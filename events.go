package shared

import "time"

// CloudEventHeaders contains the fields common to all CloudEvent messages.
type CloudEventHeaders struct {
	ID          string    `json:"id"`
	Source      string    `json:"source"`
	SpecVersion string    `json:"specversion"`
	Subject     string    `json:"subject"`
	Time        time.Time `json:"time"`
	Type        string    `json:"type"`
}

// DeviceFingerprintCloudEvent contains the fields for all fingerprint events
type DeviceFingerprintCloudEvent struct {
	CloudEventHeaders
	Data FingerprintData `json:"data"`
}

type FingerprintData struct {
	FingerprintCommonData
	Vin      string  `json:"vin"`
	Protocol string  `json:"protocol"`
	Odometer float64 `json:"odometer,omitempty"`
}

// CommonData common properties we want to send with every data payload
type FingerprintCommonData struct {
	RpiUptimeSecs  int     `json:"rpiUptimeSecs,omitempty"`
	BatteryVoltage float64 `json:"batteryVoltage,omitempty"`
	Timestamp      int64   `json:"timestamp"`
}
