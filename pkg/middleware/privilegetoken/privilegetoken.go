package privilegetoken

import (
	"math/big"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/shared/pkg/logfields"
	"github.com/DIMO-Network/shared/pkg/privileges"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type VerifyPrivilegeToken struct {
	Contract  common.Address
	ChainID   uint64
	PathParam string
}

func (p *VerifyPrivilegeToken) OneOf(privilegeIDs []privileges.Privilege) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return p.verifyPrivileges(c, privilegeIDs, false)
	}
}

func (p *VerifyPrivilegeToken) AllOf(privilegeIDs []privileges.Privilege) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return p.verifyPrivileges(c, privilegeIDs, true)
	}
}

func (v *VerifyPrivilegeToken) verifyPrivileges(c *fiber.Ctx, privilegeIDs []privileges.Privilege, requireAll bool) error {
	logger := zerolog.Ctx(c.UserContext()).With().Str("src", "mw.shared.privilegetoken").Logger()

	claims, err := getTokenClaims(c)
	if err != nil {
		return err
	}
	pathParam := v.PathParam
	if pathParam == "" {
		pathParam = "tokenID"
	}
	tokenParam := c.Params(pathParam)
	tokenID, ok := big.NewInt(0).SetString(tokenParam, 10)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid token ID")
	}
	expectedSubject := cloudevent.NFTDID{
		ChainID:         v.ChainID,
		ContractAddress: v.Contract,
		TokenID:         uint32(tokenID.Uint64()),
	}.String()
	if claims.Subject != expectedSubject {
		logger.Debug().
			Uint64(logfields.VehicleTokenID, tokenID.Uint64()).
			Str("jwtSubject", claims.Subject).
			Msg("Mismatched token ids")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong subject in token")
	}
	if claims.ContractAddress != v.Contract {
		logger.Debug().
			Uint64(logfields.VehicleTokenID, tokenID.Uint64()).
			Str("jwtContractAddress", claims.ContractAddress.String()).
			Msg("Mismatched contract addresses")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong contract address in token")
	}
	if claims.TokenID.Cmp(tokenID) != 0 {
		logger.Debug().
			Uint64(logfields.VehicleTokenID, tokenID.Uint64()).
			Uint64("jwtTokenID", claims.TokenID.Uint64()).
			Msg("Mismatched token ids")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Wrong token id in token")
	}

	if requireAll {
		for _, privilege := range privilegeIDs {
			if !slices.Contains(claims.PrivilegeIDs, privilege) {
				logPrivilegeFailure(logger, claims.TokenID.String(), privilegeIDs, claims.PrivilegeIDs, "missing required")
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Token does not contain all required privileges")
			}
		}
	} else {
		for _, privilege := range privilegeIDs {
			if slices.Contains(claims.PrivilegeIDs, privilege) {
				c.Locals("tokenClaims", claims)
				return c.Next()
			}
		}
		logPrivilegeFailure(logger, claims.TokenID.String(), privilegeIDs, claims.PrivilegeIDs, "missing any")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized! Token does not contain any of the required privileges")
	}

	c.Locals("tokenClaims", claims)
	return c.Next()
}

func logPrivilegeFailure(logger zerolog.Logger, tokenID string, required []privileges.Privilege,
	actual []privileges.Privilege, reason string) {
	logger.Debug().
		Str(logfields.VehicleTokenID, tokenID).
		Interface("requiredPrivileges", required).
		Interface("actualPrivileges", actual).
		Msgf("Token %s privilege", reason)
}

func getTokenClaims(c *fiber.Ctx) (Token, error) {
	token := c.Locals("user").(*jwt.Token)
	claims, ok := token.Claims.(Token)
	if !ok {
		return Token{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}
	return claims, nil
}
