package shared

type FingerprintData struct {
	RpiUptimeSecs  int      `json:"rpiUptimeSecs,omitempty"`
	BatteryVoltage float64  `json:"batteryVoltage,omitempty"`
	Timestamp      int64    `json:"timestamp"`
	Vin            string   `json:"vin"`
	Protocol       string   `json:"protocol"`
	Odometer       float64  `json:"odometer,omitempty"`
	Errors         []string `json:"errors,omitempty"`
}
