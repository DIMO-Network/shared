package privilegetoken

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type httpTestTemplate struct {
	description  string // description of the test case
	route        string // route path to test
	expectedCode int    // expected HTTP status code
}

type testHelper struct {
	app                 *fiber.App
	t                   *testing.T
	assert              *assert.Assertions
	logger              zerolog.Logger
	privilegeMiddleware IVerifyPrivilegeToken
}

func initTestHelper(t *testing.T) testHelper {
	assert := assert.New(t)
	app := fiber.New() // This can be moved into a new var to avoid re-initializing TBD.
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return testHelper{
		assert: assert,
		t:      t,
		app:    app,
		logger: logger,
		privilegeMiddleware: New(Config{
			Log: &logger,
		}),
	}
}

func (t testHelper) signToken(p jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, p)
}

const (
	Commands        = 1
	AllTimeLocation = 2
)

func TestSuccessOnValidSinglePrivilege(t *testing.T) {
	th := initTestHelper(t)

	test := httpTestTemplate{
		description:  "Test success response when token contains at only allowed privilege on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusOK,
	}

	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})

	th.app.Use(func(c *fiber.Ctx) error {
		token := th.signToken((jwt.MapClaims{
			"token_id":         "1",
			"contract_address": vehicleAddr,
			"privilege_ids":    []int64{Commands},
		}))

		c.Locals("user", token)
		return c.Next()
	})

	th.app.Get("/v1/test/:tokenID", th.privilegeMiddleware.OneOf(vehicleAddr, []int64{Commands}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := th.app.Test(req, 1)

	// Verify, if the status code is as expected
	th.assert.Equalf(test.expectedCode, resp.StatusCode, test.description)
}

func TestSuccessOnValidTokenPrivilegeOnMany(t *testing.T) {
	th := initTestHelper(t)

	test := httpTestTemplate{
		description:  "Test success response when token contains at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusOK,
	}

	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})

	th.app.Use(func(c *fiber.Ctx) error {
		token := th.signToken((jwt.MapClaims{
			"token_id":         "1",
			"contract_address": vehicleAddr,
			"privilege_ids":    []int64{Commands},
		}))

		c.Locals("user", token)
		return c.Next()
	})

	th.app.Get("/v1/test/:tokenID", th.privilegeMiddleware.OneOf(vehicleAddr, []int64{Commands, AllTimeLocation}), func(c *fiber.Ctx) error {

		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := th.app.Test(req, 1)

	// Verify, if the status code is as expected
	th.assert.Equalf(test.expectedCode, resp.StatusCode, test.description)
}

func TestMiddlewareWriteClaimsToContext(t *testing.T) {
	th := initTestHelper(t)

	test := httpTestTemplate{
		description:  "Test success response when token contains at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusOK,
	}

	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})
	cClaims := CustomClaims{
		ContractAddress: vehicleAddr,
		TokenID:         "1",
		PrivilegeIDs:    []int64{Commands},
	}
	th.app.Use(func(c *fiber.Ctx) error {
		token := th.signToken((jwt.MapClaims{
			"token_id":         cClaims.TokenID,
			"contract_address": cClaims.ContractAddress,
			"privilege_ids":    cClaims.PrivilegeIDs,
		}))

		c.Locals("user", token)
		return c.Next()
	})

	th.app.Get("/v1/test/:tokenID", th.privilegeMiddleware.OneOf(vehicleAddr, []int64{Commands, AllTimeLocation}), func(c *fiber.Ctx) error {
		cl := c.Locals("tokenClaims").(CustomClaims)
		th.assert.Equal(cl, cClaims)
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := th.app.Test(req, 1)

	// Verify, if the status code is as expected
	th.assert.Equalf(test.expectedCode, resp.StatusCode, test.description)
}

func TestFailureOnInvalidPrivilegeInToken(t *testing.T) {
	th := initTestHelper(t)

	test := httpTestTemplate{
		description:  "Test unauthorized response when token does not contain at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusUnauthorized,
	}

	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})

	th.app.Use(func(c *fiber.Ctx) error {
		token := th.signToken((jwt.MapClaims{
			"token_id":         "1",
			"contract_address": vehicleAddr,
			"privilege_ids":    []int64{Commands},
		}))

		c.Locals("user", token)
		return c.Next()
	})

	th.app.Get("/v1/test/:tokenID", th.privilegeMiddleware.OneOf(vehicleAddr, []int64{AllTimeLocation}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := th.app.Test(req, 1)

	// Verify, if the status code is as expected
	th.assert.Equalf(test.expectedCode, resp.StatusCode, test.description)
}

func TestFailureOnInvalidContractAddress(t *testing.T) {
	th := initTestHelper(t)

	test := httpTestTemplate{
		description:  "Test unauthorized response when token does not contain at least 1 of allowed privileges on endpoint",
		route:        fmt.Sprintf("/v1/test/%d", 1),
		expectedCode: fiber.StatusUnauthorized,
	}

	vehicleAddr := common.BytesToAddress([]byte{uint8(1)})

	th.app.Use(func(c *fiber.Ctx) error {
		token := th.signToken((jwt.MapClaims{
			"token_id":         "1",
			"contract_address": common.BytesToAddress([]byte{uint8(2)}),
			"privilege_ids":    []int64{Commands},
		}))

		c.Locals("user", token)
		return c.Next()
	})

	th.app.Get("/v1/test/:tokenID", th.privilegeMiddleware.OneOf(vehicleAddr, []int64{AllTimeLocation}), func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	req := httptest.NewRequest("GET", test.route, nil)

	resp, _ := th.app.Test(req, 1)

	// Verify, if the status code is as expected
	th.assert.Equalf(test.expectedCode, resp.StatusCode, test.description)
}
