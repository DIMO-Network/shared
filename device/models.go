package device

/*
 shared models. If update this must update both vehicle-signal-decoding and edge-network and any
 other callers of vehicle-signal-decoding-api
*/

// ConfigResponse response for what templates to use, mobile app dependency: userGetVehicleDecoding.ts
type ConfigResponse struct {
	// PidURL including the version for the template
	PidURL string `json:"pidUrl"`
	// DeviceSettingURL including the version for the settings
	DeviceSettingURL string `json:"deviceSettingUrl"`
	// DbcURL including the version for the dbc file, usually same as pidurl template version
	DbcURL string `json:"dbcUrl,omitempty"`
}
