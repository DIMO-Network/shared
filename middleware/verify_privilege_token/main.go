package verify_privilege_token

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type VerifyPrivilegeTokenInterface interface {
	HasTokenPrivilege(privilegeID int64) fiber.Handler
}

type Config struct {
	Log *zerolog.Logger
}

type verifyPrivilegeToken struct {
	cfg Config
}

func New(cfg Config) VerifyPrivilegeTokenInterface {
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
	logger := p.cfg.Log
	claims, err := getVehicleTokenClaims(c)
	if err != nil {
		logger.Debug().Str("VehicleTokenID In Request", c.Params("tokenID")).
			Str("VehicleTokenID in bearer token", claims.VehicleTokenID).
			Msg(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "Error verifying user privilege!")
	}

	tkID := c.Params("tokenID")

	if tkID != claims.VehicleTokenID {
		logger.Debug().Str("VehicleTokenID In Request", tkID).
			Str("VehicleTokenID in bearer token", claims.VehicleTokenID).
			Msg("Invalid vehicle token")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong vehicle token provided")
	}

	// Verify privilege is correct
	privilegeFound := slices.Contains(claims.PrivilegeIDs, privilegeID)
	if !privilegeFound {
		logger.Debug().Interface("Privilege In Request", claims.PrivilegeIDs).Msg("Invalid privilege requested")
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Unauthorized! Token does not contain privilege %d.", privilegeID))
	}

	c.Locals("vehicleTokenClaims", claims)

	return c.Next()
}
