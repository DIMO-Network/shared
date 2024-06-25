package sdmint

import "github.com/ethereum/go-ethereum/common"

var Type = "com.dimo.zone.device.synthetic.mint"

type Data struct {
	Integration Integration `json:"integration"`
	Vehicle     Vehicle     `json:"vehicle"`
	Device      Device      `json:"device"`
}

type Integration struct {
	TokenID       int    `json:"tokenId"`
	IntegrationID string `json:"integrationId"`
}

type Vehicle struct {
	TokenID      int    `json:"tokenId"`
	UserDeviceID string `json:"userDeviceId"`
}

type Device struct {
	TokenID           int            `json:"tokenId"`
	ExternalID        string         `json:"externalId"`
	Address           common.Address `json:"address"`
	WalletChildNumber uint32         `json:"walletChildNumber"`
}
