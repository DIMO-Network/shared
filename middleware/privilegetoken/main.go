package privilegetoken

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type IVerifyPrivilegeToken interface {
	OneOf(contract common.Address, privilegeIDs []int64) fiber.Handler
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

func (p *verifyPrivilegeToken) OneOf(contract common.Address, privilegeIDs []int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return p.checkPrivilege(c, contract, privilegeIDs)
	}
}

func (p *verifyPrivilegeToken) getClaims(c *fiber.Ctx, logger zerolog.Logger) (CustomClaims, error) {

	claims, err := getTokenClaims(c)
	if err != nil {
		logger.Debug().Str("TokenID In Request", c.Params("tokenID")).
			Str("TokenID in bearer token", claims.Sub()).
			Msg(err.Error())
		return CustomClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Error verifying user privilege!")
	}

	tkID := c.Params("tokenID")

	if tkID != claims.TokenID {
		logger.Debug().Str("DeviceTokenID", tkID).
			Str("jwtTokenID", claims.TokenID).
			Msg("Mismatched token ids")
		return CustomClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong device token provided")
	}

	return claims, nil
}

func (p *verifyPrivilegeToken) checkPrivilege(c *fiber.Ctx, contract common.Address, privilegeIDs []int64) error {
	logger := p.cfg.Log.With().Str("src", "mw.shared.privilegetoken").Logger()

	// This checks that the privileges are for the token specified by the path variable :tokenID
	claims, err := p.getClaims(c, logger)
	if err != nil {
		return err
	}

	if claims.ContractAddress != contract {
		return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("Provided token is for the wrong contract: %s", claims.ContractAddress))
	}

	tkID := c.Params("tokenID")

	privilegeFound := false

	for _, v := range privilegeIDs {
		pf := slices.Contains(claims.PrivilegeIDs, v)
		if pf {
			privilegeFound = true
			break
		}
	}

	if !privilegeFound {
		logger.Debug().Str("tokenId", tkID).Interface("requiredPrivilege", privilegeIDs).Interface("jwtPrivileges", claims.PrivilegeIDs).Msg("Token lacks the required privilege.")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Token does not contain required privileges")
	}

	// Verify privilege is correct
	c.Locals("tokenClaims", claims)

	return c.Next()
}
