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

type FingerprintData struct {
	RpiUptimeSecs  int      `json:"rpiUptimeSecs,omitempty"`
	BatteryVoltage float64  `json:"batteryVoltage,omitempty"`
	Timestamp      int64    `json:"timestamp"`
	Vin            string   `json:"vin"`
	Protocol       string   `json:"protocol"`
	Odometer       float64  `json:"odometer,omitempty"`
	Errors         []string `json:"errors,omitempty"`
}
