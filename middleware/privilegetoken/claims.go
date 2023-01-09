package privilegetoken

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	DeviceTokenID  string // Entity in this case is what we need privilege to access
	UserEthAddress string
	PrivilegeIDs   []int64
}

type privilegeTokenClaimsResponseRaw struct {
	Sub            string
	UserEthAddress string
	PrivilegeIDs   []int64
}

func getDeviceTokenClaims(c *fiber.Ctx) (TokenClaims, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return TokenClaims{}, err  
	}

	p := privilegeTokenClaimsResponseRaw{}

	if err := json.Unmarshal(jsonbody, &p); err != nil {
		return TokenClaims{}, err
	}

	return TokenClaims{
		DeviceTokenID:  p.Sub,
		UserEthAddress: p.UserEthAddress,
		PrivilegeIDs:   p.PrivilegeIDs,
	}, nil
}
