package privilegetoken

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type IVerifyPrivilegeToken interface {
	HasTokenPrivilege(privilegeID int64) fiber.Handler
}

type Config struct {
	Log *zerolog.Logger
}

type verifyPrivilegeToken struct {
	cfg Config
}

func New(cfg Config) IVerifyPrivilegeToken {
	return &verifyPrivilegeToken{
		cfg: cfg,
	}
}

func (p *verifyPrivilegeToken) HasTokenPrivilege(privilegeID int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return p.checkPrivilege(c, privilegeID)
	}
}

func (p *verifyPrivilegeToken) checkPrivilege(c *fiber.Ctx, privilegeID int64) error {
	logger := p.cfg.Log.With().Str("src", "mw.shared.privilegetoken").Logger()

	claims, err := getDeviceTokenClaims(c)
	if err != nil {
		logger.Debug().Str("DeviceTokenID In Request", c.Params("tokenID")).
			Str("DeviceTokenID in bearer token", claims.DeviceTokenID).
			Msg(err.Error())
		return fiber.NewError(fiber.StatusUnauthorized, "Error verifying user privilege!")
	}

	tkID := c.Params("tokenID")

	if tkID != claims.DeviceTokenID {
		logger.Debug().Str("DeviceTokenID In Request", tkID).
			Str("DeviceTokenID in bearer token", claims.DeviceTokenID).
			Msg("Invalid device token provided")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong device token provided")
	}

	// Verify privilege is correct
	privilegeFound := slices.Contains(claims.PrivilegeIDs, privilegeID)
	if !privilegeFound {
		logger.Debug().Str("tokenId", tkID).Int64("requiredPrivilege", privilegeID).Interface("jwtPrivileges", claims.PrivilegeIDs).Msg("Token lacks the required privilege.")
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Unauthorized! Token does not contain privilege %d.", privilegeID))
	}

	c.Locals("deviceTokenClaims", claims)

	return c.Next()
}
