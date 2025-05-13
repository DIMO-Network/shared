package privilegetoken

import (
	"math/big"

	"github.com/DIMO-Network/shared/pkg/privileges"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: move this to token-exchanges
type CustomClaims struct {
	ContractAddress common.Address         `json:"contract_address"`
	TokenID         *big.Int               `json:"token_id"`
	PrivilegeIDs    []privileges.Privilege `json:"privilege_ids"`
}

// TODO: move this to token-exchange
type Token struct {
	jwt.RegisteredClaims
	CustomClaims
}
