package verify_privilege_token

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PrivilegeTokenClaims struct {
	DeviceTokenID  string // Entity in this case is what we need privilege to access
	UserEthAddress string
	PrivilegeIDs   []int64
}

type privilegeTokenClaimsResponseRaw struct {
	Sub            string
	UserEthAddress string
	Privileges     []int64
}

func getDeviceTokenClaims(c *fiber.Ctx) (PrivilegeTokenClaims, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return PrivilegeTokenClaims{}, err
	}

	p := privilegeTokenClaimsResponseRaw{}

	if err := json.Unmarshal(jsonbody, &p); err != nil {
		return PrivilegeTokenClaims{}, err
	}

	return PrivilegeTokenClaims{
		DeviceTokenID:  p.Sub,
		UserEthAddress: p.UserEthAddress,
		PrivilegeIDs:   p.Privileges,
	}, nil
}
