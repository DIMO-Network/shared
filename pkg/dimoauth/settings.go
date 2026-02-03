package dimoauth

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// Settings contains configuration for DIMO authentication service
type Settings struct {
	AuthURL            url.URL        `env:"AUTH_URL"              yaml:"AUTH_URL"`
	TokenExchangeURL   url.URL        `env:"TOKEN_EXCHANGE_URL"    yaml:"TOKEN_EXCHANGE_URL"`
	NFTContractAddress common.Address `env:"NFT_CONTRACT_ADDRESS"  yaml:"NFT_CONTRACT_ADDRESS"`
	ClientID           common.Address `env:"CLIENT_ID"             yaml:"CLIENT_ID"`
	RedirectURL        url.URL        `env:"REDIRECT_URL"          yaml:"REDIRECT_URL"`
	PrivateKeyHex      string         `env:"PRIVATE_KEY_HEX"       yaml:"PRIVATE_KEY_HEX"`
}
