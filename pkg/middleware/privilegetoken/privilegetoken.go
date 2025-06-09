// Package privilegetoken provides a middleware that verifies the privileges of a token.
package privilegetoken

import (
	"math/big"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/shared/pkg/privileges"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

const (
	tokenIDKey = "tokenId"
)

// VerifyPrivilegeToken is a middleware that verifies the privileges of a token.
type VerifyPrivilegeToken struct {
	Contract  common.Address
	ChainID   uint64
	PathParam string
}

// OneOf verifies that the token contains at least one of the specified privileges.
func (v *VerifyPrivilegeToken) OneOf(privilegeIDs []privileges.Privilege) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return v.verifyPrivileges(c, privilegeIDs, false)
	}
}

// AllOf verifies that the token contains all of the specified privileges.
func (v *VerifyPrivilegeToken) AllOf(privilegeIDs []privileges.Privilege) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return v.verifyPrivileges(c, privilegeIDs, true)
	}
}

// verifyPrivileges verifies the privileges of a token.
func (v *VerifyPrivilegeToken) verifyPrivileges(c *fiber.Ctx, privilegeIDs []privileges.Privilege, requireAll bool) error {
	claims, err := getTokenClaims(c)
	if err != nil {
		return err
	}

	tokenID, err := validateTokenParams(c, v.PathParam)
	if err != nil {
		return err
	}

	if err := validateTokenDetails(c, claims, tokenID, v.Contract, v.ChainID); err != nil {
		return err
	}

	return checkPrivileges(c, claims, tokenID, privilegeIDs, requireAll)
}

// validateTokenParams validates the token ID from the path parameters.
func validateTokenParams(c *fiber.Ctx, pathParam string) (*big.Int, error) {
	if pathParam == "" {
		pathParam = "tokenID"
	}
	tokenParam := c.Params(pathParam)
	tokenID, ok := big.NewInt(0).SetString(tokenParam, 10)
	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid token ID")
	}
	return tokenID, nil
}

// validateTokenDetails validates the token details.
func validateTokenDetails(c *fiber.Ctx, claims Token, tokenID *big.Int, contract common.Address, chainID uint64) error {
	expectedSubject := cloudevent.NFTDID{
		ChainID:         chainID,
		ContractAddress: contract,
		TokenID:         uint32(tokenID.Uint64()), //TODO: Update this when cloudevent uses big.Int
	}.String()

	if claims.Subject != expectedSubject {
		logger(c).Debug().
			Str("tokenId", tokenID.String()).
			Str("jwtSubject", claims.Subject).
			Msg("Mismatched token subject")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Invalid token subject")
	}
	if claims.ContractAddress != contract {
		logger(c).Debug().
			Str(tokenIDKey, tokenID.String()).
			Str("jwtContractAddress", claims.ContractAddress.String()).
			Msg("Mismatched contract addresses")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Invalid contract address")
	}
	if claims.TokenID.Cmp(tokenID) != 0 {
		logger(c).Debug().
			Str(tokenIDKey, tokenID.String()).
			Str("jwtTokenId", claims.TokenID.String()).
			Msg("Mismatched token Ids")
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Invalid token ID")
	}
	return nil
}

// checkPrivileges checks the privileges of a token.
func checkPrivileges(c *fiber.Ctx, claims Token, tokenID *big.Int, privilegeIDs []privileges.Privilege, requireAll bool) error {
	if requireAll {
		for _, privilege := range privilegeIDs {
			if !slices.Contains(claims.PrivilegeIDs, privilege) {
				logger(c).Debug().
					Str(tokenIDKey, tokenID.String()).
					Interface("requiredPrivileges", privilegeIDs).
					Interface("actualPrivileges", claims.PrivilegeIDs).
					Msg("Token privilege verification failed: missing all required")
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Token does not contain all required privileges")
			}
		}
		c.Locals("tokenClaims", claims)
		return c.Next()
	}

	// If not all, then at least one is required
	for _, privilege := range privilegeIDs {
		if slices.Contains(claims.PrivilegeIDs, privilege) {
			c.Locals("tokenClaims", claims)
			return c.Next()
		}
	}
	logger(c).Debug().
		Str(tokenIDKey, tokenID.String()).
		Interface("requiredPrivileges", privilegeIDs).
		Interface("actualPrivileges", claims.PrivilegeIDs).
		Msg("Token privilege verification failed: missing any required")
	return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Token does not contain any of the required privileges")
}

func getTokenClaims(c *fiber.Ctx) (Token, error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return Token{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token in context")
	}
	claims, ok := token.Claims.(Token)
	if !ok {
		return Token{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}
	return claims, nil
}

func logger(c *fiber.Ctx) *zerolog.Logger {
	logger := zerolog.Ctx(c.UserContext()).With().Str("src", "mw.shared.privilegetoken").Logger()
	return &logger
}
