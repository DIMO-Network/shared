package shared

import (
	"fmt"

	"github.com/clarketm/json"
)

func (d DataSchemaStruct) IsValid() (bool, error) {
	valid := []string{"SAE_0", "SAE_1", "SAE_2_DISENGAGING", "SAE_2", "SAE_3_DISENGAGING", "SAE_3", "SAE_4_DISENGAGING", "SAE_4", "SAE_5", ""}

	val := d.Vehicle.ADAS.ActiveAutonomyLevel
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.ADAS.ActiveAutonomyLevel: %v", val)
	}

	valid = []string{"SAE_0", "SAE_1", "SAE_2", "SAE_3", "SAE_4", "SAE_5", ""}

	val = d.Vehicle.ADAS.SupportedAutonomyLevel
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.ADAS.SupportedAutonomyLevel: %v", val)
	}

	valid = []string{"INACTIVE", "ACTIVE", "ADAPTIVE", ""}

	val = d.Vehicle.Body.Lights.Brake.IsActive
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Lights.Brake.IsActive: %v", val)
	}

	valid = []string{"FRONT_LEFT", "FRONT_RIGHT", "MIDDLE_LEFT", "MIDDLE_RIGHT", "REAR_LEFT", "REAR_RIGHT", ""}

	val = d.Vehicle.Body.RefuelPosition
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.RefuelPosition: %v", val)
	}

	valid = []string{"OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR", ""}

	val = d.Vehicle.Body.Windshield.Front.Wiping.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Front.Wiping.Mode: %v", val)
	}

	valid = []string{"STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP", ""}

	val = d.Vehicle.Body.Windshield.Front.Wiping.System.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Front.Wiping.System.Mode: %v", val)
	}

	valid = []string{"OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR", ""}

	val = d.Vehicle.Body.Windshield.Rear.Wiping.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Rear.Wiping.Mode: %v", val)
	}

	valid = []string{"STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP", ""}

	val = d.Vehicle.Body.Windshield.Rear.Wiping.System.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Rear.Wiping.System.Mode: %v", val)
	}

	valid = []string{"UNDEFINED", "CLOSED", "OPEN", "CLOSING", "OPENING", "STALLED", ""}

	val = d.Vehicle.Cabin.Convertible.Status
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Convertible.Status: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Left.Shade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Left.Shade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Left.Window.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Left.Window.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Right.Shade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Right.Shade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Right.Window.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Right.Window.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Left.Shade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Left.Shade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Left.Window.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Left.Window.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Right.Shade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Right.Shade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Right.Window.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Right.Window.Switch: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row1.Left.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row1.Left.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row1.Right.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row1.Right.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row2.Left.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row2.Left.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row2.Right.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row2.Right.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row3.Left.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row3.Left.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row3.Right.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row3.Right.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row4.Left.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row4.Left.AirDistribution: %v", val)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row4.Right.AirDistribution
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row4.Right.AirDistribution: %v", val)
	}

	valid = []string{"YYYY_MM_DD", "DD_MM_YYYY", "MM_DD_YYYY", "YY_MM_DD", "DD_MM_YY", "MM_DD_YY", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DateFormat
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DateFormat: %v", val)
	}

	valid = []string{"DAY", "NIGHT", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DayNightMode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DayNightMode: %v", val)
	}

	valid = []string{"MILES", "KILOMETERS", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DistanceUnit
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DistanceUnit: %v", val)
	}

	valid = []string{"MILES_PER_KILOWATT_HOUR", "KILOMETERS_PER_KILOWATT_HOUR", "KILOWATT_HOURS_PER_100_MILES", "KILOWATT_HOURS_PER_100_KILOMETERS", "WATT_HOURS_PER_MILE", "WATT_HOURS_PER_KILOMETER", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.EVEconomyUnits
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.EVEconomyUnits: %v", val)
	}

	valid = []string{"MPG_UK", "MPG_US", "MILES_PER_LITER", "KILOMETERS_PER_LITER", "LITERS_PER_100_KILOMETERS", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.FuelEconomyUnits
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.FuelEconomyUnits: %v", val)
	}

	valid = []string{"C", "F", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TemperatureUnit
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TemperatureUnit: %v", val)
	}

	valid = []string{"HR_12", "HR_24", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TimeFormat
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TimeFormat: %v", val)
	}

	valid = []string{"‘PSI’", "‘KPA’", "’BAR’", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TirePressureUnit
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TirePressureUnit: %v", val)
	}

	valid = []string{"UNKNOWN", "STOP", "PLAY", "FAST_FORWARD", "FAST_BACKWARD", "SKIP_FORWARD", "SKIP_BACKWARD", ""}

	val = d.Vehicle.Cabin.Infotainment.Media.Action
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Media.Action: %v", val)
	}

	valid = []string{"UNKNOWN", "SIRIUS_XM", "AM", "FM", "DAB", "TV", "CD", "DVD", "AUX", "USB", "DISK", "BLUETOOTH", "INTERNET", "VOICE", "BEEP", ""}

	val = d.Vehicle.Cabin.Infotainment.Media.Played.Source
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Media.Played.Source: %v", val)
	}

	valid = []string{"MUTED", "ALERT_ONLY", "UNMUTED", ""}

	val = d.Vehicle.Cabin.Infotainment.Navigation.Mute
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Navigation.Mute: %v", val)
	}

	valid = []string{"NONE", "ACTIVE", "INACTIVE", ""}

	val = d.Vehicle.Cabin.Infotainment.SmartphoneProjection.Active
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.SmartphoneProjection.Active: %v", val)
	}

	valid = []string{"USB", "BLUETOOTH", "WIFI", ""}

	val = d.Vehicle.Cabin.Infotainment.SmartphoneProjection.Source
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.SmartphoneProjection.Source: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.RearShade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.RearShade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Sunroof.Shade.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Sunroof.Shade.Switch: %v", val)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", "TILT_UP", "TILT_DOWN", ""}

	val = d.Vehicle.Cabin.Sunroof.Switch
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Sunroof.Switch: %v", val)
	}

	valid = []string{"FRONT_LEFT", "FRONT_RIGHT", ""}

	val = d.Vehicle.Chassis.SteeringWheel.Position
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Chassis.SteeringWheel.Position: %v", val)
	}

	valid = []string{"NONE", "TWO_D", "TWO_D_SATELLITE_BASED_AUGMENTATION", "TWO_D_GROUND_BASED_AUGMENTATION", "TWO_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", "THREE_D", "THREE_D_SATELLITE_BASED_AUGMENTATION", "THREE_D_GROUND_BASED_AUGMENTATION", "THREE_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", ""}

	val = d.Vehicle.CurrentLocation.GNSSReceiver.FixType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.CurrentLocation.GNSSReceiver.FixType: %v", val)
	}

	valid = []string{"UNDEFINED", "LOCK", "OFF", "ACC", "ON", "START", ""}

	val = d.Vehicle.LowVoltageSystemState
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.LowVoltageSystemState: %v", val)
	}

	valid = []string{"SPARK", "COMPRESSION", ""}

	val = d.Vehicle.OBD.DriveCycleStatus.IgnitionType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.OBD.DriveCycleStatus.IgnitionType: %v", val)
	}

	valid = []string{"SPARK", "COMPRESSION", ""}

	val = d.Vehicle.OBD.Status.IgnitionType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.OBD.Status.IgnitionType: %v", val)
	}

	valid = []string{"UNKNOWN", "NATURAL", "SUPERCHARGER", "TURBOCHARGER", ""}

	val = d.Vehicle.Powertrain.CombustionEngine.AspirationType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.AspirationType: %v", val)
	}

	valid = []string{"UNKNOWN", "STRAIGHT", "V", "BOXER", "W", "ROTARY", "RADIAL", "SQUARE", "H", "U", "OPPOSED", "X", ""}

	val = d.Vehicle.Powertrain.CombustionEngine.Configuration
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.Configuration: %v", val)
	}

	valid = []string{"CRITICALLY_LOW", "LOW", "NORMAL", "HIGH", "CRITICALLY_HIGH", ""}

	val = d.Vehicle.Powertrain.CombustionEngine.EngineOilLevel
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.EngineOilLevel: %v", val)
	}

	valid = []string{"UNKNOWN", "NOT_APPLICABLE", "STOP_START", "BELT_ISG", "CIMG", "PHEV", ""}

	val = d.Vehicle.Powertrain.FuelSystem.HybridType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.FuelSystem.HybridType: %v", val)
	}

	valid = []string{"OPEN", "CLOSED", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.ChargePortFlap
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.ChargePortFlap: %v", val)
	}

	valid = []string{"MANUAL", "TIMER", "GRID", "PROFILE", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.Mode: %v", val)
	}

	valid = []string{"START", "STOP", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.StartStopCharging
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.StartStopCharging: %v", val)
	}

	valid = []string{"INACTIVE", "START_TIME", "END_TIME", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.Timer.Mode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.Timer.Mode: %v", val)
	}

	valid = []string{"UNKNOWN", "FORWARD_WHEEL_DRIVE", "REAR_WHEEL_DRIVE", "ALL_WHEEL_DRIVE", ""}

	val = d.Vehicle.Powertrain.Transmission.DriveType
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.DriveType: %v", val)
	}

	valid = []string{"MANUAL", "AUTOMATIC", ""}

	val = d.Vehicle.Powertrain.Transmission.GearChangeMode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.GearChangeMode: %v", val)
	}

	valid = []string{"NORMAL", "SPORT", "ECONOMY", "SNOW", "RAIN", ""}

	val = d.Vehicle.Powertrain.Transmission.PerformanceMode
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.PerformanceMode: %v", val)
	}

	valid = []string{"UNKNOWN", "SEQUENTIAL", "H", "AUTOMATIC", "DSG", "CVT", ""}

	val = d.Vehicle.Powertrain.Transmission.Type
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.Type: %v", val)
	}

	valid = []string{"COMBUSTION", "HYBRID", "ELECTRIC", ""}

	val = d.Vehicle.Powertrain.Type
	if !contains(valid, val) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Type: %v", val)
	}

	return true, nil
}

// Spatial acceleration. Axis definitions according to ISO 8855.
type Acceleration struct {
	// Vehicle acceleration in Y (lateral acceleration).
	Lateral float64 `json:"lateral,omitempty"`
	// Vehicle acceleration in X (longitudinal acceleration).
	Longitudinal float64 `json:"longitudinal,omitempty"`
	// Vehicle acceleration in Z (vertical acceleration).
	Vertical float64 `json:"vertical,omitempty"`
}

// Spatial rotation. Axis definitions according to ISO 8855.
type AngularVelocity struct {
	// Vehicle rotation rate along Y (lateral).
	Pitch float64 `json:"pitch,omitempty"`
	// Vehicle rotation rate along X (longitudinal).
	Roll float64 `json:"roll,omitempty"`
	// Vehicle rotation rate along Z (vertical).
	Yaw float64 `json:"yaw,omitempty"`
}

// All Advanced Driver Assist Systems data.
type DriveAssistSystems struct {
	// Antilock Braking System signals.
	ABS struct {
		// Indicates if ABS is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if ABS is currently regulating brake pressure. True = Engaged. False = Not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
		// Indicates if ABS incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
	} `json:"aBS,omitempty"`
	// Indicates the currently active level of autonomy according to SAE J3016 taxonomy. Must be one of ["SAE_0", "SAE_1", "SAE_2_DISENGAGING", "SAE_2", "SAE_3_DISENGAGING", "SAE_3", "SAE_4_DISENGAGING", "SAE_4", "SAE_5"]
	ActiveAutonomyLevel string `json:"activeAutonomyLevel,omitempty"`
	// Signals from Cruise Control system.
	CruiseControl struct {
		// Indicates if cruise control system is active (i.e. actively controls speed). True = Active. False = Inactive.
		IsActive bool `json:"isActive,omitempty"`
		// Indicates if cruise control system is enabled (e.g. ready to receive configurations and settings) True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if cruise control system incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
		// Set cruise control speed in kilometers per hour.
		SpeedSet float64 `json:"speedSet,omitempty"`
	} `json:"cruiseControl,omitempty"`
	// Emergency Brake Assist (EBA) System signals.
	EBA struct {
		// Indicates if EBA is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if EBA is currently regulating brake pressure. True = Engaged. False = Not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
		// Indicates if EBA incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
	} `json:"eBA,omitempty"`
	// Electronic Brakeforce Distribution (EBD) System signals.
	EBD struct {
		// Indicates if EBD is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if EBD is currently regulating vehicle brakeforce distribution. True = Engaged. False = Not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
		// Indicates if EBD incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
	} `json:"eBD,omitempty"`
	// Electronic Stability Control System signals.
	ESC struct {
		// Indicates if ESC is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if ESC is currently regulating vehicle stability. True = Engaged. False = Not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
		// Indicates if ESC incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
		// Indicates if the ESC system is detecting strong cross winds. True = Strong cross winds detected. False = No strong cross winds detected.
		IsStrongCrossWindDetected bool `json:"isStrongCrossWindDetected,omitempty"`
		// Road friction values reported by the ESC system.
		RoadFriction struct {
			// Lower bound road friction, as calculated by the ESC system. 5% possibility that road friction is below this value. 0 = no friction, 100 = maximum friction.
			LowerBound float64 `json:"lowerBound,omitempty"`
			// Most probable road friction, as calculated by the ESC system. Exact meaning of most probable is implementation specific. 0 = no friction, 100 = maximum friction.
			MostProbable float64 `json:"mostProbable,omitempty"`
			// Upper bound road friction, as calculated by the ESC system. 95% possibility that road friction is below this value. 0 = no friction, 100 = maximum friction.
			UpperBound float64 `json:"upperBound,omitempty"`
		} `json:"roadFriction,omitempty"`
	} `json:"eSC,omitempty"`
	// Signals from Lane Departure Detection System.
	LaneDepartureDetection struct {
		// Indicates if lane departure detection system is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if lane departure system incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
		// Indicates if lane departure detection registered a lane departure.
		IsWarning bool `json:"isWarning,omitempty"`
	} `json:"laneDepartureDetection,omitempty"`
	// Signals form Obstacle Sensor System.
	ObstacleDetection struct {
		// Indicates if obstacle sensor system is enabled (i.e. monitoring for obstacles). True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if obstacle sensor system incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
		// Indicates if obstacle sensor system registered an obstacle.
		IsWarning bool `json:"isWarning,omitempty"`
	} `json:"obstacleDetection,omitempty"`
	// Indicates the highest level of autonomy according to SAE J3016 taxonomy the vehicle is capable of. Must be one of ["SAE_0", "SAE_1", "SAE_2", "SAE_3", "SAE_4", "SAE_5"]
	SupportedAutonomyLevel string `json:"supportedAutonomyLevel,omitempty"`
	// Traction Control System signals.
	TCS struct {
		// Indicates if TCS is enabled. True = Enabled. False = Disabled.
		IsEnabled bool `json:"isEnabled,omitempty"`
		// Indicates if TCS is currently regulating traction. True = Engaged. False = Not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
		// Indicates if TCS incurred an error condition. True = Error. False = No Error.
		IsError bool `json:"isError,omitempty"`
	} `json:"tCS,omitempty"`
}

// All body components.
type Body struct {
	// Body type code as defined by ISO 3779.
	BodyType string `json:"bodyType,omitempty"`
	// Hood status.
	Hood struct {
		// Hood open or closed. True = Open. False = Closed.
		IsOpen bool `json:"isOpen,omitempty"`
	} `json:"hood,omitempty"`
	// Horn signals.
	Horn struct {
		// Horn active or inactive. True = Active. False = Inactive.
		IsActive bool `json:"isActive,omitempty"`
	} `json:"horn,omitempty"`
	// Exterior lights.
	Lights struct {
		// Backup lights.
		Backup struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn bool `json:"isOn,omitempty"`
		} `json:"backup,omitempty"`
		// Beam lights.
		Beam struct {
			// Beam lights.
			High struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn bool `json:"isOn,omitempty"`
			} `json:"high,omitempty"`
			// Beam lights.
			Low struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn bool `json:"isOn,omitempty"`
			} `json:"low,omitempty"`
		} `json:"beam,omitempty"`
		//
		Brake struct {
			// Indicates if break-light is active. INACTIVE means lights are off. ACTIVE means lights are on. ADAPTIVE means that break-light is indicating emergency-breaking. Must be one of ["INACTIVE", "ACTIVE", "ADAPTIVE"]
			IsActive string `json:"isActive,omitempty"`
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
		} `json:"brake,omitempty"`
		// Indicator lights.
		DirectionIndicator struct {
			// Indicator lights.
			Left struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is signaling or off. True = signaling. False = Off.
				IsSignaling bool `json:"isSignaling,omitempty"`
			} `json:"left,omitempty"`
			// Indicator lights.
			Right struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is signaling or off. True = signaling. False = Off.
				IsSignaling bool `json:"isSignaling,omitempty"`
			} `json:"right,omitempty"`
		} `json:"directionIndicator,omitempty"`
		// Fog lights.
		Fog struct {
			// Fog lights.
			Front struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn bool `json:"isOn,omitempty"`
			} `json:"front,omitempty"`
			// Fog lights.
			Rear struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn bool `json:"isOn,omitempty"`
			} `json:"rear,omitempty"`
		} `json:"fog,omitempty"`
		// Hazard lights.
		Hazard struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
			// Indicates if light is signaling or off. True = signaling. False = Off.
			IsSignaling bool `json:"isSignaling,omitempty"`
		} `json:"hazard,omitempty"`
		// License plate lights.
		LicensePlate struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn bool `json:"isOn,omitempty"`
		} `json:"licensePlate,omitempty"`
		// Parking lights.
		Parking struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn bool `json:"isOn,omitempty"`
		} `json:"parking,omitempty"`
		// Running lights.
		Running struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn bool `json:"isOn,omitempty"`
		} `json:"running,omitempty"`
	} `json:"lights,omitempty"`
	// All mirrors.
	Mirrors struct {
		// All mirrors.
		Left struct {
			// Mirror Heater on or off. True = Heater On. False = Heater Off.
			IsHeatingOn bool `json:"isHeatingOn,omitempty"`
			// Mirror pan as a percent. 0 = Center Position. 100 = Fully Left Position. -100 = Fully Right Position.
			Pan int8 `json:"pan,omitempty"`
			// Mirror tilt as a percent. 0 = Center Position. 100 = Fully Upward Position. -100 = Fully Downward Position.
			Tilt int8 `json:"tilt,omitempty"`
		} `json:"left,omitempty"`
		// All mirrors.
		Right struct {
			// Mirror Heater on or off. True = Heater On. False = Heater Off.
			IsHeatingOn bool `json:"isHeatingOn,omitempty"`
			// Mirror pan as a percent. 0 = Center Position. 100 = Fully Left Position. -100 = Fully Right Position.
			Pan int8 `json:"pan,omitempty"`
			// Mirror tilt as a percent. 0 = Center Position. 100 = Fully Upward Position. -100 = Fully Downward Position.
			Tilt int8 `json:"tilt,omitempty"`
		} `json:"right,omitempty"`
	} `json:"mirrors,omitempty"`
	// Rainsensor signals.
	Raindetection struct {
		// Rain intensity. 0 = Dry, No Rain. 100 = Covered.
		Intensity uint8 `json:"intensity,omitempty"`
	} `json:"raindetection,omitempty"`
	// Rear spoiler position, 0% = Spoiler fully stowed. 100% = Spoiler fully exposed.
	RearMainSpoilerPosition float64 `json:"rearMainSpoilerPosition,omitempty"`
	// Location of the fuel cap or charge port. Must be one of ["FRONT_LEFT", "FRONT_RIGHT", "MIDDLE_LEFT", "MIDDLE_RIGHT", "REAR_LEFT", "REAR_RIGHT"]
	RefuelPosition string `json:"refuelPosition,omitempty"`
	// Trunk status.
	Trunk struct {
		// Trunk status.
		Front struct {
			// Is trunk locked or unlocked. True = Locked. False = Unlocked.
			IsLocked bool `json:"isLocked,omitempty"`
			// Trunk open or closed. True = Open. False = Closed.
			IsOpen bool `json:"isOpen,omitempty"`
		} `json:"front,omitempty"`
		// Trunk status.
		Rear struct {
			// Is trunk locked or unlocked. True = Locked. False = Unlocked.
			IsLocked bool `json:"isLocked,omitempty"`
			// Trunk open or closed. True = Open. False = Closed.
			IsOpen bool `json:"isOpen,omitempty"`
		} `json:"rear,omitempty"`
	} `json:"trunk,omitempty"`
	// Windshield signals.
	Windshield struct {
		// Windshield signals.
		Front struct {
			// Windshield heater status. False - off, True - on.
			IsHeatingOn bool `json:"isHeatingOn,omitempty"`
			// Windshield washer fluid signals
			WasherFluid struct {
				// Low level indication for washer fluid. True = Level Low. False = Level OK.
				IsLevelLow bool `json:"isLevelLow,omitempty"`
				// Washer fluid level as a percent. 0 = Empty. 100 = Full.
				Level uint8 `json:"level,omitempty"`
			} `json:"washerFluid,omitempty"`
			// Windshield wiper signals.
			Wiping struct {
				// Relative intensity/sensitivity for interval and rain sensor mode as requested by user/driver. Has no significance if Windshield.Wiping.Mode is OFF/SLOW/MEDIUM/FAST 0 - wipers inactive. 1 - minimum intensity (lowest frequency/sensitivity, longest interval). 2/3/4/... - higher intensity (higher frequency/sensitivity, shorter interval). Maximum value supported is vehicle specific.
				Intensity uint8 `json:"intensity,omitempty"`
				// Wiper wear status. True = Worn, Replacement recommended or required. False = Not Worn.
				IsWipersWorn bool `json:"isWipersWorn,omitempty"`
				// Wiper mode requested by user/driver. INTERVAL indicates intermittent wiping, with fixed time interval between each wipe. RAIN_SENSOR indicates intermittent wiping based on rain intensity. Must be one of ["OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR"]
				Mode string `json:"mode,omitempty"`
				// Signals to control behavior of wipers in detail. By default VSS expects only one instance.
				System struct {
					// Actual position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific.
					ActualPosition float64 `json:"actualPosition,omitempty"`
					// Actual current used by wiper drive.
					DriveCurrent float64 `json:"driveCurrent,omitempty"`
					// Wiping frequency/speed, measured in cycles per minute. The signal concerns the actual speed of the wiper blades when moving. Intervals/pauses are excluded, i.e. the value corresponds to the number of cycles that would be completed in 1 minute if wiping permanently over default range.
					Frequency uint8 `json:"frequency,omitempty"`
					// Indicates if wiper movement is blocked. True = Movement blocked. False = Movement not blocked.
					IsBlocked bool `json:"isBlocked,omitempty"`
					// Indicates if current wipe movement is completed or near completion. True = Movement is completed or near completion. Changes to RequestedPosition will be executed first after reaching previous RequestedPosition, if it has not already been reached. False = Movement is not near completion. Any change to RequestedPosition will be executed immediately. Change of direction may not be allowed.
					IsEndingWipeCycle bool `json:"isEndingWipeCycle,omitempty"`
					// Indicates if wiper system is overheated. True = Wiper system overheated. False = Wiper system not overheated.
					IsOverheated bool `json:"isOverheated,omitempty"`
					// Indicates if a requested position has been reached. IsPositionReached refers to the previous position in case the TargetPosition is updated while IsEndingWipeCycle=True. True = Current or Previous TargetPosition reached. False = Position not (yet) reached, or wipers have moved away from the reached position.
					IsPositionReached bool `json:"isPositionReached,omitempty"`
					// Indicates system failure. True if wiping is disabled due to system failure.
					IsWiperError bool `json:"isWiperError,omitempty"`
					// Indicates wiper movement. True if wiper blades are moving. Change of direction shall be considered as IsWiping if wipers will continue to move directly after the change of direction.
					IsWiping bool `json:"isWiping,omitempty"`
					// Requested mode of wiper system. STOP_HOLD means that the wipers shall move to position given by TargetPosition and then hold the position. WIPE means that wipers shall move to the position given by TargetPosition and then hold the position if no new TargetPosition is requested. PLANT_MODE means that wiping is disabled. Exact behavior is vehicle specific. EMERGENCY_STOP means that wiping shall be immediately stopped without holding the position. Must be one of ["STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP"]
					Mode string `json:"mode,omitempty"`
					// Requested position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific. System behavior when receiving TargetPosition depends on Mode and IsEndingWipeCycle. Supported values are vehicle specific and might be dynamically corrected. If IsEndingWipeCycle=True then wipers will complete current movement before actuating new TargetPosition. If IsEndingWipeCycle=False then wipers will directly change destination if the TargetPosition is changed.
					TargetPosition float64 `json:"targetPosition,omitempty"`
				} `json:"system,omitempty"`
				// Wiper wear as percent. 0 = No Wear. 100 = Worn. Replacement required. Method for calculating or estimating wiper wear is vehicle specific. For windshields with multiple wipers the wear reported shall correspond to the most worn wiper.
				WiperWear uint8 `json:"wiperWear,omitempty"`
			} `json:"wiping,omitempty"`
		} `json:"front,omitempty"`
		// Windshield signals.
		Rear struct {
			// Windshield heater status. False - off, True - on.
			IsHeatingOn bool `json:"isHeatingOn,omitempty"`
			// Windshield washer fluid signals
			WasherFluid struct {
				// Low level indication for washer fluid. True = Level Low. False = Level OK.
				IsLevelLow bool `json:"isLevelLow,omitempty"`
				// Washer fluid level as a percent. 0 = Empty. 100 = Full.
				Level uint8 `json:"level,omitempty"`
			} `json:"washerFluid,omitempty"`
			// Windshield wiper signals.
			Wiping struct {
				// Relative intensity/sensitivity for interval and rain sensor mode as requested by user/driver. Has no significance if Windshield.Wiping.Mode is OFF/SLOW/MEDIUM/FAST 0 - wipers inactive. 1 - minimum intensity (lowest frequency/sensitivity, longest interval). 2/3/4/... - higher intensity (higher frequency/sensitivity, shorter interval). Maximum value supported is vehicle specific.
				Intensity uint8 `json:"intensity,omitempty"`
				// Wiper wear status. True = Worn, Replacement recommended or required. False = Not Worn.
				IsWipersWorn bool `json:"isWipersWorn,omitempty"`
				// Wiper mode requested by user/driver. INTERVAL indicates intermittent wiping, with fixed time interval between each wipe. RAIN_SENSOR indicates intermittent wiping based on rain intensity. Must be one of ["OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR"]
				Mode string `json:"mode,omitempty"`
				// Signals to control behavior of wipers in detail. By default VSS expects only one instance.
				System struct {
					// Actual position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific.
					ActualPosition float64 `json:"actualPosition,omitempty"`
					// Actual current used by wiper drive.
					DriveCurrent float64 `json:"driveCurrent,omitempty"`
					// Wiping frequency/speed, measured in cycles per minute. The signal concerns the actual speed of the wiper blades when moving. Intervals/pauses are excluded, i.e. the value corresponds to the number of cycles that would be completed in 1 minute if wiping permanently over default range.
					Frequency uint8 `json:"frequency,omitempty"`
					// Indicates if wiper movement is blocked. True = Movement blocked. False = Movement not blocked.
					IsBlocked bool `json:"isBlocked,omitempty"`
					// Indicates if current wipe movement is completed or near completion. True = Movement is completed or near completion. Changes to RequestedPosition will be executed first after reaching previous RequestedPosition, if it has not already been reached. False = Movement is not near completion. Any change to RequestedPosition will be executed immediately. Change of direction may not be allowed.
					IsEndingWipeCycle bool `json:"isEndingWipeCycle,omitempty"`
					// Indicates if wiper system is overheated. True = Wiper system overheated. False = Wiper system not overheated.
					IsOverheated bool `json:"isOverheated,omitempty"`
					// Indicates if a requested position has been reached. IsPositionReached refers to the previous position in case the TargetPosition is updated while IsEndingWipeCycle=True. True = Current or Previous TargetPosition reached. False = Position not (yet) reached, or wipers have moved away from the reached position.
					IsPositionReached bool `json:"isPositionReached,omitempty"`
					// Indicates system failure. True if wiping is disabled due to system failure.
					IsWiperError bool `json:"isWiperError,omitempty"`
					// Indicates wiper movement. True if wiper blades are moving. Change of direction shall be considered as IsWiping if wipers will continue to move directly after the change of direction.
					IsWiping bool `json:"isWiping,omitempty"`
					// Requested mode of wiper system. STOP_HOLD means that the wipers shall move to position given by TargetPosition and then hold the position. WIPE means that wipers shall move to the position given by TargetPosition and then hold the position if no new TargetPosition is requested. PLANT_MODE means that wiping is disabled. Exact behavior is vehicle specific. EMERGENCY_STOP means that wiping shall be immediately stopped without holding the position. Must be one of ["STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP"]
					Mode string `json:"mode,omitempty"`
					// Requested position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific. System behavior when receiving TargetPosition depends on Mode and IsEndingWipeCycle. Supported values are vehicle specific and might be dynamically corrected. If IsEndingWipeCycle=True then wipers will complete current movement before actuating new TargetPosition. If IsEndingWipeCycle=False then wipers will directly change destination if the TargetPosition is changed.
					TargetPosition float64 `json:"targetPosition,omitempty"`
				} `json:"system,omitempty"`
				// Wiper wear as percent. 0 = No Wear. 100 = Worn. Replacement required. Method for calculating or estimating wiper wear is vehicle specific. For windshields with multiple wipers the wear reported shall correspond to the most worn wiper.
				WiperWear uint8 `json:"wiperWear,omitempty"`
			} `json:"wiping,omitempty"`
		} `json:"rear,omitempty"`
	} `json:"windshield,omitempty"`
}

// All in-cabin components, including doors.
type Cabin struct {
	// Convertible roof.
	Convertible struct {
		// Roof status on convertible vehicles. Must be one of ["UNDEFINED", "CLOSED", "OPEN", "CLOSING", "OPENING", "STALLED"]
		Status string `json:"status,omitempty"`
	} `json:"convertible,omitempty"`
	// All doors, including windows and switches.
	Door struct {
		// All doors, including windows and switches.
		Row1 struct {
			// All doors, including windows and switches.
			Left struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"left,omitempty"`
			// All doors, including windows and switches.
			Right struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"right,omitempty"`
		} `json:"row1,omitempty"`
		// All doors, including windows and switches.
		Row2 struct {
			// All doors, including windows and switches.
			Left struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"left,omitempty"`
			// All doors, including windows and switches.
			Right struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch string `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"right,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"door,omitempty"`
	// Number of doors in vehicle.
	DoorCount uint8 `json:"doorCount,omitempty"`
	// The position of the driver seat in row 1.
	DriverPosition uint8 `json:"driverPosition,omitempty"`
	// Climate control
	HVAC struct {
		// Ambient air temperature inside the vehicle.
		AmbientAirTemperature float64 `json:"ambientAirTemperature,omitempty"`
		// Is Air conditioning active.
		IsAirConditioningActive bool `json:"isAirConditioningActive,omitempty"`
		// Is front defroster active.
		IsFrontDefrosterActive bool `json:"isFrontDefrosterActive,omitempty"`
		// Is rear defroster active.
		IsRearDefrosterActive bool `json:"isRearDefrosterActive,omitempty"`
		// Is recirculation active.
		IsRecirculationActive bool `json:"isRecirculationActive,omitempty"`
		// HVAC for single station in the vehicle
		Station struct {
			// HVAC for single station in the vehicle
			Row1 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row1,omitempty"`
			// HVAC for single station in the vehicle
			Row2 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row2,omitempty"`
			// HVAC for single station in the vehicle
			Row3 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row3,omitempty"`
			// HVAC for single station in the vehicle
			Row4 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution string `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row4,omitempty"`
		} `json:"station,omitempty"`
	} `json:"hVAC,omitempty"`
	// Infotainment system.
	Infotainment struct {
		// HMI related signals
		HMI struct {
			// ISO 639-1 standard language code for the current HMI
			CurrentLanguage string `json:"currentLanguage,omitempty"`
			// Date format used in the current HMI Must be one of ["YYYY_MM_DD", "DD_MM_YYYY", "MM_DD_YYYY", "YY_MM_DD", "DD_MM_YY", "MM_DD_YY"]
			DateFormat string `json:"dateFormat,omitempty"`
			// Current display theme Must be one of ["DAY", "NIGHT"]
			DayNightMode string `json:"dayNightMode,omitempty"`
			// Distance unit used in the current HMI Must be one of ["MILES", "KILOMETERS"]
			DistanceUnit string `json:"distanceUnit,omitempty"`
			// EV fuel economy unit used in the current HMI Must be one of ["MILES_PER_KILOWATT_HOUR", "KILOMETERS_PER_KILOWATT_HOUR", "KILOWATT_HOURS_PER_100_MILES", "KILOWATT_HOURS_PER_100_KILOMETERS", "WATT_HOURS_PER_MILE", "WATT_HOURS_PER_KILOMETER"]
			EVEconomyUnits string `json:"eVEconomyUnits,omitempty"`
			// Fuel economy unit used in the current HMI Must be one of ["MPG_UK", "MPG_US", "MILES_PER_LITER", "KILOMETERS_PER_LITER", "LITERS_PER_100_KILOMETERS"]
			FuelEconomyUnits string `json:"fuelEconomyUnits,omitempty"`
			// Temperature unit used in the current HMI Must be one of ["C", "F"]
			TemperatureUnit string `json:"temperatureUnit,omitempty"`
			// Time format used in the current HMI Must be one of ["HR_12", "HR_24"]
			TimeFormat string `json:"timeFormat,omitempty"`
			// Tire pressure unit used in the current HMI Must be one of ["‘PSI’", "‘KPA’", "’BAR’"]
			TirePressureUnit string `json:"tirePressureUnit,omitempty"`
		} `json:"hMI,omitempty"`
		// All Media actions
		Media struct {
			// Tells if the media was Must be one of ["UNKNOWN", "STOP", "PLAY", "FAST_FORWARD", "FAST_BACKWARD", "SKIP_FORWARD", "SKIP_BACKWARD"]
			Action string `json:"action,omitempty"`
			// URI of suggested media that was declined
			DeclinedURI string `json:"declinedURI,omitempty"`
			// Collection of signals updated in concert when a new media is played
			Played struct {
				// Name of album being played
				Album string `json:"album,omitempty"`
				// Name of artist being played
				Artist string `json:"artist,omitempty"`
				// Current playback rate of media being played.
				PlaybackRate float64 `json:"playbackRate,omitempty"`
				// Media selected for playback Must be one of ["UNKNOWN", "SIRIUS_XM", "AM", "FM", "DAB", "TV", "CD", "DVD", "AUX", "USB", "DISK", "BLUETOOTH", "INTERNET", "VOICE", "BEEP"]
				Source string `json:"source,omitempty"`
				// Name of track being played
				Track string `json:"track,omitempty"`
				// User Resource associated with the media
				URI string `json:"uRI,omitempty"`
			} `json:"played,omitempty"`
			// URI of suggested media that was selected
			SelectedURI string `json:"selectedURI,omitempty"`
			// Current Media Volume
			Volume uint8 `json:"volume,omitempty"`
		} `json:"media,omitempty"`
		// All navigation actions
		Navigation struct {
			// A navigation has been selected.
			DestinationSet struct {
				// Latitude of destination in WGS 84 geodetic coordinates.
				Latitude float64 `json:"latitude,omitempty"`
				// Longitude of destination in WGS 84 geodetic coordinates.
				Longitude float64 `json:"longitude,omitempty"`
			} `json:"destinationSet,omitempty"`
			// Navigation mute state that was selected. Must be one of ["MUTED", "ALERT_ONLY", "UNMUTED"]
			Mute string `json:"mute,omitempty"`
			// Current navigation volume
			Volume uint8 `json:"volume,omitempty"`
		} `json:"navigation,omitempty"`
		// All smartphone projection actions.
		SmartphoneProjection struct {
			// Projection activation info. Must be one of ["NONE", "ACTIVE", "INACTIVE"]
			Active string `json:"active,omitempty"`
			// Connectivity source selected for projection. Must be one of ["USB", "BLUETOOTH", "WIFI"]
			Source string `json:"source,omitempty"`
			// Supportable list for projection. Must be one of ["ANDROID_AUTO", "APPLE_CARPLAY", "MIRROR_LINK", "OTHER"]
			SupportedMode []string `json:"supportedMode,omitempty"`
		} `json:"smartphoneProjection,omitempty"`
	} `json:"infotainment,omitempty"`
	// Interior lights signals and sensors.
	Lights struct {
		// How much ambient light is detected in cabin. 0 = No ambient light. 100 = Full brightness
		AmbientLight uint8 `json:"ambientLight,omitempty"`
		// Is central dome light light on
		IsDomeOn bool `json:"isDomeOn,omitempty"`
		// Is glove box light on
		IsGloveBoxOn bool `json:"isGloveBoxOn,omitempty"`
		// Is trunk light light on
		IsTrunkOn bool `json:"isTrunkOn,omitempty"`
		// Intensity of the interior lights. 0 = Off. 100 = Full brightness.
		LightIntensity uint8 `json:"lightIntensity,omitempty"`
		// Spotlight for a specific area in the vehicle.
		Spotlight struct {
			// Spotlight for a specific area in the vehicle.
			Row1 struct {
				// Is light on the left side switched on
				IsLeftOn bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn bool `json:"isSharedOn,omitempty"`
			} `json:"row1,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row2 struct {
				// Is light on the left side switched on
				IsLeftOn bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn bool `json:"isSharedOn,omitempty"`
			} `json:"row2,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row3 struct {
				// Is light on the left side switched on
				IsLeftOn bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn bool `json:"isSharedOn,omitempty"`
			} `json:"row3,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row4 struct {
				// Is light on the left side switched on
				IsLeftOn bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn bool `json:"isSharedOn,omitempty"`
			} `json:"row4,omitempty"`
		} `json:"spotlight,omitempty"`
	} `json:"lights,omitempty"`
	// Rear window shade.
	RearShade struct {
		// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
		Position uint8 `json:"position,omitempty"`
		// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
		Switch string `json:"switch,omitempty"`
	} `json:"rearShade,omitempty"`
	// Rearview mirror.
	RearviewMirror struct {
		// Dimming level of rearview mirror. 0 = undimmed. 100 = fully dimmed.
		DimmingLevel uint8 `json:"dimmingLevel,omitempty"`
	} `json:"rearviewMirror,omitempty"`
	// All seats.
	Seat struct {
		// All seats.
		Row1 struct {
			// All seats.
			Pos1 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos1,omitempty"`
			// All seats.
			Pos2 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos2,omitempty"`
			// All seats.
			Pos3 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos3,omitempty"`
		} `json:"row1,omitempty"`
		// All seats.
		Row2 struct {
			// All seats.
			Pos1 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos1,omitempty"`
			// All seats.
			Pos2 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos2,omitempty"`
			// All seats.
			Pos3 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer string `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject string `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt float64 `json:"tilt,omitempty"`
			} `json:"pos3,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"seat,omitempty"`
	// Number of seats across each row from the front to the rear.
	SeatPosCount []uint8 `json:"seatPosCount,omitempty"`
	// Number of seat rows in vehicle.
	SeatRowCount uint8 `json:"seatRowCount,omitempty"`
	// Sun roof status.
	Sunroof struct {
		// Sunroof position. 0 = Fully closed 100 = Fully opened. -100 = Fully tilted.
		Position int8 `json:"position,omitempty"`
		// Sun roof shade status.
		Shade struct {
			// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
			Position uint8 `json:"position,omitempty"`
			// Switch controlling sliding action such as window, sunroof, or blind. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
			Switch string `json:"switch,omitempty"`
		} `json:"shade,omitempty"`
		// Switch controlling sliding action such as window, sunroof, or shade. Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", "TILT_UP", "TILT_DOWN"]
		Switch string `json:"switch,omitempty"`
	} `json:"sunroof,omitempty"`
}

// All data concerning steering, suspension, wheels, and brakes.
type Chassis struct {
	// Accelerator signals
	Accelerator struct {
		// Accelerator pedal position as percent. 0 = Not depressed. 100 = Fully depressed.
		PedalPosition uint8 `json:"pedalPosition,omitempty"`
	} `json:"accelerator,omitempty"`
	// Axle signals
	Axle struct {
		// Axle signals
		Row1 struct {
			// Aspect ratio between tire section height and tire section width, as per ETRTO / TRA standard.
			TireAspectRatio uint8 `json:"tireAspectRatio,omitempty"`
			// Outer diameter of tires, in inches, as per ETRTO / TRA standard.
			TireDiameter float64 `json:"tireDiameter,omitempty"`
			// Nominal section width of tires, in mm, as per ETRTO / TRA standard.
			TireWidth uint16 `json:"tireWidth,omitempty"`
			// Wheel signals for axle
			Wheel struct {
				// Wheel signals for axle
				Left struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"left,omitempty"`
				// Wheel signals for axle
				Right struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"right,omitempty"`
			} `json:"wheel,omitempty"`
			// Number of wheels on the axle
			WheelCount uint8 `json:"wheelCount,omitempty"`
			// Diameter of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelDiameter float64 `json:"wheelDiameter,omitempty"`
			// Width of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelWidth float64 `json:"wheelWidth,omitempty"`
		} `json:"row1,omitempty"`
		// Axle signals
		Row2 struct {
			// Aspect ratio between tire section height and tire section width, as per ETRTO / TRA standard.
			TireAspectRatio uint8 `json:"tireAspectRatio,omitempty"`
			// Outer diameter of tires, in inches, as per ETRTO / TRA standard.
			TireDiameter float64 `json:"tireDiameter,omitempty"`
			// Nominal section width of tires, in mm, as per ETRTO / TRA standard.
			TireWidth uint16 `json:"tireWidth,omitempty"`
			// Wheel signals for axle
			Wheel struct {
				// Wheel signals for axle
				Left struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"left,omitempty"`
				// Wheel signals for axle
				Right struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"right,omitempty"`
			} `json:"wheel,omitempty"`
			// Number of wheels on the axle
			WheelCount uint8 `json:"wheelCount,omitempty"`
			// Diameter of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelDiameter float64 `json:"wheelDiameter,omitempty"`
			// Width of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelWidth float64 `json:"wheelWidth,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"axle,omitempty"`
	// Number of axles on the vehicle
	AxleCount uint8 `json:"axleCount,omitempty"`
	// Brake system signals
	Brake struct {
		// Indicates if emergency braking initiated by driver is detected. True = Emergency braking detected. False = Emergency braking not detected.
		IsDriverEmergencyBrakingDetected bool `json:"isDriverEmergencyBrakingDetected,omitempty"`
		// Brake pedal position as percent. 0 = Not depressed. 100 = Fully depressed.
		PedalPosition uint8 `json:"pedalPosition,omitempty"`
	} `json:"brake,omitempty"`
	// Parking brake signals
	ParkingBrake struct {
		// Parking brake status. True = Parking Brake is Engaged. False = Parking Brake is not Engaged.
		IsEngaged bool `json:"isEngaged,omitempty"`
	} `json:"parkingBrake,omitempty"`
	// Steering wheel signals
	SteeringWheel struct {
		// Steering wheel angle. Positive = degrees to the left. Negative = degrees to the right.
		Angle int16 `json:"angle,omitempty"`
		// Steering wheel column extension from dashboard. 0 = Closest to dashboard. 100 = Furthest from dashboard.
		Extension uint8 `json:"extension,omitempty"`
		// Position of the steering wheel on the left or right side of the vehicle. Must be one of ["FRONT_LEFT", "FRONT_RIGHT"]
		Position string `json:"position,omitempty"`
		// Steering wheel column tilt. 0 = Lowest position. 100 = Highest position.
		Tilt uint8 `json:"tilt,omitempty"`
	} `json:"steeringWheel,omitempty"`
	// Overall wheel tracking, in mm.
	Track uint16 `json:"track,omitempty"`
	// Overall wheel base, in mm.
	Wheelbase uint16 `json:"wheelbase,omitempty"`
}

// Connectivity data.
type Connectivity struct {
	// Indicates if connectivity between vehicle and cloud is available. True = Connectivity is available. False = Connectivity is not available.
	IsConnectivityAvailable bool `json:"isConnectivityAvailable,omitempty"`
}

type StaticFeatures struct {
	// Attributes that identify a vehicle.
	VehicleIdentification struct {
		// The ACRISS Car Classification Code is a code used by many car rental companies.
		AcrissCode string `json:"acrissCode,omitempty"`
		// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
		BodyType string `json:"bodyType,omitempty"`
		// Vehicle brand or manufacturer.
		Brand string `json:"brand,omitempty"`
		// The date in ISO 8601 format of the first registration of the vehicle with the respective public authorities.
		DateVehicleFirstRegistered string `json:"dateVehicleFirstRegistered,omitempty"`
		// A textual description of known damages, both repaired and unrepaired.
		KnownVehicleDamages string `json:"knownVehicleDamages,omitempty"`
		// Indicates that the vehicle meets the respective emission standard.
		MeetsEmissionStandard string `json:"meetsEmissionStandard,omitempty"`
		// Vehicle model.
		Model string `json:"model,omitempty"`
		// The date in ISO 8601 format of production of the item, e.g. vehicle.
		ProductionDate string `json:"productionDate,omitempty"`
		// The date in ISO 8601 format of the item e.g. vehicle was purchased by the current owner.
		PurchaseDate string `json:"purchaseDate,omitempty"`
		// 17-character Vehicle Identification Number (VIN) as defined by ISO 3779.
		VIN string `json:"vIN,omitempty"`
		// A short text indicating the configuration of the vehicle, e.g. '5dr hatchback ST 2.5 MT 225 hp' or 'limited edition'.
		VehicleConfiguration string `json:"vehicleConfiguration,omitempty"`
		// The color or color combination of the interior of the vehicle.
		VehicleInteriorColor string `json:"vehicleInteriorColor,omitempty"`
		// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.).
		VehicleInteriorType string `json:"vehicleInteriorType,omitempty"`
		// The release date in ISO 8601 format of a vehicle model (often used to differentiate versions of the same make and model).
		VehicleModelDate string `json:"vehicleModelDate,omitempty"`
		// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
		VehicleSeatingCapacity uint16 `json:"vehicleSeatingCapacity,omitempty"`
		// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school.
		VehicleSpecialUsage string `json:"vehicleSpecialUsage,omitempty"`
		// 3-character World Manufacturer Identification (WMI) as defined by ISO 3780.
		WMI string `json:"wMI,omitempty"`
		// Model year of the vehicle.
		Year uint16 `json:"year,omitempty"`
	} `json:"vehicleIdentification,omitempty"`
	// Overall vehicle width.
	Width uint16 `json:"width,omitempty"`
	// Maximum vertical weight on the tow ball of a trailer.
	MaxTowBallWeight uint16 `json:"maxTowBallWeight,omitempty"`
	// Maximum weight of trailer.
	MaxTowWeight uint16 `json:"maxTowWeight,omitempty"`
	// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
	CargoVolume float64 `json:"cargoVolume,omitempty"`
	// Vehicle curb weight, including all liquids and full tank of fuel, but no cargo or passengers.
	CurbWeight uint16 `json:"curbWeight,omitempty"`
	// The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.
	RoofLoad int16 `json:"roofLoad,omitempty"`
}

// Signals related to low voltage battery.
type LowVoltageBattery struct {
	// Current current flowing in/out of the low voltage battery. Positive = Current flowing in to battery, e.g. during charging or driving. Negative = Current flowing out of battery, e.g. when using the battery to start a combustion engine.
	CurrentCurrent float64 `json:"currentCurrent,omitempty"`
	// Current Voltage of the low voltage battery.
	CurrentVoltage float64 `json:"currentVoltage,omitempty"`
	// Nominal capacity of the low voltage battery.
	NominalCapacity uint16 `json:"nominalCapacity,omitempty"`
	// Nominal Voltage of the battery.
	NominalVoltage uint16 `json:"nominalVoltage,omitempty"`
}

// Powertrain data for battery management, etc.
type Powertrain struct {
	// The accumulated energy from regenerative braking over lifetime.
	AccumulatedBrakingEnergy float64 `json:"accumulatedBrakingEnergy,omitempty"`
	// Engine-specific data, stopping at the bell housing.
	CombustionEngine struct {
		// Type of aspiration (natural, turbocharger, supercharger etc). Must be one of ["UNKNOWN", "NATURAL", "SUPERCHARGER", "TURBOCHARGER"]
		AspirationType string `json:"aspirationType,omitempty"`
		// Bore in millimetres.
		Bore float64 `json:"bore,omitempty"`
		// Engine compression ratio, specified in the format 'X:1', e.g. '9.2:1'.
		CompressionRatio string `json:"compressionRatio,omitempty"`
		// Engine configuration. Must be one of ["UNKNOWN", "STRAIGHT", "V", "BOXER", "W", "ROTARY", "RADIAL", "SQUARE", "H", "U", "OPPOSED", "X"]
		Configuration string `json:"configuration,omitempty"`
		// Signals related to Diesel Exhaust Fluid (DEF). DEF is called AUS32 in ISO 22241.
		DieselExhaustFluid struct {
			// Capacity in liters of the Diesel Exhaust Fluid Tank.
			Capacity float64 `json:"capacity,omitempty"`
			// Indicates if the Diesel Exhaust Fluid level is low. True if level is low. Definition of low is vehicle dependent.
			IsLevelLow bool `json:"isLevelLow,omitempty"`
			// Level of the Diesel Exhaust Fluid tank as percent of capacity. 0 = empty. 100 = full.
			Level uint8 `json:"level,omitempty"`
			// Remaining range in meters of the Diesel Exhaust Fluid present in the vehicle.
			Range uint32 `json:"range,omitempty"`
		} `json:"dieselExhaustFluid,omitempty"`
		// Diesel Particulate Filter signals.
		DieselParticulateFilter struct {
			// Delta Pressure of Diesel Particulate Filter.
			DeltaPressure float64 `json:"deltaPressure,omitempty"`
			// Inlet temperature of Diesel Particulate Filter.
			InletTemperature float64 `json:"inletTemperature,omitempty"`
			// Outlet temperature of Diesel Particulate Filter.
			OutletTemperature float64 `json:"outletTemperature,omitempty"`
		} `json:"dieselParticulateFilter,omitempty"`
		// Displacement in cubic centimetres.
		Displacement uint16 `json:"displacement,omitempty"`
		// Engine coolant temperature.
		ECT int16 `json:"eCT,omitempty"`
		// Engine oil pressure.
		EOP uint16 `json:"eOP,omitempty"`
		// Engine oil temperature.
		EOT int16 `json:"eOT,omitempty"`
		// Engine code designation, as specified by vehicle manufacturer.
		EngineCode string `json:"engineCode,omitempty"`
		// Engine coolant capacity in liters.
		EngineCoolantCapacity float64 `json:"engineCoolantCapacity,omitempty"`
		// Accumulated time during engine lifetime with 'engine speed (rpm) > 0'.
		EngineHours float64 `json:"engineHours,omitempty"`
		// Engine oil capacity in liters.
		EngineOilCapacity float64 `json:"engineOilCapacity,omitempty"`
		// Engine oil level. Must be one of ["CRITICALLY_LOW", "LOW", "NORMAL", "HIGH", "CRITICALLY_HIGH"]
		EngineOilLevel string `json:"engineOilLevel,omitempty"`
		// Accumulated idling time during engine lifetime. Definition of idling is not standardized.
		IdleHours float64 `json:"idleHours,omitempty"`
		// Engine Running. True if engine is rotating (Speed > 0).
		IsRunning bool `json:"isRunning,omitempty"`
		// Grams of air drawn into engine per second.
		MAF uint16 `json:"mAF,omitempty"`
		// Manifold absolute pressure possibly boosted using forced induction.
		MAP uint16 `json:"mAP,omitempty"`
		// Peak power, in kilowatts, that engine can generate.
		MaxPower uint16 `json:"maxPower,omitempty"`
		// Peak torque, in newton meter, that the engine can generate.
		MaxTorque uint16 `json:"maxTorque,omitempty"`
		// Number of cylinders.
		NumberOfCylinders uint16 `json:"numberOfCylinders,omitempty"`
		// Number of valves per cylinder.
		NumberOfValvesPerCylinder uint16 `json:"numberOfValvesPerCylinder,omitempty"`
		// Remaining engine oil life in seconds. Negative values can be used to indicate that lifetime has been exceeded.
		OilLifeRemaining int32 `json:"oilLifeRemaining,omitempty"`
		// Current engine power output. Shall be reported as 0 during engine breaking.
		Power uint16 `json:"power,omitempty"`
		// Engine speed measured as rotations per minute.
		Speed uint16 `json:"speed,omitempty"`
		// Stroke length in millimetres.
		StrokeLength float64 `json:"strokeLength,omitempty"`
		// Current throttle position.
		TPS uint8 `json:"tPS,omitempty"`
		// Current engine torque. Shall be reported as 0 during engine breaking.
		Torque uint16 `json:"torque,omitempty"`
	} `json:"combustionEngine,omitempty"`
	// Electric Motor specific data.
	ElectricMotor struct {
		// Motor coolant temperature (if applicable).
		CoolantTemperature int16 `json:"coolantTemperature,omitempty"`
		// Engine code designation, as specified by vehicle manufacturer.
		EngineCode string `json:"engineCode,omitempty"`
		// Peak power, in kilowatts, that motor(s) can generate.
		MaxPower uint16 `json:"maxPower,omitempty"`
		// Peak regen/brake power, in kilowatts, that motor(s) can generate.
		MaxRegenPower uint16 `json:"maxRegenPower,omitempty"`
		// Peak regen/brake torque, in newton meter, that the motor(s) can generate.
		MaxRegenTorque uint16 `json:"maxRegenTorque,omitempty"`
		// Peak power, in newton meter, that the motor(s) can generate.
		MaxTorque uint16 `json:"maxTorque,omitempty"`
		// Current motor power output. Negative values indicate regen mode.
		Power int16 `json:"power,omitempty"`
		// Motor rotational speed measured as rotations per minute. Negative values indicate reverse driving mode.
		Speed int32 `json:"speed,omitempty"`
		// Motor temperature.
		Temperature int16 `json:"temperature,omitempty"`
		// Current motor torque. Negative values indicate regen mode.
		Torque int16 `json:"torque,omitempty"`
	} `json:"electricMotor,omitempty"`
	// Fuel system data.
	FuelSystem struct {
		// Average consumption in liters per 100 km.
		AverageConsumption float64 `json:"averageConsumption,omitempty"`
		// Fuel amount in liters consumed since start of current trip.
		ConsumptionSinceStart float64 `json:"consumptionSinceStart,omitempty"`
		// Defines the hybrid type of the vehicle. Must be one of ["UNKNOWN", "NOT_APPLICABLE", "STOP_START", "BELT_ISG", "CIMG", "PHEV"]
		HybridType string `json:"hybridType,omitempty"`
		// Current consumption in liters per 100 km.
		InstantConsumption float64 `json:"instantConsumption,omitempty"`
		// Indicates whether eco start stop is currently enabled.
		IsEngineStopStartEnabled bool `json:"isEngineStopStartEnabled,omitempty"`
		// Indicates that the fuel level is low (e.g. <50km range).
		IsFuelLevelLow bool `json:"isFuelLevelLow,omitempty"`
		// Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
		Level uint8 `json:"level,omitempty"`
		// Remaining range in meters using only liquid fuel.
		Range uint32 `json:"range,omitempty"`
		// Detailed information on fuels supported by the vehicle. Identifiers originating from DIN EN 16942:2021-08, appendix B, with additional suffix for octane (RON) where relevant. Must be one of ["E5_95", "E5_98", "E10_95", "E10_98", "E85", "B7", "B10", "B20", "B30", "B100", "XTL", "LPG", "CNG", "LNG", "H2", "OTHER"]
		SupportedFuel []string `json:"supportedFuel,omitempty"`
		// High level information of fuel types supported Must be one of ["GASOLINE", "DIESEL", "E85", "LPG", "CNG", "LNG", "H2", "OTHER"]
		SupportedFuelTypes []string `json:"supportedFuelTypes,omitempty"`
		// Capacity of the fuel tank in liters.
		TankCapacity float64 `json:"tankCapacity,omitempty"`
		// Time in seconds elapsed since start of current trip.
		TimeSinceStart uint32 `json:"timeSinceStart,omitempty"`
	} `json:"fuelSystem,omitempty"`
	// Remaining range in meters using all energy sources available in the vehicle.
	Range uint32 `json:"range,omitempty"`
	// Battery Management data.
	TractionBattery struct {
		// The accumulated energy delivered to the battery during charging over lifetime of the battery.
		AccumulatedChargedEnergy float64 `json:"accumulatedChargedEnergy,omitempty"`
		// The accumulated charge throughput delivered to the battery during charging over lifetime of the battery.
		AccumulatedChargedThroughput float64 `json:"accumulatedChargedThroughput,omitempty"`
		// The accumulated energy leaving HV battery for propulsion and auxiliary loads over lifetime of the battery.
		AccumulatedConsumedEnergy float64 `json:"accumulatedConsumedEnergy,omitempty"`
		// The accumulated charge throughput leaving HV battery for propulsion and auxiliary loads over lifetime of the battery.
		AccumulatedConsumedThroughput float64 `json:"accumulatedConsumedThroughput,omitempty"`
		// Properties related to battery charging.
		Charging struct {
			// Current charging current.
			ChargeCurrent struct {
				// Current DC charging current at inlet. Negative if returning energy to grid.
				DC float64 `json:"dC,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 1. Negative if returning energy to grid.
				Phase1 float64 `json:"phase1,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 2. Negative if returning energy to grid.
				Phase2 float64 `json:"phase2,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 3. Negative if returning energy to grid.
				Phase3 float64 `json:"phase3,omitempty"`
			} `json:"chargeCurrent,omitempty"`
			// Target charge limit (state of charge) for battery.
			ChargeLimit uint8 `json:"chargeLimit,omitempty"`
			// Type of charge plug (charging inlet) available on the vehicle. IEC types refer to IEC 62196,  GBT refers to  GB/T 20234. Must be one of ["IEC_TYPE_1_AC", "IEC_TYPE_2_AC", "IEC_TYPE_3_AC", "IEC_TYPE_4_DC", "IEC_TYPE_1_CCS_DC", "IEC_TYPE_2_CCS_DC", "TESLA_ROADSTER", "TESLA_HPWC", "TESLA_SUPERCHARGER", "GBT_AC", "GBT_DC", "OTHER"]
			ChargePlugType []string `json:"chargePlugType,omitempty"`
			// Status of the charge port cover, can potentially be controlled manually. Must be one of ["OPEN", "CLOSED"]
			ChargePortFlap string `json:"chargePortFlap,omitempty"`
			// Current charging rate, as in kilometers of range added per hour.
			ChargeRate float64 `json:"chargeRate,omitempty"`
			// Current charging voltage, as measured at the charging inlet.
			ChargeVoltage struct {
				// Current DC charging voltage at charging inlet.
				DC float64 `json:"dC,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 1.
				Phase1 float64 `json:"phase1,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 2.
				Phase2 float64 `json:"phase2,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 3.
				Phase3 float64 `json:"phase3,omitempty"`
			} `json:"chargeVoltage,omitempty"`
			// True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
			IsCharging bool `json:"isCharging,omitempty"`
			// Indicates if a charging cable is physically connected to the vehicle or not.
			IsChargingCableConnected bool `json:"isChargingCableConnected,omitempty"`
			// Is charging cable locked to prevent removal.
			IsChargingCableLocked bool `json:"isChargingCableLocked,omitempty"`
			// True if discharging (vehicle to grid) is ongoing. Discharging is considered to be ongoing if energy is flowing from vehicle to charger/grid.
			IsDischarging bool `json:"isDischarging,omitempty"`
			// Maximum charging current that can be accepted by the system, as measured at the charging inlet.
			MaximumChargingCurrent struct {
				// Maximum DC charging current at inlet that can be accepted by the system.
				DC float64 `json:"dC,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 1 that can be accepted by the system.
				Phase1 float64 `json:"phase1,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 2 that can be accepted by the system.
				Phase2 float64 `json:"phase2,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 3 that can be accepted by the system.
				Phase3 float64 `json:"phase3,omitempty"`
			} `json:"maximumChargingCurrent,omitempty"`
			// Control of the charge process. MANUAL means manually initiated (plug-in event, companion app, etc). TIMER means timer-based. GRID means grid-controlled (eg ISO 15118). PROFILE means controlled by profile download to vehicle. Must be one of ["MANUAL", "TIMER", "GRID", "PROFILE"]
			Mode string `json:"mode,omitempty"`
			// Electrical energy lost by power dissipation to heat inside the AC/DC converter.
			PowerLoss float64 `json:"powerLoss,omitempty"`
			// Start or stop the charging process. Must be one of ["START", "STOP"]
			StartStopCharging string `json:"startStopCharging,omitempty"`
			// Current temperature of AC/DC converter converting grid voltage to battery voltage.
			Temperature float64 `json:"temperature,omitempty"`
			// The time needed for the current charging process to reach Charging.ChargeLimit. 0 if charging is complete or no charging process is active or planned.
			TimeToComplete uint32 `json:"timeToComplete,omitempty"`
			// Properties related to timing of battery charging sessions.
			Timer struct {
				// Defines timer mode for charging: INACTIVE - no timer set, charging may start as soon as battery is connected to a charger. START_TIME - charging shall start at Charging.Timer.Time. END_TIME - charging shall be finished (reach Charging.ChargeLimit) at Charging.Timer.Time. When charging is completed the vehicle shall change mode to 'inactive' or set a new Charging.Timer.Time. Charging shall start immediately if mode is 'starttime' or 'endtime' and Charging.Timer.Time is a time in the past. Must be one of ["INACTIVE", "START_TIME", "END_TIME"]
				Mode string `json:"mode,omitempty"`
				// Time for next charging-related action, formatted according to ISO 8601 with UTC time zone. Value has no significance if Charging.Timer.Mode is 'inactive'.
				Time string `json:"time,omitempty"`
			} `json:"timer,omitempty"`
		} `json:"charging,omitempty"`
		// Current current flowing in/out of battery. Positive = Current flowing in to battery, e.g. during charging. Negative = Current flowing out of battery, e.g. during driving.
		CurrentCurrent float64 `json:"currentCurrent,omitempty"`
		// Current electrical energy flowing in/out of battery. Positive = Energy flowing in to battery, e.g. during charging. Negative = Energy flowing out of battery, e.g. during driving.
		CurrentPower float64 `json:"currentPower,omitempty"`
		// Current Voltage of the battery.
		CurrentVoltage float64 `json:"currentVoltage,omitempty"`
		// Properties related to DC/DC converter converting high voltage (from high voltage battery) to vehicle low voltage (supply voltage, typically 12 Volts).
		DCDC struct {
			// Electrical energy lost by power dissipation to heat inside DC/DC converter.
			PowerLoss float64 `json:"powerLoss,omitempty"`
			// Current temperature of DC/DC converter converting battery high voltage to vehicle low voltage (typically 12 Volts).
			Temperature float64 `json:"temperature,omitempty"`
		} `json:"dCDC,omitempty"`
		// Gross capacity of the battery.
		GrossCapacity uint16 `json:"grossCapacity,omitempty"`
		// Battery Identification Number as assigned by OEM.
		Id string `json:"id,omitempty"`
		// Indicating if the ground (negative terminator) of the traction battery is connected to the powertrain.
		IsGroundConnected bool `json:"isGroundConnected,omitempty"`
		// Indicating if the power (positive terminator) of the traction battery is connected to the powertrain.
		IsPowerConnected bool `json:"isPowerConnected,omitempty"`
		// Max allowed voltage of the battery, e.g. during charging.
		MaxVoltage uint16 `json:"maxVoltage,omitempty"`
		// Total net capacity of the battery considering aging.
		NetCapacity uint16 `json:"netCapacity,omitempty"`
		// Nominal Voltage of the battery.
		NominalVoltage uint16 `json:"nominalVoltage,omitempty"`
		// Electrical energy lost by power dissipation to heat inside the battery.
		PowerLoss float64 `json:"powerLoss,omitempty"`
		// Production date of battery in ISO8601 format, e.g. YYYY-MM-DD.
		ProductionDate string `json:"productionDate,omitempty"`
		// Remaining range in meters using only battery.
		Range uint32 `json:"range,omitempty"`
		// Information on the state of charge of the vehicle's high voltage battery.
		StateOfCharge struct {
			// Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
			Current float64 `json:"current,omitempty"`
			// State of charge displayed to the customer.
			Displayed float64 `json:"displayed,omitempty"`
		} `json:"stateOfCharge,omitempty"`
		// Calculated battery state of health at standard conditions.
		StateOfHealth float64 `json:"stateOfHealth,omitempty"`
		// Temperature Information for the battery pack.
		Temperature struct {
			// Current average temperature of the battery cells.
			Average float64 `json:"average,omitempty"`
			// Current maximum temperature of the battery cells, i.e. temperature of the hottest cell.
			Max float64 `json:"max,omitempty"`
			// Current minimum temperature of the battery cells, i.e. temperature of the coldest cell.
			Min float64 `json:"min,omitempty"`
		} `json:"temperature,omitempty"`
	} `json:"tractionBattery,omitempty"`
	// Transmission-specific data, stopping at the drive shafts.
	Transmission struct {
		// Clutch engagement. 0% = Clutch fully disengaged. 100% = Clutch fully engaged.
		ClutchEngagement float64 `json:"clutchEngagement,omitempty"`
		// Clutch wear as a percent. 0 = no wear. 100 = worn.
		ClutchWear uint8 `json:"clutchWear,omitempty"`
		// The current gear. 0=Neutral, 1/2/..=Forward, -1/-2/..=Reverse.
		CurrentGear int8 `json:"currentGear,omitempty"`
		// Front Diff Lock engagement. 0% = Diff lock fully disengaged. 100% = Diff lock fully engaged.
		DiffLockFrontEngagement float64 `json:"diffLockFrontEngagement,omitempty"`
		// Rear Diff Lock engagement. 0% = Diff lock fully disengaged. 100% = Diff lock fully engaged.
		DiffLockRearEngagement float64 `json:"diffLockRearEngagement,omitempty"`
		// Drive type. Must be one of ["UNKNOWN", "FORWARD_WHEEL_DRIVE", "REAR_WHEEL_DRIVE", "ALL_WHEEL_DRIVE"]
		DriveType string `json:"driveType,omitempty"`
		// Is the gearbox in automatic or manual (paddle) mode. Must be one of ["MANUAL", "AUTOMATIC"]
		GearChangeMode string `json:"gearChangeMode,omitempty"`
		// Number of forward gears in the transmission. -1 = CVT.
		GearCount int8 `json:"gearCount,omitempty"`
		// Is electrical powertrain mechanically connected/engaged to the drivetrain or not. False = Disconnected/Disengaged. True = Connected/Engaged.
		IsElectricalPowertrainEngaged bool `json:"isElectricalPowertrainEngaged,omitempty"`
		// Is gearbox in low range mode or not. False = Normal/High range engaged. True = Low range engaged.
		IsLowRangeEngaged bool `json:"isLowRangeEngaged,omitempty"`
		// Is the transmission park lock engaged or not. False = Disengaged. True = Engaged.
		IsParkLockEngaged bool `json:"isParkLockEngaged,omitempty"`
		// Current gearbox performance mode. Must be one of ["NORMAL", "SPORT", "ECONOMY", "SNOW", "RAIN"]
		PerformanceMode string `json:"performanceMode,omitempty"`
		// The selected gear. 0=Neutral, 1/2/..=Forward, -1/-2/..=Reverse, 126=Park, 127=Drive.
		SelectedGear int8 `json:"selectedGear,omitempty"`
		// The current gearbox temperature.
		Temperature int16 `json:"temperature,omitempty"`
		// Torque distribution between front and rear axle in percent. -100% = Full torque to front axle, 0% = 50:50 Front/Rear, 100% = Full torque to rear axle.
		TorqueDistribution float64 `json:"torqueDistribution,omitempty"`
		// Odometer reading, total distance travelled during the lifetime of the transmission.
		TravelledDistance float64 `json:"travelledDistance,omitempty"`
		// Transmission type. Must be one of ["UNKNOWN", "SEQUENTIAL", "H", "AUTOMATIC", "DSG", "CVT"]
		Type string `json:"type,omitempty"`
	} `json:"transmission,omitempty"`
	// Defines the powertrain type of the vehicle. Must be one of ["COMBUSTION", "HYBRID", "ELECTRIC"]
	Type string `json:"type,omitempty"`
}

// Service data.
type Service struct {
	// Remaining distance to service (of any kind). Negative values indicate service overdue.
	DistanceToService float64 `json:"distanceToService,omitempty"`
	// Indicates if vehicle needs service (of any kind). True = Service needed now or in the near future. False = No known need for service.
	IsServiceDue bool `json:"isServiceDue,omitempty"`
	// Remaining time to service (of any kind). Negative values indicate service overdue.
	TimeToService int32 `json:"timeToService,omitempty"`
}

// Trailer signals.
type Trailer struct {
	// Signal indicating if trailer is connected or not.
	IsConnected bool `json:"isConnected,omitempty"`
}

// OBD data.
type OBD struct {
	// PID 43 - Absolute load value
	AbsoluteLoad float64 `json:"absoluteLoad,omitempty"`
	// PID 49 - Accelerator pedal position D
	AcceleratorPositionD float64 `json:"acceleratorPositionD,omitempty"`
	// PID 4A - Accelerator pedal position E
	AcceleratorPositionE float64 `json:"acceleratorPositionE,omitempty"`
	// PID 4B - Accelerator pedal position F
	AcceleratorPositionF float64 `json:"acceleratorPositionF,omitempty"`
	// PID 12 - Secondary air status
	AirStatus string `json:"airStatus,omitempty"`
	// PID 46 - Ambient air temperature
	AmbientAirTemperature float64 `json:"ambientAirTemperature,omitempty"`
	// PID 33 - Barometric pressure
	BarometricPressure float64 `json:"barometricPressure,omitempty"`
	// Catalyst signals
	Catalyst struct {
		// Catalyst bank 1 signals
		Bank1 struct {
			// PID 3C - Catalyst temperature from bank 1, sensor 1
			Temperature1 float64 `json:"temperature1,omitempty"`
			// PID 3E - Catalyst temperature from bank 1, sensor 2
			Temperature2 float64 `json:"temperature2,omitempty"`
		} `json:"bank1,omitempty"`
		// Catalyst bank 2 signals
		Bank2 struct {
			// PID 3D - Catalyst temperature from bank 2, sensor 1
			Temperature1 float64 `json:"temperature1,omitempty"`
			// PID 3F - Catalyst temperature from bank 2, sensor 2
			Temperature2 float64 `json:"temperature2,omitempty"`
		} `json:"bank2,omitempty"`
	} `json:"catalyst,omitempty"`
	// PID 2C - Commanded exhaust gas recirculation (EGR)
	CommandedEGR float64 `json:"commandedEGR,omitempty"`
	// PID 2E - Commanded evaporative purge (EVAP) valve
	CommandedEVAP float64 `json:"commandedEVAP,omitempty"`
	// PID 44 - Commanded equivalence ratio
	CommandedEquivalenceRatio float64 `json:"commandedEquivalenceRatio,omitempty"`
	// PID 42 - Control module voltage
	ControlModuleVoltage float64 `json:"controlModuleVoltage,omitempty"`
	// PID 05 - Coolant temperature
	CoolantTemperature float64 `json:"coolantTemperature,omitempty"`
	// List of currently active DTCs formatted according OBD II (SAE-J2012DA_201812) standard ([P|C|B|U]XXXXX )
	DTCList []string `json:"dTCList,omitempty"`
	// PID 31 - Distance traveled since codes cleared
	DistanceSinceDTCClear float64 `json:"distanceSinceDTCClear,omitempty"`
	// PID 21 - Distance traveled with MIL on
	DistanceWithMIL float64 `json:"distanceWithMIL,omitempty"`
	// PID 41 - OBD status for the current drive cycle
	DriveCycleStatus struct {
		// Number of sensor Trouble Codes (DTC)
		DTCCount uint8 `json:"dTCCount,omitempty"`
		// Type of the ignition for ICE - spark = spark plug ignition, compression = self-igniting (Diesel engines) Must be one of ["SPARK", "COMPRESSION"]
		IgnitionType string `json:"ignitionType,omitempty"`
		// Malfunction Indicator Light (MIL) - False = Off, True = On
		IsMILOn bool `json:"isMILOn,omitempty"`
	} `json:"driveCycleStatus,omitempty"`
	// PID 2D - Exhaust gas recirculation (EGR) error
	EGRError float64 `json:"eGRError,omitempty"`
	// PID 32 - Evaporative purge (EVAP) system pressure
	EVAPVaporPressure float64 `json:"eVAPVaporPressure,omitempty"`
	// PID 53 - Absolute evaporative purge (EVAP) system pressure
	EVAPVaporPressureAbsolute float64 `json:"eVAPVaporPressureAbsolute,omitempty"`
	// PID 54 - Alternate evaporative purge (EVAP) system pressure
	EVAPVaporPressureAlternate float64 `json:"eVAPVaporPressureAlternate,omitempty"`
	// PID 04 - Engine load in percent - 0 = no load, 100 = full load
	EngineLoad float64 `json:"engineLoad,omitempty"`
	// PID 0C - Engine speed measured as rotations per minute
	EngineSpeed float64 `json:"engineSpeed,omitempty"`
	// PID 52 - Percentage of ethanol in the fuel
	EthanolPercent float64 `json:"ethanolPercent,omitempty"`
	// PID 02 - DTC that triggered the freeze frame
	FreezeDTC string `json:"freezeDTC,omitempty"`
	// PID 5D - Fuel injection timing
	FuelInjectionTiming float64 `json:"fuelInjectionTiming,omitempty"`
	// PID 2F - Fuel level in the fuel tank
	FuelLevel float64 `json:"fuelLevel,omitempty"`
	// PID 0A - Fuel pressure
	FuelPressure float64 `json:"fuelPressure,omitempty"`
	// PID 59 - Absolute fuel rail pressure
	FuelRailPressureAbsolute float64 `json:"fuelRailPressureAbsolute,omitempty"`
	// PID 23 - Fuel rail pressure direct inject
	FuelRailPressureDirect float64 `json:"fuelRailPressureDirect,omitempty"`
	// PID 22 - Fuel rail pressure relative to vacuum
	FuelRailPressureVac float64 `json:"fuelRailPressureVac,omitempty"`
	// PID 5E - Engine fuel rate
	FuelRate float64 `json:"fuelRate,omitempty"`
	// PID 03 - Fuel status
	FuelStatus string `json:"fuelStatus,omitempty"`
	// PID 51 - Fuel type
	FuelType string `json:"fuelType,omitempty"`
	// PID 5B - Remaining life of hybrid battery
	HybridBatteryRemaining float64 `json:"hybridBatteryRemaining,omitempty"`
	// PID 0F - Intake temperature
	IntakeTemp float64 `json:"intakeTemp,omitempty"`
	// PID 1E - Auxiliary input status (power take off)
	IsPTOActive bool `json:"isPTOActive,omitempty"`
	// PID 07 - Long Term (learned) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
	LongTermFuelTrim1 float64 `json:"longTermFuelTrim1,omitempty"`
	// PID 09 - Long Term (learned) Fuel Trim - Bank 2 - negative percent leaner, positive percent richer
	LongTermFuelTrim2 float64 `json:"longTermFuelTrim2,omitempty"`
	// PID 56 (byte A) - Long term secondary O2 trim - Bank 1
	LongTermO2Trim1 float64 `json:"longTermO2Trim1,omitempty"`
	// PID 58 (byte A) - Long term secondary O2 trim - Bank 2
	LongTermO2Trim2 float64 `json:"longTermO2Trim2,omitempty"`
	// PID 56 (byte B) - Long term secondary O2 trim - Bank 3
	LongTermO2Trim3 float64 `json:"longTermO2Trim3,omitempty"`
	// PID 58 (byte B) - Long term secondary O2 trim - Bank 4
	LongTermO2Trim4 float64 `json:"longTermO2Trim4,omitempty"`
	// PID 10 - Grams of air drawn into engine per second
	MAF float64 `json:"mAF,omitempty"`
	// PID 0B - Intake manifold pressure
	MAP float64 `json:"mAP,omitempty"`
	// PID 50 - Maximum flow for mass air flow sensor
	MaxMAF float64 `json:"maxMAF,omitempty"`
	// Oxygen sensors (PID 14 - PID 1B)
	O2 struct {
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor1 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor1,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor2 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor2,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor3 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor3,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor4 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor4,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor5 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor5,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor6 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor6,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor7 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor7,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor8 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor8,omitempty"`
	} `json:"o2,omitempty"`
	// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
	O2WR struct {
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor1 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor1,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor2 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor2,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor3 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor3,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor4 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor4,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor5 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor5,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor6 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor6,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor7 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor7,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor8 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage float64 `json:"voltage,omitempty"`
		} `json:"sensor8,omitempty"`
	} `json:"o2WR,omitempty"`
	// PID 1C - OBD standards this vehicle conforms to
	OBDStandards uint8 `json:"oBDStandards,omitempty"`
	// PID 5C - Engine oil temperature
	OilTemperature float64 `json:"oilTemperature,omitempty"`
	// PID 13 - Presence of oxygen sensors in 2 banks. [A0..A3] == Bank 1, Sensors 1-4. [A4..A7] == Bank 2, Sensors 1-4
	OxygenSensorsIn2Banks uint8 `json:"oxygenSensorsIn2Banks,omitempty"`
	// PID 1D - Presence of oxygen sensors in 4 banks. Similar to PID 13, but [A0..A7] == [B1S1, B1S2, B2S1, B2S2, B3S1, B3S2, B4S1, B4S2]
	OxygenSensorsIn4Banks uint8 `json:"oxygenSensorsIn4Banks,omitempty"`
	// PID 00 - Bit array of the supported pids 01 to 20
	PidsA uint32 `json:"pidsA,omitempty"`
	// PID 20 - Bit array of the supported pids 21 to 40
	PidsB uint32 `json:"pidsB,omitempty"`
	// PID 40 - Bit array of the supported pids 41 to 60
	PidsC uint32 `json:"pidsC,omitempty"`
	// PID 5A - Relative accelerator pedal position
	RelativeAcceleratorPosition float64 `json:"relativeAcceleratorPosition,omitempty"`
	// PID 45 - Relative throttle position
	RelativeThrottlePosition float64 `json:"relativeThrottlePosition,omitempty"`
	// PID 1F - Engine run time
	RunTime float64 `json:"runTime,omitempty"`
	// PID 4D - Run time with MIL on
	RunTimeMIL float64 `json:"runTimeMIL,omitempty"`
	// PID 06 - Short Term (immediate) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
	ShortTermFuelTrim1 float64 `json:"shortTermFuelTrim1,omitempty"`
	// PID 08 - Short Term (immediate) Fuel Trim - Bank 2 - negative percent leaner, positive percent richer
	ShortTermFuelTrim2 float64 `json:"shortTermFuelTrim2,omitempty"`
	// PID 55 (byte A) - Short term secondary O2 trim - Bank 1
	ShortTermO2Trim1 float64 `json:"shortTermO2Trim1,omitempty"`
	// PID 57 (byte A) - Short term secondary O2 trim - Bank 2
	ShortTermO2Trim2 float64 `json:"shortTermO2Trim2,omitempty"`
	// PID 55 (byte B) - Short term secondary O2 trim - Bank 3
	ShortTermO2Trim3 float64 `json:"shortTermO2Trim3,omitempty"`
	// PID 57 (byte B) - Short term secondary O2 trim - Bank 4
	ShortTermO2Trim4 float64 `json:"shortTermO2Trim4,omitempty"`
	// PID 0D - Vehicle speed
	Speed float64 `json:"speed,omitempty"`
	// PID 01 - OBD status
	Status struct {
		// Number of sensor Trouble Codes (DTC)
		DTCCount uint8 `json:"dTCCount,omitempty"`
		// Type of the ignition for ICE - spark = spark plug ignition, compression = self-igniting (Diesel engines) Must be one of ["SPARK", "COMPRESSION"]
		IgnitionType string `json:"ignitionType,omitempty"`
		// Malfunction Indicator Light (MIL) False = Off, True = On
		IsMILOn bool `json:"isMILOn,omitempty"`
	} `json:"status,omitempty"`
	// PID 4C - Commanded throttle actuator
	ThrottleActuator float64 `json:"throttleActuator,omitempty"`
	// PID 11 - Throttle position - 0 = closed throttle, 100 = open throttle
	ThrottlePosition float64 `json:"throttlePosition,omitempty"`
	// PID 47 - Absolute throttle position B
	ThrottlePositionB float64 `json:"throttlePositionB,omitempty"`
	// PID 48 - Absolute throttle position C
	ThrottlePositionC float64 `json:"throttlePositionC,omitempty"`
	// PID 4E - Time since trouble codes cleared
	TimeSinceDTCCleared float64 `json:"timeSinceDTCCleared,omitempty"`
	// PID 0E - Time advance
	TimingAdvance float64 `json:"timingAdvance,omitempty"`
	// PID 30 - Number of warm-ups since codes cleared
	WarmupsSinceDTCClear uint8 `json:"warmupsSinceDTCClear,omitempty"`
}

// Driver data.
type Driver struct {
	// Probability of attentiveness of the driver.
	AttentiveProbability float64 `json:"attentiveProbability,omitempty"`
	// Distraction level of the driver will be the level how much the driver is distracted, by multiple factors. E.g. Driving situation, acustical or optical signales inside the cockpit, phone calls.
	DistractionLevel float64 `json:"distractionLevel,omitempty"`
	// Fatigueness level of driver. Evaluated by multiple factors like trip time, behaviour of steering, eye status.
	FatigueLevel float64 `json:"fatigueLevel,omitempty"`
	// Heart rate of the driver.
	HeartRate uint16 `json:"heartRate,omitempty"`
	// Identifier attributes based on OAuth 2.0.
	Identifier struct {
		// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
		Issuer string `json:"issuer,omitempty"`
		// Subject for the authentication of the occupant. E.g. UserID 7331677.
		Subject string `json:"subject,omitempty"`
	} `json:"identifier,omitempty"`
	// Has driver the eyes on road or not?
	IsEyesOnRoad bool `json:"isEyesOnRoad,omitempty"`
}

// Information about exterior measured by vehicle.
type Exterior struct {
	// Air temperature outside the vehicle.
	AirTemperature float64 `json:"airTemperature,omitempty"`
	// Relative humidity outside the vehicle. 0 = Dry, 100 = Air fully saturated.
	Humidity float64 `json:"humidity,omitempty"`
	// Light intensity outside the vehicle. 0 = No light detected, 100 = Fully lit.
	LightIntensity float64 `json:"lightIntensity,omitempty"`
}

// The current latitude and longitude of the vehicle.
type CurrentLocation struct {
	// Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	Altitude float64 `json:"altitude,omitempty"`
	// Information on the GNSS receiver used for determining current location.
	GNSSReceiver struct {
		// Fix status of GNSS receiver. Must be one of ["NONE", "TWO_D", "TWO_D_SATELLITE_BASED_AUGMENTATION", "TWO_D_GROUND_BASED_AUGMENTATION", "TWO_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", "THREE_D", "THREE_D_SATELLITE_BASED_AUGMENTATION", "THREE_D_GROUND_BASED_AUGMENTATION", "THREE_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION"]
		FixType string `json:"fixType,omitempty"`
		// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle.
		MountingPosition struct {
			// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = forward of rear axle. Negative values = backward of rear axle.
			X int16 `json:"x,omitempty"`
			// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = left of origin. Negative values = right of origin. Left/Right is as seen from driver perspective, i.e. by a person looking forward.
			Y int16 `json:"y,omitempty"`
			// Mounting position of GNSS receiver on Z-axis. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = above center of rear axle. Negative values = below center of rear axle.
			Z int16 `json:"z,omitempty"`
		} `json:"mountingPosition,omitempty"`
	} `json:"gNSSReceiver,omitempty"`
	// Current heading relative to geographic north. 0 = North, 90 = East, 180 = South, 270 = West.
	Heading float64 `json:"heading,omitempty"`
	// Accuracy of the latitude and longitude coordinates.
	HorizontalAccuracy float64 `json:"horizontalAccuracy,omitempty"`
	// Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	Latitude float64 `json:"latitude,omitempty"`
	// Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	Longitude float64 `json:"longitude,omitempty"`
	// Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	Timestamp string `json:"timestamp,omitempty"`
	// Accuracy of altitude.
	VerticalAccuracy float64 `json:"verticalAccuracy,omitempty"`
}

// Supported Version of VSS.
type VersionVSS struct {
	// Label to further describe the version.
	Label string `json:"label,omitempty"`
	// Supported Version of VSS - Major version.
	Major uint32 `json:"major,omitempty"`
	// Supported Version of VSS - Minor version.
	Minor uint32 `json:"minor,omitempty"`
	// Supported Version of VSS - Patch version.
	Patch uint32 `json:"patch,omitempty"`
}

type DataSchemaStruct struct {
	// High-level vehicle data.
	Vehicle struct {
		// All Advanced Driver Assist Systems data.
		ADAS DriveAssistSystems `json:"driveAssistSystems,omitempty"`
		// Spatial acceleration. Axis definitions according to ISO 8855.
		Acceleration Acceleration `json:"acceleration,omitempty"`
		// Spatial rotation. Axis definitions according to ISO 8855.
		AngularVelocity AngularVelocity `json:"angularVelocity,omitempty"`
		// Average speed for the current trip.
		AverageSpeed float64 `json:"averageSpeed,omitempty"`
		// All body components.
		Body Body `json:"body,omitempty"`
		// All in-cabin components, including doors.
		Cabin Cabin `json:"cabin,omitempty"`
		// All data concerning steering, suspension, wheels, and brakes.
		Chassis Chassis `json:"chassis,omitempty"`
		// Connectivity data.
		Connectivity Connectivity `json:"connectivity,omitempty"`
		// The current latitude and longitude of the vehicle.
		CurrentLocation CurrentLocation `json:"currentLocation,omitempty"`
		// Current overall Vehicle weight. Including passengers, cargo and other load inside the car.
		CurrentOverallWeight uint16 `json:"currentOverallWeight,omitempty"`
		// Driver data.
		Driver Driver `json:"driver,omitempty"`
		// The CO2 emissions.
		EmissionsCO2 int16 `json:"emissionsCO2,omitempty"`
		// Information about exterior measured by vehicle.
		Exterior Exterior `json:"exterior,omitempty"`
		// Curb weight of vehicle, including all liquids and full tank of fuel and full load of cargo and passengers.
		GrossWeight uint16 `json:"grossWeight,omitempty"`
		// Overall vehicle height.
		Height uint16 `json:"height,omitempty"`
		// Vehicle breakdown or any similar event causing vehicle to stop on the road, that might pose a risk to other road users. True = Vehicle broken down on the road, due to e.g. engine problems, flat tire, out of gas, brake problems. False = Vehicle not broken down.
		IsBrokenDown bool `json:"isBrokenDown,omitempty"`
		// Indicates whether the vehicle is stationary or moving.
		IsMoving bool `json:"isMoving,omitempty"`
		// Overall vehicle length.
		Length uint16 `json:"length,omitempty"`
		// Signals related to low voltage battery.
		LowVoltageBattery LowVoltageBattery `json:"lowVoltageBattery,omitempty"`
		// State of the supply voltage of the control units (usually 12V). Must be one of ["UNDEFINED", "LOCK", "OFF", "ACC", "ON", "START"]
		LowVoltageSystemState string `json:"lowVoltageSystemState,omitempty"`
		// OBD data.
		OBD OBD `json:"oBD,omitempty"`
		// Powertrain data for battery management, etc.
		Powertrain Powertrain `json:"powertrain,omitempty"`
		// Service data.
		Service Service `json:"service,omitempty"`
		// Vehicle speed.
		Speed float64 `json:"speed,omitempty"`
		// Trailer signals.
		Trailer Trailer `json:"trailer,omitempty"`
		// Odometer reading, total distance travelled during the lifetime of the vehicle.
		TravelledDistance float64 `json:"travelledDistance,omitempty"`
		// Current trip meter reading.
		TripMeterReading float64 `json:"tripMeterReading,omitempty"`
		// Supported Version of VSS.
		VersionVSS     VersionVSS     `json:"versionVSS,omitempty"`
		StaticFeatures StaticFeatures `json:"staticFeatures,omitempty"`
	} `json:"vehicle,omitempty"`
}

func NewVehicleStatus() DataSchemaStruct {
	return DataSchemaStruct{}
}

func (d DataSchemaStruct) Marshal() ([]byte, error) {
	if valid, err := d.IsValid(); !valid {
		return []byte{}, err
	}

	return json.Marshal(d)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
