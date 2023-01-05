package tokenexchange

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type IPrivilegeHandler interface {
	HasTokenPrivilege(privilegeID int64) fiber.Handler
}

type PrivilegeHandler struct {
	Log *zerolog.Logger
}

func New(cfg PrivilegeHandler) IPrivilegeHandler {
	return &PrivilegeHandler{
		Log: cfg.Log,
	}
}

func (p *PrivilegeHandler) HasTokenPrivilege(privilegeID int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return p.checkPrivilege(c, privilegeID)
	}
}

func (p *PrivilegeHandler) checkPrivilege(c *fiber.Ctx, privilegeID int64) error {
	claims, err := getVehicleTokenClaims(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	tkID := c.Params("tokenID")

	if tkID != claims.VehicleTokenID {
		p.Log.Debug().Str("VehicleTokenID In Request", tkID).
			Str("VehicleTokenID in bearer token", claims.VehicleTokenID).
			Msg("Invalid vehicle token")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong vehicle token provided")
	}

	privilegeFound := slices.Contains(claims.Privileges, privilegeID)
	if !privilegeFound {
		p.Log.Debug().Interface("Privilege In Request", claims.Privileges).Msg("Invalid privilege requested")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Token does not contain privilege 1.")
	}

	// Verify privilege is correct
	c.Locals("vehicleTokenClaims", VehicleTokenClaims{
		VehicleTokenID: claims.VehicleTokenID,
		UserEthAddress: claims.UserEthAddress,
		Privileges:     claims.Privileges,
	})

	return c.Next()
}
