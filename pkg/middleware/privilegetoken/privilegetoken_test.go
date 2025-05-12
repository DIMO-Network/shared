package privilegetoken

import (
	"fmt"
	"io"
	"math/big"
	"net/http/httptest"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/shared/pkg/privileges"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpTestTemplate struct {
	description  string // description of the test case
	route        string // route path to test
	expectedCode int    // expected HTTP status code
}

func signToken(p jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, p)
}
func TestSuccessOnValidSinglePrivilege(t *testing.T) {
	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})
	privilegeMiddleware := VerifyPrivilegeToken{
		Contract:  vehicleAddr,
		ChainID:   1,
		PathParam: "tokenID",
	}
	app := fiber.New()

	test := httpTestTemplate{
		description:  "Test success response when token contains at only allowed privilege on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusOK,
	}

	app.Use(func(c *fiber.Ctx) error {
		token := signToken(Token{
			CustomClaims: CustomClaims{
				ContractAddress: vehicleAddr,
				TokenID:         big.NewInt(1),
				PrivilegeIDs:    []privileges.Privilege{privileges.VehicleCommands},
			},
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: cloudevent.NFTDID{
					ChainID:         1,
					ContractAddress: vehicleAddr,
					TokenID:         1,
				}.String(),
			},
		})

		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/v1/test/:tokenID", privilegeMiddleware.OneOf([]privileges.Privilege{privileges.VehicleCommands}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, err := app.Test(req)
	require.NoError(t, err)
	body, _ := io.ReadAll(resp.Body)
	// Verify, if the status code is as expected
	assert.Equalf(t, test.expectedCode, resp.StatusCode, string(body))
}

func TestSuccessOnValidTokenPrivilegeOnMany(t *testing.T) {
	app := fiber.New()
	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})
	privilegeMiddleware := VerifyPrivilegeToken{
		Contract:  vehicleAddr,
		ChainID:   1,
		PathParam: "tokenID",
	}

	test := httpTestTemplate{
		description:  "Test success response when token contains at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusOK,
	}

	app.Use(func(c *fiber.Ctx) error {
		token := signToken(Token{
			CustomClaims: CustomClaims{
				ContractAddress: vehicleAddr,
				TokenID:         big.NewInt(1),
				PrivilegeIDs:    []privileges.Privilege{privileges.VehicleCommands},
			},
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: cloudevent.NFTDID{
					ChainID:         1,
					ContractAddress: vehicleAddr,
					TokenID:         1,
				}.String(),
			},
		})

		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/v1/test/:tokenID", privilegeMiddleware.OneOf([]privileges.Privilege{privileges.VehicleCommands, privileges.VehicleAllTimeLocation}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := app.Test(req, 1)

	// Verify, if the status code is as expected
	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}

func TestFailureOnInvalidPrivilegeInToken(t *testing.T) {
	app := fiber.New()
	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})
	privilegeMiddleware := VerifyPrivilegeToken{
		Contract:  vehicleAddr,
		ChainID:   1,
		PathParam: "tokenID",
	}

	test := httpTestTemplate{
		description:  "Test unauthorized response when token does not contain at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusUnauthorized,
	}

	app.Use(func(c *fiber.Ctx) error {
		token := signToken(Token{
			CustomClaims: CustomClaims{
				ContractAddress: vehicleAddr,
				TokenID:         big.NewInt(1),
				PrivilegeIDs:    []privileges.Privilege{privileges.VehicleCommands},
			},
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: cloudevent.NFTDID{
					ChainID:         1,
					ContractAddress: vehicleAddr,
					TokenID:         1,
				}.String(),
			},
		})

		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/v1/test/:tokenID", privilegeMiddleware.OneOf([]privileges.Privilege{privileges.VehicleAllTimeLocation}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := app.Test(req, 1)

	// Verify, if the status code is as expected
	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}

func TestFailureOnInvalidContractAddress(t *testing.T) {
	app := fiber.New()
	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})
	privilegeMiddleware := VerifyPrivilegeToken{
		Contract:  vehicleAddr,
		ChainID:   1,
		PathParam: "tokenID",
	}

	test := httpTestTemplate{
		description:  "Test unauthorized response when token does not contain at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusUnauthorized,
	}

	app.Use(func(c *fiber.Ctx) error {
		token := signToken(Token{
			CustomClaims: CustomClaims{
				ContractAddress: vehicleAddr,
				TokenID:         big.NewInt(1),
				PrivilegeIDs:    []privileges.Privilege{privileges.VehicleCommands},
			},
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: cloudevent.NFTDID{
					ChainID:         1,
					ContractAddress: vehicleAddr,
					TokenID:         1,
				}.String(),
			},
		})
		c.Locals("user", token)
		return c.Next()
	})

	app.Get("/v1/test/:tokenID", privilegeMiddleware.OneOf([]privileges.Privilege{privileges.VehicleAllTimeLocation}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := app.Test(req, 1)

	// Verify, if the status code is as expected
	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}
