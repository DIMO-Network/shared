package sdtask

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type CredentialData struct {
	TaskID        string    `json:"taskId"`
	UserDeviceID  string    `json:"userDeviceId"`
	IntegrationID string    `json:"integrationId"`
	AccessToken   string    `json:"accessToken"`
	Expiry        time.Time `json:"expiry"`
	RefreshToken  string    `json:"refreshToken"`
	// Version is an API version. Currently only used for Tesla: 1 is the old
	// "Owner api" and 2 is the new "Fleet API".
	Version int `json:"version,omitempty"`
	// Region is the two-letter code for the region. Currently only used for
	// Tesla, and can only be "na" or "eu".
	Region          string           `json:"region,omitempty"`
	SyntheticDevice *SyntheticDevice `json:"syntheticDevice,omitempty"`
}

type SyntheticDevice struct {
	TokenID            int            `json:"tokenId"`
	Address            common.Address `json:"address"`
	IntegrationTokenID int            `json:"integrationTokenId"`
	WalletChildNumber  int            `json:"walletChildNumber"`
	VehicleTokenID     int            `json:"vehicleTokenId"`
}
