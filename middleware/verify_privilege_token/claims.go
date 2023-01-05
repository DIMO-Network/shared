package verify_privilege_token

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type VehicleTokenClaims struct {
	VehicleTokenID string
	UserEthAddress string
	PrivilegeIDs   []int64
}

type vehicleTokenClaimsResponseRaw struct {
	Sub            string
	UserEthAddress string
	Privileges     []int64
}

func getVehicleTokenClaims(c *fiber.Ctx) (VehicleTokenClaims, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return VehicleTokenClaims{}, err
	}

	p := vehicleTokenClaimsResponseRaw{}

	if err := json.Unmarshal(jsonbody, &p); err != nil {
		return VehicleTokenClaims{}, err
	}

	return VehicleTokenClaims{
		VehicleTokenID: p.Sub,
		UserEthAddress: p.UserEthAddress,
		PrivilegeIDs:   p.Privileges,
	}, nil
}
