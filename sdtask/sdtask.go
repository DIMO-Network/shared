package sdtask

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type CredentialData struct {
	TaskID          string           `json:"taskId"`
	UserDeviceID    string           `json:"userDeviceId"`
	IntegrationID   string           `json:"integrationId"`
	AccessToken     string           `json:"accessToken"`
	Expiry          time.Time        `json:"expiry"`
	RefreshToken    string           `json:"refreshToken"`
	Version         int              `json:"version,omitempty"`
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
