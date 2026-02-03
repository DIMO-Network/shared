package dimoauth

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// Settings contains configuration for DIMO authentication service
type Settings struct {
	AuthURL            string `env:"AUTH_URL"              yaml:"AUTH_URL"`
	TokenExchangeURL   string `env:"TOKEN_EXCHANGE_URL"    yaml:"TOKEN_EXCHANGE_URL"`
	NFTContractAddress string `env:"NFT_CONTRACT_ADDRESS"  yaml:"NFT_CONTRACT_ADDRESS"`
	ClientID           string `env:"CLIENT_ID"             yaml:"CLIENT_ID"`
	Domain             string `env:"DOMAIN"                yaml:"DOMAIN"`
	PrivateKeyHex      string `env:"PRIVATE_KEY_HEX"       yaml:"PRIVATE_KEY_HEX"`
}

// ParseURLs converts string URLs to url.URL and validates addresses
func (s *Settings) ParseURLs() (authURL url.URL, tokenExchangeURL url.URL, nftContractAddress common.Address, clientID common.Address, err error) {
	parsedAuthURL, err := url.Parse(s.AuthURL)
	if err != nil {
		return url.URL{}, url.URL{}, common.Address{}, common.Address{}, err
	}

	parsedTokenExchangeURL, err := url.Parse(s.TokenExchangeURL)
	if err != nil {
		return url.URL{}, url.URL{}, common.Address{}, common.Address{}, err
	}

	nftContractAddress = common.HexToAddress(s.NFTContractAddress)
	clientID = common.HexToAddress(s.ClientID)

	return *parsedAuthURL, *parsedTokenExchangeURL, nftContractAddress, clientID, nil
}
