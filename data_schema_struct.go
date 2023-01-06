package shared

import (
	"fmt"

	"github.com/clarketm/json"
	"github.com/volatiletech/null/v8"
)

func (v VehicleAttributes) IsValid() (bool, error) {
	valid := []string{"SAE_0", "SAE_1", "SAE_2", "SAE_3", "SAE_4", "SAE_5", ""}

	val := v.Vehicle.ADAS.SupportedAutonomyLevel
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.ADAS.SupportedAutonomyLevel: %v", val.String)
	}

	valid = []string{"FRONT_LEFT", "FRONT_RIGHT", "MIDDLE_LEFT", "MIDDLE_RIGHT", "REAR_LEFT", "REAR_RIGHT", ""}

	val = v.Vehicle.Body.RefuelPosition
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.RefuelPosition: %v", val.String)
	}

	valid = []string{"FRONT_LEFT", "FRONT_RIGHT", ""}

	val = v.Vehicle.Chassis.SteeringWheel.Position
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Chassis.SteeringWheel.Position: %v", val.String)
	}

	valid = []string{"UNKNOWN", "NATURAL", "SUPERCHARGER", "TURBOCHARGER", ""}

	val = v.Vehicle.Powertrain.CombustionEngine.AspirationType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.AspirationType: %v", val.String)
	}

	valid = []string{"UNKNOWN", "STRAIGHT", "V", "BOXER", "W", "ROTARY", "RADIAL", "SQUARE", "H", "U", "OPPOSED", "X", ""}

	val = v.Vehicle.Powertrain.CombustionEngine.Configuration
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.Configuration: %v", val.String)
	}

	valid = []string{"UNKNOWN", "NOT_APPLICABLE", "STOP_START", "BELT_ISG", "CIMG", "PHEV", ""}

	val = v.Vehicle.Powertrain.FuelSystem.HybridType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.FuelSystem.HybridType: %v", val.String)
	}

	valid = []string{"UNKNOWN", "FORWARD_WHEEL_DRIVE", "REAR_WHEEL_DRIVE", "ALL_WHEEL_DRIVE", ""}

	val = v.Vehicle.Powertrain.Transmission.DriveType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.DriveType: %v", val.String)
	}

	valid = []string{"UNKNOWN", "SEQUENTIAL", "H", "AUTOMATIC", "DSG", "CVT", ""}

	val = v.Vehicle.Powertrain.Transmission.Type
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.Type: %v", val.String)
	}

	valid = []string{"COMBUSTION", "HYBRID", "ELECTRIC", ""}

	val = v.Vehicle.Powertrain.Type
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Type: %v", val.String)
	}

	return true, nil
}

func (d DataSchemaStruct) IsValid() (bool, error) {
	valid := []string{"SAE_0", "SAE_1", "SAE_2_DISENGAGING", "SAE_2", "SAE_3_DISENGAGING", "SAE_3", "SAE_4_DISENGAGING", "SAE_4", "SAE_5", ""}

	val := d.Vehicle.ADAS.ActiveAutonomyLevel
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.ADAS.ActiveAutonomyLevel: %v", val.String)
	}

	valid = []string{"INACTIVE", "ACTIVE", "ADAPTIVE", ""}

	val = d.Vehicle.Body.Lights.Brake.IsActive
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Lights.Brake.IsActive: %v", val.String)
	}

	valid = []string{"OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR", ""}

	val = d.Vehicle.Body.Windshield.Front.Wiping.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Front.Wiping.Mode: %v", val.String)
	}

	valid = []string{"STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP", ""}

	val = d.Vehicle.Body.Windshield.Front.Wiping.System.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Front.Wiping.System.Mode: %v", val.String)
	}

	valid = []string{"OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR", ""}

	val = d.Vehicle.Body.Windshield.Rear.Wiping.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Rear.Wiping.Mode: %v", val.String)
	}

	valid = []string{"STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP", ""}

	val = d.Vehicle.Body.Windshield.Rear.Wiping.System.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Body.Windshield.Rear.Wiping.System.Mode: %v", val.String)
	}

	valid = []string{"UNDEFINED", "CLOSED", "OPEN", "CLOSING", "OPENING", "STALLED", ""}

	val = d.Vehicle.Cabin.Convertible.Status
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Convertible.Status: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Left.Shade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Left.Shade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Left.Window.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Left.Window.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Right.Shade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Right.Shade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row1.Right.Window.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row1.Right.Window.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Left.Shade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Left.Shade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Left.Window.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Left.Window.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Right.Shade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Right.Shade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Door.Row2.Right.Window.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Door.Row2.Right.Window.Switch: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row1.Left.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row1.Left.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row1.Right.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row1.Right.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row2.Left.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row2.Left.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row2.Right.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row2.Right.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row3.Left.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row3.Left.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row3.Right.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row3.Right.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row4.Left.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row4.Left.AirDistribution: %v", val.String)
	}

	valid = []string{"UP", "MIDDLE", "DOWN", ""}

	val = d.Vehicle.Cabin.HVAC.Station.Row4.Right.AirDistribution
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.HVAC.Station.Row4.Right.AirDistribution: %v", val.String)
	}

	valid = []string{"YYYY_MM_DD", "DD_MM_YYYY", "MM_DD_YYYY", "YY_MM_DD", "DD_MM_YY", "MM_DD_YY", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DateFormat
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DateFormat: %v", val.String)
	}

	valid = []string{"DAY", "NIGHT", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DayNightMode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DayNightMode: %v", val.String)
	}

	valid = []string{"MILES", "KILOMETERS", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.DistanceUnit
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.DistanceUnit: %v", val.String)
	}

	valid = []string{"MILES_PER_KILOWATT_HOUR", "KILOMETERS_PER_KILOWATT_HOUR", "KILOWATT_HOURS_PER_100_MILES", "KILOWATT_HOURS_PER_100_KILOMETERS", "WATT_HOURS_PER_MILE", "WATT_HOURS_PER_KILOMETER", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.EVEconomyUnits
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.EVEconomyUnits: %v", val.String)
	}

	valid = []string{"MPG_UK", "MPG_US", "MILES_PER_LITER", "KILOMETERS_PER_LITER", "LITERS_PER_100_KILOMETERS", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.FuelEconomyUnits
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.FuelEconomyUnits: %v", val.String)
	}

	valid = []string{"C", "F", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TemperatureUnit
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TemperatureUnit: %v", val.String)
	}

	valid = []string{"HR_12", "HR_24", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TimeFormat
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TimeFormat: %v", val.String)
	}

	valid = []string{"‘PSI’", "‘KPA’", "’BAR’", ""}

	val = d.Vehicle.Cabin.Infotainment.HMI.TirePressureUnit
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.HMI.TirePressureUnit: %v", val.String)
	}

	valid = []string{"UNKNOWN", "STOP", "PLAY", "FAST_FORWARD", "FAST_BACKWARD", "SKIP_FORWARD", "SKIP_BACKWARD", ""}

	val = d.Vehicle.Cabin.Infotainment.Media.Action
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Media.Action: %v", val.String)
	}

	valid = []string{"UNKNOWN", "SIRIUS_XM", "AM", "FM", "DAB", "TV", "CD", "DVD", "AUX", "USB", "DISK", "BLUETOOTH", "INTERNET", "VOICE", "BEEP", ""}

	val = d.Vehicle.Cabin.Infotainment.Media.Played.Source
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Media.Played.Source: %v", val.String)
	}

	valid = []string{"MUTED", "ALERT_ONLY", "UNMUTED", ""}

	val = d.Vehicle.Cabin.Infotainment.Navigation.Mute
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.Navigation.Mute: %v", val.String)
	}

	valid = []string{"NONE", "ACTIVE", "INACTIVE", ""}

	val = d.Vehicle.Cabin.Infotainment.SmartphoneProjection.Active
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.SmartphoneProjection.Active: %v", val.String)
	}

	valid = []string{"USB", "BLUETOOTH", "WIFI", ""}

	val = d.Vehicle.Cabin.Infotainment.SmartphoneProjection.Source
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Infotainment.SmartphoneProjection.Source: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.RearShade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.RearShade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", ""}

	val = d.Vehicle.Cabin.Sunroof.Shade.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Sunroof.Shade.Switch: %v", val.String)
	}

	valid = []string{"INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", "TILT_UP", "TILT_DOWN", ""}

	val = d.Vehicle.Cabin.Sunroof.Switch
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Cabin.Sunroof.Switch: %v", val.String)
	}

	valid = []string{"NONE", "TWO_D", "TWO_D_SATELLITE_BASED_AUGMENTATION", "TWO_D_GROUND_BASED_AUGMENTATION", "TWO_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", "THREE_D", "THREE_D_SATELLITE_BASED_AUGMENTATION", "THREE_D_GROUND_BASED_AUGMENTATION", "THREE_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", ""}

	val = d.Vehicle.CurrentLocation.GNSSReceiver.FixType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.CurrentLocation.GNSSReceiver.FixType: %v", val.String)
	}

	valid = []string{"UNDEFINED", "LOCK", "OFF", "ACC", "ON", "START", ""}

	val = d.Vehicle.LowVoltageSystemState
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.LowVoltageSystemState: %v", val.String)
	}

	valid = []string{"SPARK", "COMPRESSION", ""}

	val = d.Vehicle.OBD.DriveCycleStatus.IgnitionType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.OBD.DriveCycleStatus.IgnitionType: %v", val.String)
	}

	valid = []string{"SPARK", "COMPRESSION", ""}

	val = d.Vehicle.OBD.Status.IgnitionType
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.OBD.Status.IgnitionType: %v", val.String)
	}

	valid = []string{"CRITICALLY_LOW", "LOW", "NORMAL", "HIGH", "CRITICALLY_HIGH", ""}

	val = d.Vehicle.Powertrain.CombustionEngine.EngineOilLevel
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.CombustionEngine.EngineOilLevel: %v", val.String)
	}

	valid = []string{"OPEN", "CLOSED", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.ChargePortFlap
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.ChargePortFlap: %v", val.String)
	}

	valid = []string{"MANUAL", "TIMER", "GRID", "PROFILE", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.Mode: %v", val.String)
	}

	valid = []string{"START", "STOP", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.StartStopCharging
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.StartStopCharging: %v", val.String)
	}

	valid = []string{"INACTIVE", "START_TIME", "END_TIME", ""}

	val = d.Vehicle.Powertrain.TractionBattery.Charging.Timer.Mode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.TractionBattery.Charging.Timer.Mode: %v", val.String)
	}

	valid = []string{"MANUAL", "AUTOMATIC", ""}

	val = d.Vehicle.Powertrain.Transmission.GearChangeMode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.GearChangeMode: %v", val.String)
	}

	valid = []string{"NORMAL", "SPORT", "ECONOMY", "SNOW", "RAIN", ""}

	val = d.Vehicle.Powertrain.Transmission.PerformanceMode
	if !contains(valid, val.String) {
		return false, fmt.Errorf("invalid value at Vehicle.Powertrain.Transmission.PerformanceMode: %v", val.String)
	}

	return true, nil
}

// All Advanced Driver Assist Systems attributes.
type AdvancedDriveAssistAttributes struct {
	// Indicates the highest level of autonomy according to SAE J3016 taxonomy the vehicle is capable of.
	// Must be one of ["SAE_0", "SAE_1", "SAE_2", "SAE_3", "SAE_4", "SAE_5"]
	SupportedAutonomyLevel null.String `json:"supportedAutonomyLevel,omitempty"`
}

// All body attributes.
type BodyAttributes struct {
	// Body type code as defined by ISO 3779.
	BodyType null.String `json:"bodyType,omitempty"`
	// Location of the fuel cap or charge port.
	// Must be one of ["FRONT_LEFT", "FRONT_RIGHT", "MIDDLE_LEFT", "MIDDLE_RIGHT", "REAR_LEFT", "REAR_RIGHT"]
	RefuelPosition null.String `json:"refuelPosition,omitempty"`
}

// All in-cabin attributes
type CabinAttributes struct {
	// Number of doors in vehicle.
	DoorCount null.Uint8 `json:"doorCount,omitempty"`
	// The position of the driver seat in row 1.
	DriverPosition null.Uint8 `json:"driverPosition,omitempty"`
	// Infotainment system.
	Infotainment struct {
		// All smartphone projection actions.
		SmartphoneProjection struct {
			// Supportable list for projection.
			// Must be one of ["ANDROID_AUTO", "APPLE_CARPLAY", "MIRROR_LINK", "OTHER"]
			SupportedMode []null.String `json:"supportedMode,omitempty"`
		} `json:"smartphoneProjection,omitempty"`
	} `json:"infotainment,omitempty"`
	// Number of seats across each row from the front to the rear.
	SeatPosCount []null.Uint8 `json:"seatPosCount,omitempty"`
	// Number of seat rows in vehicle.
	SeatRowCount null.Uint8 `json:"seatRowCount,omitempty"`
}

// All attributes concerning steering, suspension, wheels, and brakes.
type ChassisAttributes struct {
	// Axle signals
	Axle struct {
		// Axle signals
		Row1 struct {
			// Aspect ratio between tire section height and tire section width, as per ETRTO / TRA standard.
			TireAspectRatio null.Uint8 `json:"tireAspectRatio,omitempty"`
			// Outer diameter of tires, in inches, as per ETRTO / TRA standard.
			TireDiameter null.Float64 `json:"tireDiameter,omitempty"`
			// Nominal section width of tires, in mm, as per ETRTO / TRA standard.
			TireWidth null.Uint16 `json:"tireWidth,omitempty"`
			// Number of wheels on the axle
			WheelCount null.Uint8 `json:"wheelCount,omitempty"`
			// Diameter of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelDiameter null.Float64 `json:"wheelDiameter,omitempty"`
			// Width of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelWidth null.Float64 `json:"wheelWidth,omitempty"`
		} `json:"row1,omitempty"`
		// Axle signals
		Row2 struct {
			// Aspect ratio between tire section height and tire section width, as per ETRTO / TRA standard.
			TireAspectRatio null.Uint8 `json:"tireAspectRatio,omitempty"`
			// Outer diameter of tires, in inches, as per ETRTO / TRA standard.
			TireDiameter null.Float64 `json:"tireDiameter,omitempty"`
			// Nominal section width of tires, in mm, as per ETRTO / TRA standard.
			TireWidth null.Uint16 `json:"tireWidth,omitempty"`
			// Number of wheels on the axle
			WheelCount null.Uint8 `json:"wheelCount,omitempty"`
			// Diameter of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelDiameter null.Float64 `json:"wheelDiameter,omitempty"`
			// Width of wheels (rims without tires), in inches, as per ETRTO / TRA standard.
			WheelWidth null.Float64 `json:"wheelWidth,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"axle,omitempty"`
	// Number of axles on the vehicle
	AxleCount null.Uint8 `json:"axleCount,omitempty"`
	// Steering wheel signals
	SteeringWheel struct {
		// Position of the steering wheel on the left or right side of the vehicle.
		// Must be one of ["FRONT_LEFT", "FRONT_RIGHT"]
		Position null.String `json:"position,omitempty"`
	} `json:"steeringWheel,omitempty"`
	// Overall wheel tracking, in mm.
	Track null.Uint16 `json:"track,omitempty"`
	// Overall wheel base, in mm.
	Wheelbase null.Uint16 `json:"wheelbase,omitempty"`
}

// The current latitude and longitude attributes of the vehicle.
type LocationAttributes struct {
	// Information on the GNSS receiver used for determining current location.
	GNSSReceiver struct {
		// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle.
		MountingPosition struct {
			// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = forward of rear axle. Negative values = backward of rear axle.
			X null.Int16 `json:"x,omitempty"`
			// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = left of origin. Negative values = right of origin. Left/Right is as seen from driver perspective, i.e. by a person looking forward.
			Y null.Int16 `json:"y,omitempty"`
			// Mounting position of GNSS receiver on Z-axis. Axis definitions according to ISO 8855. Origin at center of (first) rear axle. Positive values = above center of rear axle. Negative values = below center of rear axle.
			Z null.Int16 `json:"z,omitempty"`
		} `json:"mountingPosition,omitempty"`
	} `json:"gNSSReceiver,omitempty"`
}

// Signals related to low voltage battery.
type LowVoltageBatteryAttributes struct {
	// Nominal capacity of the low voltage battery.
	NominalCapacity null.Uint16 `json:"nominalCapacity,omitempty"`
	// Nominal Voltage of the battery.
	NominalVoltage null.Uint16 `json:"nominalVoltage,omitempty"`
}

// OBD data.
type OBDAttributes struct {
	// PID 1C - OBD standards this vehicle conforms to
	OBDStandards null.Uint8 `json:"oBDStandards,omitempty"`
}

// Powertrain data for battery management, etc.
type PowertrainAttributes struct {
	// Engine-specific data, stopping at the bell housing.
	CombustionEngine struct {
		// Type of aspiration (natural, turbocharger, supercharger etc).
		// Must be one of ["UNKNOWN", "NATURAL", "SUPERCHARGER", "TURBOCHARGER"]
		AspirationType null.String `json:"aspirationType,omitempty"`
		// Bore in millimetres.
		Bore null.Float64 `json:"bore,omitempty"`
		// Engine compression ratio, specified in the format 'X:1', e.g. '9.2:1'.
		CompressionRatio null.String `json:"compressionRatio,omitempty"`
		// Engine configuration.
		// Must be one of ["UNKNOWN", "STRAIGHT", "V", "BOXER", "W", "ROTARY", "RADIAL", "SQUARE", "H", "U", "OPPOSED", "X"]
		Configuration null.String `json:"configuration,omitempty"`
		// Signals related to Diesel Exhaust Fluid (DEF). DEF is called AUS32 in ISO 22241.
		DieselExhaustFluid struct {
			// Capacity in liters of the Diesel Exhaust Fluid Tank.
			Capacity null.Float64 `json:"capacity,omitempty"`
		} `json:"dieselExhaustFluid,omitempty"`
		// Displacement in cubic centimetres.
		Displacement null.Uint16 `json:"displacement,omitempty"`
		// Engine code designation, as specified by vehicle manufacturer.
		EngineCode null.String `json:"engineCode,omitempty"`
		// Engine coolant capacity in liters.
		EngineCoolantCapacity null.Float64 `json:"engineCoolantCapacity,omitempty"`
		// Engine oil capacity in liters.
		EngineOilCapacity null.Float64 `json:"engineOilCapacity,omitempty"`
		// Peak power, in kilowatts, that engine can generate.
		MaxPower null.Uint16 `json:"maxPower,omitempty"`
		// Peak torque, in newton meter, that the engine can generate.
		MaxTorque null.Uint16 `json:"maxTorque,omitempty"`
		// Number of cylinders.
		NumberOfCylinders null.Uint16 `json:"numberOfCylinders,omitempty"`
		// Number of valves per cylinder.
		NumberOfValvesPerCylinder null.Uint16 `json:"numberOfValvesPerCylinder,omitempty"`
		// Stroke length in millimetres.
		StrokeLength null.Float64 `json:"strokeLength,omitempty"`
	} `json:"combustionEngine,omitempty"`
	// Electric Motor specific data.
	ElectricMotor struct {
		// Engine code designation, as specified by vehicle manufacturer.
		EngineCode null.String `json:"engineCode,omitempty"`
		// Peak power, in kilowatts, that motor(s) can generate.
		MaxPower null.Uint16 `json:"maxPower,omitempty"`
		// Peak regen/brake power, in kilowatts, that motor(s) can generate.
		MaxRegenPower null.Uint16 `json:"maxRegenPower,omitempty"`
		// Peak regen/brake torque, in newton meter, that the motor(s) can generate.
		MaxRegenTorque null.Uint16 `json:"maxRegenTorque,omitempty"`
		// Peak power, in newton meter, that the motor(s) can generate.
		MaxTorque null.Uint16 `json:"maxTorque,omitempty"`
	} `json:"electricMotor,omitempty"`
	// Fuel system data.
	FuelSystem struct {
		// Defines the hybrid type of the vehicle.
		// Must be one of ["UNKNOWN", "NOT_APPLICABLE", "STOP_START", "BELT_ISG", "CIMG", "PHEV"]
		HybridType null.String `json:"hybridType,omitempty"`
		// Detailed information on fuels supported by the vehicle. Identifiers originating from DIN EN 16942:2021-08, appendix B, with additional suffix for octane (RON) where relevant.
		// Must be one of ["E5_95", "E5_98", "E10_95", "E10_98", "E85", "B7", "B10", "B20", "B30", "B100", "XTL", "LPG", "CNG", "LNG", "H2", "OTHER"]
		SupportedFuel []null.String `json:"supportedFuel,omitempty"`
		// High level information of fuel types supported
		// Must be one of ["GASOLINE", "DIESEL", "E85", "LPG", "CNG", "LNG", "H2", "OTHER"]
		SupportedFuelTypes []null.String `json:"supportedFuelTypes,omitempty"`
		// Capacity of the fuel tank in liters.
		TankCapacity null.Float64 `json:"tankCapacity,omitempty"`
	} `json:"fuelSystem,omitempty"`
	// Battery Management data.
	TractionBattery struct {
		// Properties related to battery charging.
		Charging struct {
			// Type of charge plug (charging inlet) available on the vehicle. IEC types refer to IEC 62196,  GBT refers to  GB/T 20234.
			// Must be one of ["IEC_TYPE_1_AC", "IEC_TYPE_2_AC", "IEC_TYPE_3_AC", "IEC_TYPE_4_DC", "IEC_TYPE_1_CCS_DC", "IEC_TYPE_2_CCS_DC", "TESLA_ROADSTER", "TESLA_HPWC", "TESLA_SUPERCHARGER", "GBT_AC", "GBT_DC", "OTHER"]
			ChargePlugType []null.String `json:"chargePlugType,omitempty"`
		} `json:"charging,omitempty"`
		// Gross capacity of the battery.
		GrossCapacity null.Uint16 `json:"grossCapacity,omitempty"`
		// Battery Identification Number as assigned by OEM.
		Id null.String `json:"id,omitempty"`
		// Max allowed voltage of the battery, e.g. during charging.
		MaxVoltage null.Uint16 `json:"maxVoltage,omitempty"`
		// Nominal Voltage of the battery.
		NominalVoltage null.Uint16 `json:"nominalVoltage,omitempty"`
		// Production date of battery in ISO8601 format, e.g. YYYY-MM-DD.
		ProductionDate null.String `json:"productionDate,omitempty"`
	} `json:"tractionBattery,omitempty"`
	// Transmission-specific data, stopping at the drive shafts.
	Transmission struct {
		// Drive type.
		// Must be one of ["UNKNOWN", "FORWARD_WHEEL_DRIVE", "REAR_WHEEL_DRIVE", "ALL_WHEEL_DRIVE"]
		DriveType null.String `json:"driveType,omitempty"`
		// Number of forward gears in the transmission. -1 = CVT.
		GearCount null.Int8 `json:"gearCount,omitempty"`
		// Transmission type.
		// Must be one of ["UNKNOWN", "SEQUENTIAL", "H", "AUTOMATIC", "DSG", "CVT"]
		Type null.String `json:"type,omitempty"`
	} `json:"transmission,omitempty"`
	// Defines the powertrain type of the vehicle.
	// Must be one of ["COMBUSTION", "HYBRID", "ELECTRIC"]
	Type null.String `json:"type,omitempty"`
}

// Attributes that identify a vehicle.
type VehicleIdentificationAttributes struct {
	// The ACRISS Car Classification Code is a code used by many car rental companies.
	AcrissCode null.String `json:"acrissCode,omitempty"`
	// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	BodyType null.String `json:"bodyType,omitempty"`
	// Vehicle brand or manufacturer.
	Brand null.String `json:"brand,omitempty"`
	// The date in ISO 8601 format of the first registration of the vehicle with the respective public authorities.
	DateVehicleFirstRegistered null.String `json:"dateVehicleFirstRegistered,omitempty"`
	// A textual description of known damages, both repaired and unrepaired.
	KnownVehicleDamages null.String `json:"knownVehicleDamages,omitempty"`
	// Indicates that the vehicle meets the respective emission standard.
	MeetsEmissionStandard null.String `json:"meetsEmissionStandard,omitempty"`
	// Vehicle model.
	Model null.String `json:"model,omitempty"`
	// The date in ISO 8601 format of production of the item, e.g. vehicle.
	ProductionDate null.String `json:"productionDate,omitempty"`
	// The date in ISO 8601 format of the item e.g. vehicle was purchased by the current owner.
	PurchaseDate null.String `json:"purchaseDate,omitempty"`
	// 17-character Vehicle Identification Number (VIN) as defined by ISO 3779.
	VIN null.String `json:"vIN,omitempty"`
	// A short text indicating the configuration of the vehicle, e.g. '5dr hatchback ST 2.5 MT 225 hp' or 'limited edition'.
	VehicleConfiguration null.String `json:"vehicleConfiguration,omitempty"`
	// The color or color combination of the interior of the vehicle.
	VehicleInteriorColor null.String `json:"vehicleInteriorColor,omitempty"`
	// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.).
	VehicleInteriorType null.String `json:"vehicleInteriorType,omitempty"`
	// The release date in ISO 8601 format of a vehicle model (often used to differentiate versions of the same make and model).
	VehicleModelDate null.String `json:"vehicleModelDate,omitempty"`
	// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
	VehicleSeatingCapacity null.Uint16 `json:"vehicleSeatingCapacity,omitempty"`
	// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school.
	VehicleSpecialUsage null.String `json:"vehicleSpecialUsage,omitempty"`
	// 3-character World Manufacturer Identification (WMI) as defined by ISO 3780.
	WMI null.String `json:"wMI,omitempty"`
	// Model year of the vehicle.
	Year null.Uint16 `json:"year,omitempty"`
}

// Supported Version of VSS.
type VersionVSSAttributes struct {
	// Label to further describe the version.
	Label null.String `json:"label,omitempty"`
	// Supported Version of VSS - Major version.
	Major null.Uint32 `json:"major,omitempty"`
	// Supported Version of VSS - Minor version.
	Minor null.Uint32 `json:"minor,omitempty"`
	// Supported Version of VSS - Patch version.
	Patch null.Uint32 `json:"patch,omitempty"`
}

// VehicleAttributes represent static features of the vehicle
type VehicleAttributes struct {
	// High-level vehicle data.
	Vehicle struct {
		// All Advanced Driver Assist Systems data.
		ADAS AdvancedDriveAssistAttributes `json:"advancedDriveAssist,omitempty"`
		// All body components.
		Body BodyAttributes `json:"body,omitempty"`
		// All in-cabin components, including doors.
		Cabin CabinAttributes `json:"cabin,omitempty"`
		// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
		CargoVolume null.Float64 `json:"cargoVolume,omitempty"`
		// All data concerning steering, suspension, wheels, and brakes.
		Chassis ChassisAttributes `json:"chassis,omitempty"`
		// Vehicle curb weight, including all liquids and full tank of fuel, but no cargo or passengers.
		CurbWeight null.Uint16 `json:"curbWeight,omitempty"`
		// The current latitude and longitude of the vehicle.
		CurrentLocation LocationAttributes `json:"currentLocation,omitempty"`
		// The CO2 emissions.
		EmissionsCO2 null.Int16 `json:"emissionsCO2,omitempty"`
		// Curb weight of vehicle, including all liquids and full tank of fuel and full load of cargo and passengers.
		GrossWeight null.Uint16 `json:"grossWeight,omitempty"`
		// Overall vehicle height.
		Height null.Uint16 `json:"height,omitempty"`
		// Overall vehicle length.
		Length null.Uint16 `json:"length,omitempty"`
		// Signals related to low voltage battery.
		LowVoltageBattery LowVoltageBatteryAttributes `json:"lowVoltageBattery,omitempty"`
		// Maximum vertical weight on the tow ball of a trailer.
		MaxTowBallWeight null.Uint16 `json:"maxTowBallWeight,omitempty"`
		// Maximum weight of trailer.
		MaxTowWeight null.Uint16 `json:"maxTowWeight,omitempty"`
		// OBD data.
		OBD OBDAttributes `json:"oBD,omitempty"`
		// Powertrain data for battery management, etc.
		Powertrain PowertrainAttributes `json:"powertrain,omitempty"`
		// The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.
		RoofLoad null.Int16 `json:"roofLoad,omitempty"`
		// Attributes that identify a vehicle.
		VehicleIdentification VehicleIdentificationAttributes `json:"vehicleIdentification,omitempty"`
		// Supported Version of VSS.
		VersionVSS VersionVSSAttributes `json:"versionVSS,omitempty"`
		// Overall vehicle width.
		Width null.Uint16 `json:"width,omitempty"`
	} `json:"vehicle,omitempty"`
}

// All Advanced Driver Assist Systems data.
type AdvancedDriveAssistVariables struct {
	// Antilock Braking System signals.
	ABS struct {
		// Indicates if ABS is enabled. True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if ABS is currently regulating brake pressure. True = Engaged. False = Not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
		// Indicates if ABS incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
	} `json:"aBS,omitempty"`
	// Indicates the currently active level of autonomy according to SAE J3016 taxonomy.
	// Must be one of ["SAE_0", "SAE_1", "SAE_2_DISENGAGING", "SAE_2", "SAE_3_DISENGAGING", "SAE_3", "SAE_4_DISENGAGING", "SAE_4", "SAE_5"]
	ActiveAutonomyLevel null.String `json:"activeAutonomyLevel,omitempty"`
	// Signals from Cruise Control system.
	CruiseControl struct {
		// Indicates if cruise control system is active (i.e. actively controls speed). True = Active. False = Inactive.
		IsActive null.Bool `json:"isActive,omitempty"`
		// Indicates if cruise control system is enabled (e.g. ready to receive configurations and settings) True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if cruise control system incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
		// Set cruise control speed in kilometers per hour.
		SpeedSet null.Float64 `json:"speedSet,omitempty"`
	} `json:"cruiseControl,omitempty"`
	// Emergency Brake Assist (EBA) System signals.
	EBA struct {
		// Indicates if EBA is enabled. True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if EBA is currently regulating brake pressure. True = Engaged. False = Not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
		// Indicates if EBA incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
	} `json:"eBA,omitempty"`
	// Electronic Brakeforce Distribution (EBD) System signals.
	EBD struct {
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if EBD is currently regulating vehicle brakeforce distribution. True = Engaged. False = Not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
		// Indicates if EBD incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
	} `json:"eBD,omitempty"`
	// Electronic Stability Control System signals.
	ESC struct {
		// Indicates if ESC is enabled. True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if ESC is currently regulating vehicle stability. True = Engaged. False = Not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
		// Indicates if ESC incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
		// Indicates if the ESC system is detecting strong cross winds. True = Strong cross winds detected. False = No strong cross winds detected.
		IsStrongCrossWindDetected null.Bool `json:"isStrongCrossWindDetected,omitempty"`
		// Road friction values reported by the ESC system.
		RoadFriction struct {
			// Lower bound road friction, as calculated by the ESC system. 5% possibility that road friction is below this value. 0 = no friction, 100 = maximum friction.
			LowerBound null.Float64 `json:"lowerBound,omitempty"`
			// Most probable road friction, as calculated by the ESC system. Exact meaning of most probable is implementation specific. 0 = no friction, 100 = maximum friction.
			MostProbable null.Float64 `json:"mostProbable,omitempty"`
			// Upper bound road friction, as calculated by the ESC system. 95% possibility that road friction is below this value. 0 = no friction, 100 = maximum friction.
			UpperBound null.Float64 `json:"upperBound,omitempty"`
		} `json:"roadFriction,omitempty"`
	} `json:"eSC,omitempty"`
	// Signals from Lane Departure Detection System.
	LaneDepartureDetection struct {
		// Indicates if lane departure detection system is enabled. True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if lane departure system incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
		// Indicates if lane departure detection registered a lane departure.
		IsWarning null.Bool `json:"isWarning,omitempty"`
	} `json:"laneDepartureDetection,omitempty"`
	// Signals form Obstacle Sensor System.
	ObstacleDetection struct {
		// Indicates if obstacle sensor system is enabled (i.e. monitoring for obstacles). True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if obstacle sensor system incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
		// Indicates if obstacle sensor system registered an obstacle.
		IsWarning null.Bool `json:"isWarning,omitempty"`
	} `json:"obstacleDetection,omitempty"`
	// Traction Control System signals.
	TCS struct {
		// Indicates if TCS is enabled. True = Enabled. False = Disabled.
		IsEnabled null.Bool `json:"isEnabled,omitempty"`
		// Indicates if TCS is currently regulating traction. True = Engaged. False = Not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
		// Indicates if TCS incurred an error condition. True = Error. False = No Error.
		IsError null.Bool `json:"isError,omitempty"`
	} `json:"tCS,omitempty"`
}

// Spatial acceleration. Axis definitions according to ISO 8855.
type AccelerationVariables struct {
	// Vehicle acceleration in Y (lateral acceleration).
	Lateral null.Float64 `json:"lateral,omitempty"`
	// Vehicle acceleration in X (longitudinal acceleration).
	Longitudinal null.Float64 `json:"longitudinal,omitempty"`
	// Vehicle acceleration in Z (vertical acceleration).
	Vertical null.Float64 `json:"vertical,omitempty"`
}

// Spatial rotation. Axis definitions according to ISO 8855.
type AngularVelocityVariables struct {
	// Vehicle rotation rate along Y (lateral).
	Pitch null.Float64 `json:"pitch,omitempty"`
	// Vehicle rotation rate along X (longitudinal).
	Roll null.Float64 `json:"roll,omitempty"`
	// Vehicle rotation rate along Z (vertical).
	Yaw null.Float64 `json:"yaw,omitempty"`
}

// All body components.
type BodyVariables struct {
	// Hood status.
	Hood struct {
		// Hood open or closed. True = Open. False = Closed.
		IsOpen null.Bool `json:"isOpen,omitempty"`
	} `json:"hood,omitempty"`
	// Horn signals.
	Horn struct {
		// Horn active or inactive. True = Active. False = Inactive.
		IsActive null.Bool `json:"isActive,omitempty"`
	} `json:"horn,omitempty"`
	// Exterior lights.
	Lights struct {
		// Backup lights.
		Backup struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn null.Bool `json:"isOn,omitempty"`
		} `json:"backup,omitempty"`
		// Beam lights.
		Beam struct {
			// Beam lights.
			High struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn null.Bool `json:"isOn,omitempty"`
			} `json:"high,omitempty"`
			// Beam lights.
			Low struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn null.Bool `json:"isOn,omitempty"`
			} `json:"low,omitempty"`
		} `json:"beam,omitempty"`
		//
		Brake struct {
			// Indicates if break-light is active. INACTIVE means lights are off. ACTIVE means lights are on. ADAPTIVE means that break-light is indicating emergency-breaking.
			// Must be one of ["INACTIVE", "ACTIVE", "ADAPTIVE"]
			IsActive null.String `json:"isActive,omitempty"`
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
		} `json:"brake,omitempty"`
		// Indicator lights.
		DirectionIndicator struct {
			// Indicator lights.
			Left struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is signaling or off. True = signaling. False = Off.
				IsSignaling null.Bool `json:"isSignaling,omitempty"`
			} `json:"left,omitempty"`
			// Indicator lights.
			Right struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is signaling or off. True = signaling. False = Off.
				IsSignaling null.Bool `json:"isSignaling,omitempty"`
			} `json:"right,omitempty"`
		} `json:"directionIndicator,omitempty"`
		// Fog lights.
		Fog struct {
			// Fog lights.
			Front struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn null.Bool `json:"isOn,omitempty"`
			} `json:"front,omitempty"`
			// Fog lights.
			Rear struct {
				// Indicates if light is defect. True = Light is defect. False = Light has no defect.
				IsDefect null.Bool `json:"isDefect,omitempty"`
				// Indicates if light is on or off. True = On. False = Off.
				IsOn null.Bool `json:"isOn,omitempty"`
			} `json:"rear,omitempty"`
		} `json:"fog,omitempty"`
		// Hazard lights.
		Hazard struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
			// Indicates if light is signaling or off. True = signaling. False = Off.
			IsSignaling null.Bool `json:"isSignaling,omitempty"`
		} `json:"hazard,omitempty"`
		// License plate lights.
		LicensePlate struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn null.Bool `json:"isOn,omitempty"`
		} `json:"licensePlate,omitempty"`
		// Parking lights.
		Parking struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn null.Bool `json:"isOn,omitempty"`
		} `json:"parking,omitempty"`
		// Running lights.
		Running struct {
			// Indicates if light is defect. True = Light is defect. False = Light has no defect.
			IsDefect null.Bool `json:"isDefect,omitempty"`
			// Indicates if light is on or off. True = On. False = Off.
			IsOn null.Bool `json:"isOn,omitempty"`
		} `json:"running,omitempty"`
	} `json:"lights,omitempty"`
	// All mirrors.
	Mirrors struct {
		// All mirrors.
		Left struct {
			// Mirror Heater on or off. True = Heater On. False = Heater Off.
			IsHeatingOn null.Bool `json:"isHeatingOn,omitempty"`
			// Mirror pan as a percent. 0 = Center Position. 100 = Fully Left Position. -100 = Fully Right Position.
			Pan null.Int8 `json:"pan,omitempty"`
			// Mirror tilt as a percent. 0 = Center Position. 100 = Fully Upward Position. -100 = Fully Downward Position.
			Tilt null.Int8 `json:"tilt,omitempty"`
		} `json:"left,omitempty"`
		// All mirrors.
		Right struct {
			// Mirror Heater on or off. True = Heater On. False = Heater Off.
			IsHeatingOn null.Bool `json:"isHeatingOn,omitempty"`
			// Mirror pan as a percent. 0 = Center Position. 100 = Fully Left Position. -100 = Fully Right Position.
			Pan null.Int8 `json:"pan,omitempty"`
			// Mirror tilt as a percent. 0 = Center Position. 100 = Fully Upward Position. -100 = Fully Downward Position.
			Tilt null.Int8 `json:"tilt,omitempty"`
		} `json:"right,omitempty"`
	} `json:"mirrors,omitempty"`
	// Rainsensor signals.
	Raindetection struct {
		// Rain intensity. 0 = Dry, No Rain. 100 = Covered.
		Intensity null.Uint8 `json:"intensity,omitempty"`
	} `json:"raindetection,omitempty"`
	// Rear spoiler position, 0% = Spoiler fully stowed. 100% = Spoiler fully exposed.
	RearMainSpoilerPosition null.Float64 `json:"rearMainSpoilerPosition,omitempty"`
	// Trunk status.
	Trunk struct {
		// Trunk status.
		Front struct {
			// Is trunk locked or unlocked. True = Locked. False = Unlocked.
			IsLocked null.Bool `json:"isLocked,omitempty"`
			// Trunk open or closed. True = Open. False = Closed.
			IsOpen null.Bool `json:"isOpen,omitempty"`
		} `json:"front,omitempty"`
		// Trunk status.
		Rear struct {
			// Is trunk locked or unlocked. True = Locked. False = Unlocked.
			IsLocked null.Bool `json:"isLocked,omitempty"`
			// Trunk open or closed. True = Open. False = Closed.
			IsOpen null.Bool `json:"isOpen,omitempty"`
		} `json:"rear,omitempty"`
	} `json:"trunk,omitempty"`
	// Windshield signals.
	Windshield struct {
		// Windshield signals.
		Front struct {
			// Windshield heater status. False - off, True - on.
			IsHeatingOn null.Bool `json:"isHeatingOn,omitempty"`
			// Windshield washer fluid signals
			WasherFluid struct {
				// Low level indication for washer fluid. True = Level Low. False = Level OK.
				IsLevelLow null.Bool `json:"isLevelLow,omitempty"`
				// Washer fluid level as a percent. 0 = Empty. 100 = Full.
				Level null.Uint8 `json:"level,omitempty"`
			} `json:"washerFluid,omitempty"`
			// Windshield wiper signals.
			Wiping struct {
				// Relative intensity/sensitivity for interval and rain sensor mode as requested by user/driver. Has no significance if Windshield.Wiping.Mode is OFF/SLOW/MEDIUM/FAST 0 - wipers inactive. 1 - minimum intensity (lowest frequency/sensitivity, longest interval.String). 2/3/4/... - higher intensity (higher frequency/sensitivity, shorter interval.String). Maximum value supported is vehicle specific.
				Intensity null.Uint8 `json:"intensity,omitempty"`
				// Wiper wear status. True = Worn, Replacement recommended or required. False = Not Worn.
				IsWipersWorn null.Bool `json:"isWipersWorn,omitempty"`
				// Wiper mode requested by user/driver. INTERVAL indicates intermittent wiping, with fixed time interval between each wipe. RAIN_SENSOR indicates intermittent wiping based on rain intensity.
				// Must be one of ["OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR"]
				Mode null.String `json:"mode,omitempty"`
				// Signals to control behavior of wipers in detail. By default VSS expects only one instance.
				System struct {
					// Actual position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific.
					ActualPosition null.Float64 `json:"actualPosition,omitempty"`
					// Actual current used by wiper drive.
					DriveCurrent null.Float64 `json:"driveCurrent,omitempty"`
					// Wiping frequency/speed, measured in cycles per minute. The signal concerns the actual speed of the wiper blades when moving. Intervals/pauses are excluded, i.e. the value corresponds to the number of cycles that would be completed in 1 minute if wiping permanently over default range.
					Frequency null.Uint8 `json:"frequency,omitempty"`
					// Indicates if wiper movement is blocked. True = Movement blocked. False = Movement not blocked.
					IsBlocked null.Bool `json:"isBlocked,omitempty"`
					// Indicates if current wipe movement is completed or near completion. True = Movement is completed or near completion. Changes to RequestedPosition will be executed first after reaching previous RequestedPosition, if it has not already been reached. False = Movement is not near completion. Any change to RequestedPosition will be executed immediately. Change of direction may not be allowed.
					IsEndingWipeCycle null.Bool `json:"isEndingWipeCycle,omitempty"`
					// Indicates if wiper system is overheated. True = Wiper system overheated. False = Wiper system not overheated.
					IsOverheated null.Bool `json:"isOverheated,omitempty"`
					// Indicates if a requested position has been reached. IsPositionReached refers to the previous position in case the TargetPosition is updated while IsEndingWipeCycle=True. True = Current or Previous TargetPosition reached. False = Position not (yet) reached, or wipers have moved away from the reached position.
					IsPositionReached null.Bool `json:"isPositionReached,omitempty"`
					// Indicates system failure. True if wiping is disabled due to system failure.
					IsWiperError null.Bool `json:"isWiperError,omitempty"`
					// Indicates wiper movement. True if wiper blades are moving. Change of direction shall be considered as IsWiping if wipers will continue to move directly after the change of direction.
					IsWiping null.Bool `json:"isWiping,omitempty"`
					// Requested mode of wiper system. STOP_HOLD means that the wipers shall move to position given by TargetPosition and then hold the position. WIPE means that wipers shall move to the position given by TargetPosition and then hold the position if no new TargetPosition is requested. PLANT_MODE means that wiping is disabled. Exact behavior is vehicle specific. EMERGENCY_STOP means that wiping shall be immediately stopped without holding the position.
					// Must be one of ["STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP"]
					Mode null.String `json:"mode,omitempty"`
					// Requested position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific. System behavior when receiving TargetPosition depends on Mode and IsEndingWipeCycle. Supported values are vehicle specific and might be dynamically corrected. If IsEndingWipeCycle=True then wipers will complete current movement before actuating new TargetPosition. If IsEndingWipeCycle=False then wipers will directly change destination if the TargetPosition is changed.
					TargetPosition null.Float64 `json:"targetPosition,omitempty"`
				} `json:"system,omitempty"`
				// Wiper wear as percent. 0 = No Wear. 100 = Worn. Replacement required. Method for calculating or estimating wiper wear is vehicle specific. For windshields with multiple wipers the wear reported shall correspond to the most worn wiper.
				WiperWear null.Uint8 `json:"wiperWear,omitempty"`
			} `json:"wiping,omitempty"`
		} `json:"front,omitempty"`
		// Windshield signals.
		Rear struct {
			// Windshield heater status. False - off, True - on.
			IsHeatingOn null.Bool `json:"isHeatingOn,omitempty"`
			// Windshield washer fluid signals
			WasherFluid struct {
				// Low level indication for washer fluid. True = Level Low. False = Level OK.
				IsLevelLow null.Bool `json:"isLevelLow,omitempty"`
				// Washer fluid level as a percent. 0 = Empty. 100 = Full.
				Level null.Uint8 `json:"level,omitempty"`
			} `json:"washerFluid,omitempty"`
			// Windshield wiper signals.
			Wiping struct {
				// Relative intensity/sensitivity for interval and rain sensor mode as requested by user/driver. Has no significance if Windshield.Wiping.Mode is OFF/SLOW/MEDIUM/FAST 0 - wipers inactive. 1 - minimum intensity (lowest frequency/sensitivity, longest interval.String). 2/3/4/... - higher intensity (higher frequency/sensitivity, shorter interval.String). Maximum value supported is vehicle specific.
				Intensity null.Uint8 `json:"intensity,omitempty"`
				// Wiper wear status. True = Worn, Replacement recommended or required. False = Not Worn.
				IsWipersWorn null.Bool `json:"isWipersWorn,omitempty"`
				// Wiper mode requested by user/driver. INTERVAL indicates intermittent wiping, with fixed time interval between each wipe. RAIN_SENSOR indicates intermittent wiping based on rain intensity.
				// Must be one of ["OFF", "SLOW", "MEDIUM", "FAST", "INTERVAL", "RAIN_SENSOR"]
				Mode null.String `json:"mode,omitempty"`
				// Signals to control behavior of wipers in detail. By default VSS expects only one instance.
				System struct {
					// Actual position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific.
					ActualPosition null.Float64 `json:"actualPosition,omitempty"`
					// Actual current used by wiper drive.
					DriveCurrent null.Float64 `json:"driveCurrent,omitempty"`
					// Wiping frequency/speed, measured in cycles per minute. The signal concerns the actual speed of the wiper blades when moving. Intervals/pauses are excluded, i.e. the value corresponds to the number of cycles that would be completed in 1 minute if wiping permanently over default range.
					Frequency null.Uint8 `json:"frequency,omitempty"`
					// Indicates if wiper movement is blocked. True = Movement blocked. False = Movement not blocked.
					IsBlocked null.Bool `json:"isBlocked,omitempty"`
					// Indicates if current wipe movement is completed or near completion. True = Movement is completed or near completion. Changes to RequestedPosition will be executed first after reaching previous RequestedPosition, if it has not already been reached. False = Movement is not near completion. Any change to RequestedPosition will be executed immediately. Change of direction may not be allowed.
					IsEndingWipeCycle null.Bool `json:"isEndingWipeCycle,omitempty"`
					// Indicates if wiper system is overheated. True = Wiper system overheated. False = Wiper system not overheated.
					IsOverheated null.Bool `json:"isOverheated,omitempty"`
					// Indicates if a requested position has been reached. IsPositionReached refers to the previous position in case the TargetPosition is updated while IsEndingWipeCycle=True. True = Current or Previous TargetPosition reached. False = Position not (yet) reached, or wipers have moved away from the reached position.
					IsPositionReached null.Bool `json:"isPositionReached,omitempty"`
					// Indicates system failure. True if wiping is disabled due to system failure.
					IsWiperError null.Bool `json:"isWiperError,omitempty"`
					// Indicates wiper movement. True if wiper blades are moving. Change of direction shall be considered as IsWiping if wipers will continue to move directly after the change of direction.
					IsWiping null.Bool `json:"isWiping,omitempty"`
					// Requested mode of wiper system. STOP_HOLD means that the wipers shall move to position given by TargetPosition and then hold the position. WIPE means that wipers shall move to the position given by TargetPosition and then hold the position if no new TargetPosition is requested. PLANT_MODE means that wiping is disabled. Exact behavior is vehicle specific. EMERGENCY_STOP means that wiping shall be immediately stopped without holding the position.
					// Must be one of ["STOP_HOLD", "WIPE", "PLANT_MODE", "EMERGENCY_STOP"]
					Mode null.String `json:"mode,omitempty"`
					// Requested position of main wiper blade for the wiper system relative to reference position. Location of reference position (0 degrees) and direction of positive/negative degrees is vehicle specific. System behavior when receiving TargetPosition depends on Mode and IsEndingWipeCycle. Supported values are vehicle specific and might be dynamically corrected. If IsEndingWipeCycle=True then wipers will complete current movement before actuating new TargetPosition. If IsEndingWipeCycle=False then wipers will directly change destination if the TargetPosition is changed.
					TargetPosition null.Float64 `json:"targetPosition,omitempty"`
				} `json:"system,omitempty"`
				// Wiper wear as percent. 0 = No Wear. 100 = Worn. Replacement required. Method for calculating or estimating wiper wear is vehicle specific. For windshields with multiple wipers the wear reported shall correspond to the most worn wiper.
				WiperWear null.Uint8 `json:"wiperWear,omitempty"`
			} `json:"wiping,omitempty"`
		} `json:"rear,omitempty"`
	} `json:"windshield,omitempty"`
}

// All in-cabin components, including doors.
type CabinVariables struct {
	// Convertible roof.
	Convertible struct {
		// Roof status on convertible vehicles.
		// Must be one of ["UNDEFINED", "CLOSED", "OPEN", "CLOSING", "OPENING", "STALLED"]
		Status null.String `json:"status,omitempty"`
	} `json:"convertible,omitempty"`
	// All doors, including windows and switches.
	Door struct {
		// All doors, including windows and switches.
		Row1 struct {
			// All doors, including windows and switches.
			Left struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive null.Bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked null.Bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen null.Bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged null.Bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen null.Bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"left,omitempty"`
			// All doors, including windows and switches.
			Right struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive null.Bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked null.Bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen null.Bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged null.Bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen null.Bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"right,omitempty"`
		} `json:"row1,omitempty"`
		// All doors, including windows and switches.
		Row2 struct {
			// All doors, including windows and switches.
			Left struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive null.Bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked null.Bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen null.Bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged null.Bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen null.Bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"left,omitempty"`
			// All doors, including windows and switches.
			Right struct {
				// Is door child lock engaged. True = Engaged. False = Disengaged.
				IsChildLockActive null.Bool `json:"isChildLockActive,omitempty"`
				// Is door locked or unlocked. True = Locked. False = Unlocked.
				IsLocked null.Bool `json:"isLocked,omitempty"`
				// Is door open or closed
				IsOpen null.Bool `json:"isOpen,omitempty"`
				// Side window shade
				Shade struct {
					// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"shade,omitempty"`
				// Door window status
				Window struct {
					// Is window child lock engaged. True = Engaged. False = Disengaged.
					IsChildLockEngaged null.Bool `json:"isChildLockEngaged,omitempty"`
					// Is window open or closed?
					IsOpen null.Bool `json:"isOpen,omitempty"`
					// Window position. 0 = Fully closed 100 = Fully opened.
					Position null.Uint8 `json:"position,omitempty"`
					// Switch controlling sliding action such as window, sunroof, or blind.
					// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
					Switch null.String `json:"switch,omitempty"`
				} `json:"window,omitempty"`
			} `json:"right,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"door,omitempty"`
	// Climate control
	HVAC struct {
		// Ambient air temperature inside the vehicle.
		AmbientAirTemperature null.Float64 `json:"ambientAirTemperature,omitempty"`
		// Is Air conditioning active.
		IsAirConditioningActive null.Bool `json:"isAirConditioningActive,omitempty"`
		// Is front defroster active.
		IsFrontDefrosterActive null.Bool `json:"isFrontDefrosterActive,omitempty"`
		// Is rear defroster active.
		IsRearDefrosterActive null.Bool `json:"isRearDefrosterActive,omitempty"`
		// Is recirculation active.
		IsRecirculationActive null.Bool `json:"isRecirculationActive,omitempty"`
		// HVAC for single station in the vehicle
		Station struct {
			// HVAC for single station in the vehicle
			Row1 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row1,omitempty"`
			// HVAC for single station in the vehicle
			Row2 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row2,omitempty"`
			// HVAC for single station in the vehicle
			Row3 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row3,omitempty"`
			// HVAC for single station in the vehicle
			Row4 struct {
				// HVAC for single station in the vehicle
				Left struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"left,omitempty"`
				// HVAC for single station in the vehicle
				Right struct {
					// Direction of airstream
					// Must be one of ["UP", "MIDDLE", "DOWN"]
					AirDistribution null.String `json:"airDistribution,omitempty"`
					// Fan Speed, 0 = off. 100 = max
					FanSpeed null.Uint8 `json:"fanSpeed,omitempty"`
					// Temperature
					Temperature null.Int8 `json:"temperature,omitempty"`
				} `json:"right,omitempty"`
			} `json:"row4,omitempty"`
		} `json:"station,omitempty"`
	} `json:"hVAC,omitempty"`
	// Infotainment system.
	Infotainment struct {
		// HMI related signals
		HMI struct {
			// ISO 639-1 standard language code for the current HMI
			CurrentLanguage null.String `json:"currentLanguage,omitempty"`
			// Date format used in the current HMI
			// Must be one of ["YYYY_MM_DD", "DD_MM_YYYY", "MM_DD_YYYY", "YY_MM_DD", "DD_MM_YY", "MM_DD_YY"]
			DateFormat null.String `json:"dateFormat,omitempty"`
			// Current display theme
			// Must be one of ["DAY", "NIGHT"]
			DayNightMode null.String `json:"dayNightMode,omitempty"`
			// Distance unit used in the current HMI
			// Must be one of ["MILES", "KILOMETERS"]
			DistanceUnit null.String `json:"distanceUnit,omitempty"`
			// EV fuel economy unit used in the current HMI
			// Must be one of ["MILES_PER_KILOWATT_HOUR", "KILOMETERS_PER_KILOWATT_HOUR", "KILOWATT_HOURS_PER_100_MILES", "KILOWATT_HOURS_PER_100_KILOMETERS", "WATT_HOURS_PER_MILE", "WATT_HOURS_PER_KILOMETER"]
			EVEconomyUnits null.String `json:"eVEconomyUnits,omitempty"`
			// Fuel economy unit used in the current HMI
			// Must be one of ["MPG_UK", "MPG_US", "MILES_PER_LITER", "KILOMETERS_PER_LITER", "LITERS_PER_100_KILOMETERS"]
			FuelEconomyUnits null.String `json:"fuelEconomyUnits,omitempty"`
			// Temperature unit used in the current HMI
			// Must be one of ["C", "F"]
			TemperatureUnit null.String `json:"temperatureUnit,omitempty"`
			// Time format used in the current HMI
			// Must be one of ["HR_12", "HR_24"]
			TimeFormat null.String `json:"timeFormat,omitempty"`
			// Tire pressure unit used in the current HMI
			// Must be one of ["‘PSI’", "‘KPA’", "’BAR’"]
			TirePressureUnit null.String `json:"tirePressureUnit,omitempty"`
		} `json:"hMI,omitempty"`
		// All Media actions
		Media struct {
			// Tells if the media was
			// Must be one of ["UNKNOWN", "STOP", "PLAY", "FAST_FORWARD", "FAST_BACKWARD", "SKIP_FORWARD", "SKIP_BACKWARD"]
			Action null.String `json:"action,omitempty"`
			// URI of suggested media that was declined
			DeclinedURI null.String `json:"declinedURI,omitempty"`
			// Collection of signals updated in concert when a new media is played
			Played struct {
				// Name of album being played
				Album null.String `json:"album,omitempty"`
				// Name of artist being played
				Artist null.String `json:"artist,omitempty"`
				// Current playback rate of media being played.
				PlaybackRate null.Float64 `json:"playbackRate,omitempty"`
				// Media selected for playback
				// Must be one of ["UNKNOWN", "SIRIUS_XM", "AM", "FM", "DAB", "TV", "CD", "DVD", "AUX", "USB", "DISK", "BLUETOOTH", "INTERNET", "VOICE", "BEEP"]
				Source null.String `json:"source,omitempty"`
				// Name of track being played
				Track null.String `json:"track,omitempty"`
				// User Resource associated with the media
				URI null.String `json:"uRI,omitempty"`
			} `json:"played,omitempty"`
			// URI of suggested media that was selected
			SelectedURI null.String `json:"selectedURI,omitempty"`
			// Current Media Volume
			Volume null.Uint8 `json:"volume,omitempty"`
		} `json:"media,omitempty"`
		// All navigation actions
		Navigation struct {
			// A navigation has been selected.
			DestinationSet struct {
				// Latitude of destination in WGS 84 geodetic coordinates.
				Latitude null.Float64 `json:"latitude,omitempty"`
				// Longitude of destination in WGS 84 geodetic coordinates.
				Longitude null.Float64 `json:"longitude,omitempty"`
			} `json:"destinationSet,omitempty"`
			// Navigation mute state that was selected.
			// Must be one of ["MUTED", "ALERT_ONLY", "UNMUTED"]
			Mute null.String `json:"mute,omitempty"`
			// Current navigation volume
			Volume null.Uint8 `json:"volume,omitempty"`
		} `json:"navigation,omitempty"`
		// All smartphone projection actions.
		SmartphoneProjection struct {
			// Projection activation info.
			// Must be one of ["NONE", "ACTIVE", "INACTIVE"]
			Active null.String `json:"active,omitempty"`
			// Connectivity source selected for projection.
			// Must be one of ["USB", "BLUETOOTH", "WIFI"]
			Source null.String `json:"source,omitempty"`
		} `json:"smartphoneProjection,omitempty"`
	} `json:"infotainment,omitempty"`
	// Interior lights signals and sensors.
	Lights struct {
		// How much ambient light is detected in cabin. 0 = No ambient light. 100 = Full brightness
		AmbientLight null.Uint8 `json:"ambientLight,omitempty"`
		// Is central dome light light on
		IsDomeOn null.Bool `json:"isDomeOn,omitempty"`
		// Is glove box light on
		IsGloveBoxOn null.Bool `json:"isGloveBoxOn,omitempty"`
		// Is trunk light light on
		IsTrunkOn null.Bool `json:"isTrunkOn,omitempty"`
		// Intensity of the interior lights. 0 = Off. 100 = Full brightness.
		LightIntensity null.Uint8 `json:"lightIntensity,omitempty"`
		// Spotlight for a specific area in the vehicle.
		Spotlight struct {
			// Spotlight for a specific area in the vehicle.
			Row1 struct {
				// Is light on the left side switched on
				IsLeftOn null.Bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn null.Bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn null.Bool `json:"isSharedOn,omitempty"`
			} `json:"row1,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row2 struct {
				// Is light on the left side switched on
				IsLeftOn null.Bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn null.Bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn null.Bool `json:"isSharedOn,omitempty"`
			} `json:"row2,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row3 struct {
				// Is light on the left side switched on
				IsLeftOn null.Bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn null.Bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn null.Bool `json:"isSharedOn,omitempty"`
			} `json:"row3,omitempty"`
			// Spotlight for a specific area in the vehicle.
			Row4 struct {
				// Is light on the left side switched on
				IsLeftOn null.Bool `json:"isLeftOn,omitempty"`
				// Is light on the right side switched on
				IsRightOn null.Bool `json:"isRightOn,omitempty"`
				// Is a shared light across a specific row on
				IsSharedOn null.Bool `json:"isSharedOn,omitempty"`
			} `json:"row4,omitempty"`
		} `json:"spotlight,omitempty"`
	} `json:"lights,omitempty"`
	// Rear window shade.
	RearShade struct {
		// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
		Position null.Uint8 `json:"position,omitempty"`
		// Switch controlling sliding action such as window, sunroof, or blind.
		// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
		Switch null.String `json:"switch,omitempty"`
	} `json:"rearShade,omitempty"`
	// Rearview mirror.
	RearviewMirror struct {
		// Dimming level of rearview mirror. 0 = undimmed. 100 = fully dimmed.
		DimmingLevel null.Uint8 `json:"dimmingLevel,omitempty"`
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
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos1,omitempty"`
			// All seats.
			Pos2 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos2,omitempty"`
			// All seats.
			Pos3 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos3,omitempty"`
		} `json:"row1,omitempty"`
		// All seats.
		Row2 struct {
			// All seats.
			Pos1 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos1,omitempty"`
			// All seats.
			Pos2 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos2,omitempty"`
			// All seats.
			Pos3 struct {
				// Airbag signals.
				Airbag struct {
					// Airbag deployment status. True = Airbag deployed. False = Airbag not deployed.
					IsDeployed null.Bool `json:"isDeployed,omitempty"`
				} `json:"airbag,omitempty"`
				// Describes signals related to the backrest of the seat.
				Backrest struct {
					// Adjustable lumbar support mechanisms in seats allow the user to change the seat back shape.
					Lumbar struct {
						// Height of lumbar support. Position is relative within available movable range of the lumbar support. 0 = Lowermost position supported.
						Height null.Uint8 `json:"height,omitempty"`
						// Lumbar support (in/out position). 0 = Innermost position. 100 = Outermost position.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"lumbar,omitempty"`
					// Backrest recline compared to seat z-axis (seat vertical axis). 0 degrees = Upright/Vertical backrest. Negative degrees for forward recline. Positive degrees for backward recline.
					Recline null.Float64 `json:"recline,omitempty"`
					// Backrest side bolster (lumbar side support) settings.
					SideBolster struct {
						// Side bolster support. 0 = Minimum support (widest side bolster setting). 100 = Maximum support.
						Support null.Float64 `json:"support,omitempty"`
					} `json:"sideBolster,omitempty"`
				} `json:"backrest,omitempty"`
				// Headrest settings.
				Headrest struct {
					// Headrest angle, relative to backrest, 0 degrees if parallel to backrest, Positive degrees = tilted forward.
					Angle null.Float64 `json:"angle,omitempty"`
					// Position of headrest relative to movable range of the head rest. 0 = Bottommost position supported.
					Height null.Uint8 `json:"height,omitempty"`
				} `json:"headrest,omitempty"`
				// Seat cooling / heating. 0 = off. -100 = max cold. +100 = max heat.
				Heating null.Int8 `json:"heating,omitempty"`
				// Seat position on vehicle z-axis. Position is relative within available movable range of the seating. 0 = Lowermost position supported.
				Height null.Uint16 `json:"height,omitempty"`
				// Is the belt engaged.
				IsBelted null.Bool `json:"isBelted,omitempty"`
				// Does the seat have a passenger in it.
				IsOccupied null.Bool `json:"isOccupied,omitempty"`
				// Seat massage level. 0 = off. 100 = max massage.
				Massage null.Uint8 `json:"massage,omitempty"`
				// Occupant data.
				Occupant struct {
					// Identifier attributes based on OAuth 2.0.
					Identifier struct {
						// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
						Issuer null.String `json:"issuer,omitempty"`
						// Subject for the authentication of the occupant. E.g. UserID 7331677.
						Subject null.String `json:"subject,omitempty"`
					} `json:"identifier,omitempty"`
				} `json:"occupant,omitempty"`
				// Seat position on vehicle x-axis. Position is relative to the frontmost position supported by the seat. 0 = Frontmost position supported.
				Position null.Uint16 `json:"position,omitempty"`
				// Describes signals related to the seat bottom of the seat.
				Seating struct {
					// Length adjustment of seating. 0 = Adjustable part of seating in rearmost position (Shortest length of seating).
					Length null.Uint16 `json:"length,omitempty"`
				} `json:"seating,omitempty"`
				// Seat switch signals
				Switch struct {
					// Describes switches related to the backrest of the seat.
					Backrest struct {
						// Backrest recline backward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineBackwardEngaged null.Bool `json:"isReclineBackwardEngaged,omitempty"`
						// Backrest recline forward switch engaged (SingleSeat.Backrest.Recline).
						IsReclineForwardEngaged null.Bool `json:"isReclineForwardEngaged,omitempty"`
						// Switches for SingleSeat.Backrest.Lumbar.
						Lumbar struct {
							// Lumbar down switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
							// Is switch for less lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more lumbar support engaged (SingleSeat.Backrest.Lumbar.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
							// Lumbar up switch engaged (SingleSeat.Backrest.Lumbar.Support).
							IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
						} `json:"lumbar,omitempty"`
						// Switches for SingleSeat.Backrest.SideBolster.
						SideBolster struct {
							// Is switch for less side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsLessSupportEngaged null.Bool `json:"isLessSupportEngaged,omitempty"`
							// Is switch for more side bolster support engaged (SingleSeat.Backrest.SideBolster.Support).
							IsMoreSupportEngaged null.Bool `json:"isMoreSupportEngaged,omitempty"`
						} `json:"sideBolster,omitempty"`
					} `json:"backrest,omitempty"`
					// Switches for SingleSeat.Headrest.
					Headrest struct {
						// Head rest backward switch engaged (SingleSeat.Headrest.Angle).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Head rest down switch engaged (SingleSeat.Headrest.Height).
						IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
						// Head rest forward switch engaged (SingleSeat.Headrest.Angle).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
						// Head rest up switch engaged (SingleSeat.Headrest.Height).
						IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					} `json:"headrest,omitempty"`
					// Seat backward switch engaged (SingleSeat.Position).
					IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
					// Cooler switch for Seat heater (SingleSeat.Heating).
					IsCoolerEngaged null.Bool `json:"isCoolerEngaged,omitempty"`
					// Seat down switch engaged (SingleSeat.Height).
					IsDownEngaged null.Bool `json:"isDownEngaged,omitempty"`
					// Seat forward switch engaged (SingleSeat.Position).
					IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					// Tilt backward switch engaged (SingleSeat.Tilt).
					IsTiltBackwardEngaged null.Bool `json:"isTiltBackwardEngaged,omitempty"`
					// Tilt forward switch engaged (SingleSeat.Tilt).
					IsTiltForwardEngaged null.Bool `json:"isTiltForwardEngaged,omitempty"`
					// Seat up switch engaged (SingleSeat.Height).
					IsUpEngaged null.Bool `json:"isUpEngaged,omitempty"`
					// Warmer switch for Seat heater (SingleSeat.Heating).
					IsWarmerEngaged null.Bool `json:"isWarmerEngaged,omitempty"`
					// Switches for SingleSeat.Massage.
					Massage struct {
						// Decrease massage level switch engaged (SingleSeat.Massage).
						IsDecreaseEngaged null.Bool `json:"isDecreaseEngaged,omitempty"`
						// Increase massage level switch engaged (SingleSeat.Massage).
						IsIncreaseEngaged null.Bool `json:"isIncreaseEngaged,omitempty"`
					} `json:"massage,omitempty"`
					// Describes switches related to the seating of the seat.
					Seating struct {
						// Is switch to decrease seating length engaged (SingleSeat.Seating.Length).
						IsBackwardEngaged null.Bool `json:"isBackwardEngaged,omitempty"`
						// Is switch to increase seating length engaged (SingleSeat.Seating.Length).
						IsForwardEngaged null.Bool `json:"isForwardEngaged,omitempty"`
					} `json:"seating,omitempty"`
				} `json:"switch,omitempty"`
				// Tilting of seat (seating and backrest) relative to vehicle x-axis. 0 = seat bottom is flat, seat bottom and vehicle x-axis are parallel. Positive degrees = seat tilted backwards, seat x-axis tilted upward, seat z-axis is tilted backward.
				Tilt null.Float64 `json:"tilt,omitempty"`
			} `json:"pos3,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"seat,omitempty"`
	// Sun roof status.
	Sunroof struct {
		// Sunroof position. 0 = Fully closed 100 = Fully opened. -100 = Fully tilted.
		Position null.Int8 `json:"position,omitempty"`
		// Sun roof shade status.
		Shade struct {
			// Position of window blind. 0 = Fully retracted. 100 = Fully deployed.
			Position null.Uint8 `json:"position,omitempty"`
			// Switch controlling sliding action such as window, sunroof, or blind.
			// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN"]
			Switch null.String `json:"switch,omitempty"`
		} `json:"shade,omitempty"`
		// Switch controlling sliding action such as window, sunroof, or shade.
		// Must be one of ["INACTIVE", "CLOSE", "OPEN", "ONE_SHOT_CLOSE", "ONE_SHOT_OPEN", "TILT_UP", "TILT_DOWN"]
		Switch null.String `json:"switch,omitempty"`
	} `json:"sunroof,omitempty"`
}

// All data concerning steering, suspension, wheels, and brakes.
type ChassisVariables struct {
	// Accelerator signals
	Accelerator struct {
		// Accelerator pedal position as percent. 0 = Not depressed. 100 = Fully depressed.
		PedalPosition null.Uint8 `json:"pedalPosition,omitempty"`
	} `json:"accelerator,omitempty"`
	// Axle signals
	Axle struct {
		// Axle signals
		Row1 struct {
			// Wheel signals for axle
			Wheel struct {
				// Wheel signals for axle
				Left struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel null.Uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn null.Bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow null.Bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear null.Uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed null.Float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow null.Bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure null.Uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature null.Float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"left,omitempty"`
				// Wheel signals for axle
				Right struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel null.Uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn null.Bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow null.Bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear null.Uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed null.Float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow null.Bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure null.Uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature null.Float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"right,omitempty"`
			} `json:"wheel,omitempty"`
		} `json:"row1,omitempty"`
		// Axle signals
		Row2 struct {
			// Wheel signals for axle
			Wheel struct {
				// Wheel signals for axle
				Left struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel null.Uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn null.Bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow null.Bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear null.Uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed null.Float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow null.Bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure null.Uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature null.Float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"left,omitempty"`
				// Wheel signals for axle
				Right struct {
					// Brake signals for wheel
					Brake struct {
						// Brake fluid level as percent. 0 = Empty. 100 = Full.
						FluidLevel null.Uint8 `json:"fluidLevel,omitempty"`
						// Brake pad wear status. True = Worn. False = Not Worn.
						IsBrakesWorn null.Bool `json:"isBrakesWorn,omitempty"`
						// Brake fluid level status. True = Brake fluid level low. False = Brake fluid level OK.
						IsFluidLevelLow null.Bool `json:"isFluidLevelLow,omitempty"`
						// Brake pad wear as percent. 0 = No Wear. 100 = Worn.
						PadWear null.Uint8 `json:"padWear,omitempty"`
					} `json:"brake,omitempty"`
					// Rotational speed of a vehicle's wheel.
					Speed null.Float64 `json:"speed,omitempty"`
					// Tire signals for wheel.
					Tire struct {
						// Tire Pressure Status. True = Low tire pressure. False = Good tire pressure.
						IsPressureLow null.Bool `json:"isPressureLow,omitempty"`
						// Tire pressure in kilo-Pascal.
						Pressure null.Uint16 `json:"pressure,omitempty"`
						// Tire temperature in Celsius.
						Temperature null.Float64 `json:"temperature,omitempty"`
					} `json:"tire,omitempty"`
				} `json:"right,omitempty"`
			} `json:"wheel,omitempty"`
		} `json:"row2,omitempty"`
	} `json:"axle,omitempty"`
	// Brake system signals
	Brake struct {
		// Indicates if emergency braking initiated by driver is detected. True = Emergency braking detected. False = Emergency braking not detected.
		IsDriverEmergencyBrakingDetected null.Bool `json:"isDriverEmergencyBrakingDetected,omitempty"`
		// Brake pedal position as percent. 0 = Not depressed. 100 = Fully depressed.
		PedalPosition null.Uint8 `json:"pedalPosition,omitempty"`
	} `json:"brake,omitempty"`
	// Parking brake signals
	ParkingBrake struct {
		// Parking brake status. True = Parking Brake is Engaged. False = Parking Brake is not Engaged.
		IsEngaged null.Bool `json:"isEngaged,omitempty"`
	} `json:"parkingBrake,omitempty"`
	// Steering wheel signals
	SteeringWheel struct {
		// Steering wheel angle. Positive = degrees to the left. Negative = degrees to the right.
		Angle null.Int16 `json:"angle,omitempty"`
		// Steering wheel column extension from dashboard. 0 = Closest to dashboard. 100 = Furthest from dashboard.
		Extension null.Uint8 `json:"extension,omitempty"`
		// Steering wheel column tilt. 0 = Lowest position. 100 = Highest position.
		Tilt null.Uint8 `json:"tilt,omitempty"`
	} `json:"steeringWheel,omitempty"`
}

// Connectivity data.
type ConnectivityVariables struct {
	// Indicates if connectivity between vehicle and cloud is available. True = Connectivity is available. False = Connectivity is not available.
	IsConnectivityAvailable null.Bool `json:"isConnectivityAvailable,omitempty"`
}

// The current latitude and longitude of the vehicle.
type CurrentLocationVariables struct {
	// Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	Altitude null.Float64 `json:"altitude,omitempty"`
	// Information on the GNSS receiver used for determining current location.
	GNSSReceiver struct {
		// Fix status of GNSS receiver.
		// Must be one of ["NONE", "TWO_D", "TWO_D_SATELLITE_BASED_AUGMENTATION", "TWO_D_GROUND_BASED_AUGMENTATION", "TWO_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION", "THREE_D", "THREE_D_SATELLITE_BASED_AUGMENTATION", "THREE_D_GROUND_BASED_AUGMENTATION", "THREE_D_SATELLITE_AND_GROUND_BASED_AUGMENTATION"]
		FixType null.String `json:"fixType,omitempty"`
		// Mounting position of GNSS receiver antenna relative to vehicle coordinate system. Axis definitions according to ISO 8855. Origin at center of (first) rear axle.
		MountingPosition struct {
		} `json:"mountingPosition,omitempty"`
	} `json:"gNSSReceiver,omitempty"`
	// Current heading relative to geographic north. 0 = North, 90 = East, 180 = South, 270 = West.
	Heading null.Float64 `json:"heading,omitempty"`
	// Accuracy of the latitude and longitude coordinates.
	HorizontalAccuracy null.Float64 `json:"horizontalAccuracy,omitempty"`
	// Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	Latitude null.Float64 `json:"latitude,omitempty"`
	// Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	Longitude null.Float64 `json:"longitude,omitempty"`
	// Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	Timestamp null.String `json:"timestamp,omitempty"`
	// Accuracy of altitude.
	VerticalAccuracy null.Float64 `json:"verticalAccuracy,omitempty"`
}

// Driver data.
type DriverVariables struct {
	// Probability of attentiveness of the driver.
	AttentiveProbability null.Float64 `json:"attentiveProbability,omitempty"`
	// Distraction level of the driver will be the level how much the driver is distracted, by multiple factors. E.g. Driving situation, acustical or optical signales inside the cockpit, phone calls.
	DistractionLevel null.Float64 `json:"distractionLevel,omitempty"`
	// Fatigueness level of driver. Evaluated by multiple factors like trip time, behaviour of steering, eye status.
	FatigueLevel null.Float64 `json:"fatigueLevel,omitempty"`
	// Heart rate of the driver.
	HeartRate null.Uint16 `json:"heartRate,omitempty"`
	// Identifier attributes based on OAuth 2.0.
	Identifier struct {
		// Unique Issuer for the authentication of the occupant. E.g. https://accounts.funcorp.com.
		Issuer null.String `json:"issuer,omitempty"`
		// Subject for the authentication of the occupant. E.g. UserID 7331677.
		Subject null.String `json:"subject,omitempty"`
	} `json:"identifier,omitempty"`
	// Has driver the eyes on road or not?
	IsEyesOnRoad null.Bool `json:"isEyesOnRoad,omitempty"`
}

// Information about exterior measured by vehicle.
type ExteriorVariables struct {
	// Air temperature outside the vehicle.
	AirTemperature null.Float64 `json:"airTemperature,omitempty"`
	// Relative humidity outside the vehicle. 0 = Dry, 100 = Air fully saturated.
	Humidity null.Float64 `json:"humidity,omitempty"`
	// Light intensity outside the vehicle. 0 = No light detected, 100 = Fully lit.
	LightIntensity null.Float64 `json:"lightIntensity,omitempty"`
}

// Signals related to low voltage battery.
type LowVoltageBatteryVariables struct {
	// Current current flowing in/out of the low voltage battery. Positive = Current flowing in to battery, e.g. during charging or driving. Negative = Current flowing out of battery, e.g. when using the battery to start a combustion engine.
	CurrentCurrent null.Float64 `json:"currentCurrent,omitempty"`
	// Current Voltage of the low voltage battery.
	CurrentVoltage null.Float64 `json:"currentVoltage,omitempty"`
}

// OBD data.
type OBDVariables struct {
	// PID 43 - Absolute load value
	AbsoluteLoad null.Float64 `json:"absoluteLoad,omitempty"`
	// PID 49 - Accelerator pedal position D
	AcceleratorPositionD null.Float64 `json:"acceleratorPositionD,omitempty"`
	// PID 4A - Accelerator pedal position E
	AcceleratorPositionE null.Float64 `json:"acceleratorPositionE,omitempty"`
	// PID 4B - Accelerator pedal position F
	AcceleratorPositionF null.Float64 `json:"acceleratorPositionF,omitempty"`
	// PID 12 - Secondary air status
	AirStatus null.String `json:"airStatus,omitempty"`
	// PID 46 - Ambient air temperature
	AmbientAirTemperature null.Float64 `json:"ambientAirTemperature,omitempty"`
	// PID 33 - Barometric pressure
	BarometricPressure null.Float64 `json:"barometricPressure,omitempty"`
	// Catalyst signals
	Catalyst struct {
		// Catalyst bank 1 signals
		Bank1 struct {
			// PID 3C - Catalyst temperature from bank 1, sensor 1
			Temperature1 null.Float64 `json:"temperature1,omitempty"`
			// PID 3E - Catalyst temperature from bank 1, sensor 2
			Temperature2 null.Float64 `json:"temperature2,omitempty"`
		} `json:"bank1,omitempty"`
		// Catalyst bank 2 signals
		Bank2 struct {
			// PID 3D - Catalyst temperature from bank 2, sensor 1
			Temperature1 null.Float64 `json:"temperature1,omitempty"`
			// PID 3F - Catalyst temperature from bank 2, sensor 2
			Temperature2 null.Float64 `json:"temperature2,omitempty"`
		} `json:"bank2,omitempty"`
	} `json:"catalyst,omitempty"`
	// PID 2C - Commanded exhaust gas recirculation (EGR)
	CommandedEGR null.Float64 `json:"commandedEGR,omitempty"`
	// PID 2E - Commanded evaporative purge (EVAP) valve
	CommandedEVAP null.Float64 `json:"commandedEVAP,omitempty"`
	// PID 44 - Commanded equivalence ratio
	CommandedEquivalenceRatio null.Float64 `json:"commandedEquivalenceRatio,omitempty"`
	// PID 42 - Control module voltage
	ControlModuleVoltage null.Float64 `json:"controlModuleVoltage,omitempty"`
	// PID 05 - Coolant temperature
	CoolantTemperature null.Float64 `json:"coolantTemperature,omitempty"`
	// List of currently active DTCs formatted according OBD II (SAE-J2012DA_201812) standard ([P|C|B|U]XXXXX )
	DTCList []null.String `json:"dTCList,omitempty"`
	// PID 31 - Distance traveled since codes cleared
	DistanceSinceDTCClear null.Float64 `json:"distanceSinceDTCClear,omitempty"`
	// PID 21 - Distance traveled with MIL on
	DistanceWithMIL null.Float64 `json:"distanceWithMIL,omitempty"`
	// PID 41 - OBD status for the current drive cycle
	DriveCycleStatus struct {
		// Number of sensor Trouble Codes (DTC)
		DTCCount null.Uint8 `json:"dTCCount,omitempty"`
		// Type of the ignition for ICE - spark = spark plug ignition, compression = self-igniting (Diesel engines)
		// Must be one of ["SPARK", "COMPRESSION"]
		IgnitionType null.String `json:"ignitionType,omitempty"`
		// Malfunction Indicator Light (MIL) - False = Off, True = On
		IsMILOn null.Bool `json:"isMILOn,omitempty"`
	} `json:"driveCycleStatus,omitempty"`
	// PID 2D - Exhaust gas recirculation (EGR) error
	EGRError null.Float64 `json:"eGRError,omitempty"`
	// PID 32 - Evaporative purge (EVAP) system pressure
	EVAPVaporPressure null.Float64 `json:"eVAPVaporPressure,omitempty"`
	// PID 53 - Absolute evaporative purge (EVAP) system pressure
	EVAPVaporPressureAbsolute null.Float64 `json:"eVAPVaporPressureAbsolute,omitempty"`
	// PID 54 - Alternate evaporative purge (EVAP) system pressure
	EVAPVaporPressureAlternate null.Float64 `json:"eVAPVaporPressureAlternate,omitempty"`
	// PID 04 - Engine load in percent - 0 = no load, 100 = full load
	EngineLoad null.Float64 `json:"engineLoad,omitempty"`
	// PID 0C - Engine speed measured as rotations per minute
	EngineSpeed null.Float64 `json:"engineSpeed,omitempty"`
	// PID 52 - Percentage of ethanol in the fuel
	EthanolPercent null.Float64 `json:"ethanolPercent,omitempty"`
	// PID 02 - DTC that triggered the freeze frame
	FreezeDTC null.String `json:"freezeDTC,omitempty"`
	// PID 5D - Fuel injection timing
	FuelInjectionTiming null.Float64 `json:"fuelInjectionTiming,omitempty"`
	// PID 2F - Fuel level in the fuel tank
	FuelLevel null.Float64 `json:"fuelLevel,omitempty"`
	// PID 0A - Fuel pressure
	FuelPressure null.Float64 `json:"fuelPressure,omitempty"`
	// PID 59 - Absolute fuel rail pressure
	FuelRailPressureAbsolute null.Float64 `json:"fuelRailPressureAbsolute,omitempty"`
	// PID 23 - Fuel rail pressure direct inject
	FuelRailPressureDirect null.Float64 `json:"fuelRailPressureDirect,omitempty"`
	// PID 22 - Fuel rail pressure relative to vacuum
	FuelRailPressureVac null.Float64 `json:"fuelRailPressureVac,omitempty"`
	// PID 5E - Engine fuel rate
	FuelRate null.Float64 `json:"fuelRate,omitempty"`
	// PID 03 - Fuel status
	FuelStatus null.String `json:"fuelStatus,omitempty"`
	// PID 51 - Fuel type
	FuelType null.String `json:"fuelType,omitempty"`
	// PID 5B - Remaining life of hybrid battery
	HybridBatteryRemaining null.Float64 `json:"hybridBatteryRemaining,omitempty"`
	// PID 0F - Intake temperature
	IntakeTemp null.Float64 `json:"intakeTemp,omitempty"`
	// PID 1E - Auxiliary input status (power take off)
	IsPTOActive null.Bool `json:"isPTOActive,omitempty"`
	// PID 07 - Long Term (learned) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
	LongTermFuelTrim1 null.Float64 `json:"longTermFuelTrim1,omitempty"`
	// PID 09 - Long Term (learned) Fuel Trim - Bank 2 - negative percent leaner, positive percent richer
	LongTermFuelTrim2 null.Float64 `json:"longTermFuelTrim2,omitempty"`
	// PID 56 (byte A) - Long term secondary O2 trim - Bank 1
	LongTermO2Trim1 null.Float64 `json:"longTermO2Trim1,omitempty"`
	// PID 58 (byte A) - Long term secondary O2 trim - Bank 2
	LongTermO2Trim2 null.Float64 `json:"longTermO2Trim2,omitempty"`
	// PID 56 (byte B) - Long term secondary O2 trim - Bank 3
	LongTermO2Trim3 null.Float64 `json:"longTermO2Trim3,omitempty"`
	// PID 58 (byte B) - Long term secondary O2 trim - Bank 4
	LongTermO2Trim4 null.Float64 `json:"longTermO2Trim4,omitempty"`
	// PID 10 - Grams of air drawn into engine per second
	MAF null.Float64 `json:"mAF,omitempty"`
	// PID 0B - Intake manifold pressure
	MAP null.Float64 `json:"mAP,omitempty"`
	// PID 50 - Maximum flow for mass air flow sensor
	MaxMAF null.Float64 `json:"maxMAF,omitempty"`
	// Oxygen sensors (PID 14 - PID 1B)
	O2 struct {
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor1 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor1,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor2 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor2,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor3 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor3,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor4 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor4,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor5 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor5,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor6 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor6,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor7 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor7,omitempty"`
		// Oxygen sensors (PID 14 - PID 1B)
		Sensor8 struct {
			// PID 1x (byte B) - Short term fuel trim
			ShortTermFuelTrim null.Float64 `json:"shortTermFuelTrim,omitempty"`
			// PID 1x (byte A) - Sensor voltage
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor8,omitempty"`
	} `json:"o2,omitempty"`
	// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
	O2WR struct {
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor1 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor1,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor2 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor2,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor3 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor3,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor4 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor4,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor5 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor5,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor6 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor6,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor7 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor7,omitempty"`
		// Wide range/band oxygen sensors (PID 24 - 2B and PID 34 - 3B)
		Sensor8 struct {
			// PID 3x (byte CD) - Current for wide range/band oxygen sensor
			Current null.Float64 `json:"current,omitempty"`
			// PID 2x (byte AB) and PID 3x (byte AB) - Lambda for wide range/band oxygen sensor
			Lambda null.Float64 `json:"lambda,omitempty"`
			// PID 2x (byte CD) - Voltage for wide range/band oxygen sensor
			Voltage null.Float64 `json:"voltage,omitempty"`
		} `json:"sensor8,omitempty"`
	} `json:"o2WR,omitempty"`
	// PID 5C - Engine oil temperature
	OilTemperature null.Float64 `json:"oilTemperature,omitempty"`
	// PID 13 - Presence of oxygen sensors in 2 banks. [A0..A3] == Bank 1, Sensors 1-4. [A4..A7] == Bank 2, Sensors 1-4
	OxygenSensorsIn2Banks null.Uint8 `json:"oxygenSensorsIn2Banks,omitempty"`
	// PID 1D - Presence of oxygen sensors in 4 banks. Similar to PID 13, but [A0..A7] == [B1S1, B1S2, B2S1, B2S2, B3S1, B3S2, B4S1, B4S2]
	OxygenSensorsIn4Banks null.Uint8 `json:"oxygenSensorsIn4Banks,omitempty"`
	// PID 00 - Bit array of the supported pids 01 to 20
	PidsA null.Uint32 `json:"pidsA,omitempty"`
	// PID 20 - Bit array of the supported pids 21 to 40
	PidsB null.Uint32 `json:"pidsB,omitempty"`
	// PID 40 - Bit array of the supported pids 41 to 60
	PidsC null.Uint32 `json:"pidsC,omitempty"`
	// PID 5A - Relative accelerator pedal position
	RelativeAcceleratorPosition null.Float64 `json:"relativeAcceleratorPosition,omitempty"`
	// PID 45 - Relative throttle position
	RelativeThrottlePosition null.Float64 `json:"relativeThrottlePosition,omitempty"`
	// PID 1F - Engine run time
	RunTime null.Float64 `json:"runTime,omitempty"`
	// PID 4D - Run time with MIL on
	RunTimeMIL null.Float64 `json:"runTimeMIL,omitempty"`
	// PID 06 - Short Term (immediate) Fuel Trim - Bank 1 - negative percent leaner, positive percent richer
	ShortTermFuelTrim1 null.Float64 `json:"shortTermFuelTrim1,omitempty"`
	// PID 08 - Short Term (immediate) Fuel Trim - Bank 2 - negative percent leaner, positive percent richer
	ShortTermFuelTrim2 null.Float64 `json:"shortTermFuelTrim2,omitempty"`
	// PID 55 (byte A) - Short term secondary O2 trim - Bank 1
	ShortTermO2Trim1 null.Float64 `json:"shortTermO2Trim1,omitempty"`
	// PID 57 (byte A) - Short term secondary O2 trim - Bank 2
	ShortTermO2Trim2 null.Float64 `json:"shortTermO2Trim2,omitempty"`
	// PID 55 (byte B) - Short term secondary O2 trim - Bank 3
	ShortTermO2Trim3 null.Float64 `json:"shortTermO2Trim3,omitempty"`
	// PID 57 (byte B) - Short term secondary O2 trim - Bank 4
	ShortTermO2Trim4 null.Float64 `json:"shortTermO2Trim4,omitempty"`
	// PID 0D - Vehicle speed
	Speed null.Float64 `json:"speed,omitempty"`
	// PID 01 - OBD status
	Status struct {
		// Number of sensor Trouble Codes (DTC)
		DTCCount null.Uint8 `json:"dTCCount,omitempty"`
		// Type of the ignition for ICE - spark = spark plug ignition, compression = self-igniting (Diesel engines)
		// Must be one of ["SPARK", "COMPRESSION"]
		IgnitionType null.String `json:"ignitionType,omitempty"`
		// Malfunction Indicator Light (MIL) False = Off, True = On
		IsMILOn null.Bool `json:"isMILOn,omitempty"`
	} `json:"status,omitempty"`
	// PID 4C - Commanded throttle actuator
	ThrottleActuator null.Float64 `json:"throttleActuator,omitempty"`
	// PID 11 - Throttle position - 0 = closed throttle, 100 = open throttle
	ThrottlePosition null.Float64 `json:"throttlePosition,omitempty"`
	// PID 47 - Absolute throttle position B
	ThrottlePositionB null.Float64 `json:"throttlePositionB,omitempty"`
	// PID 48 - Absolute throttle position C
	ThrottlePositionC null.Float64 `json:"throttlePositionC,omitempty"`
	// PID 4E - Time since trouble codes cleared
	TimeSinceDTCCleared null.Float64 `json:"timeSinceDTCCleared,omitempty"`
	// PID 0E - Time advance
	TimingAdvance null.Float64 `json:"timingAdvance,omitempty"`
	// PID 30 - Number of warm-ups since codes cleared
	WarmupsSinceDTCClear null.Uint8 `json:"warmupsSinceDTCClear,omitempty"`
}

// Powertrain data for battery management, etc.
type PowertrainVariables struct {
	// The accumulated energy from regenerative braking over lifetime.
	AccumulatedBrakingEnergy null.Float64 `json:"accumulatedBrakingEnergy,omitempty"`
	// Engine-specific data, stopping at the bell housing.
	CombustionEngine struct {
		// Signals related to Diesel Exhaust Fluid (DEF). DEF is called AUS32 in ISO 22241.
		DieselExhaustFluid struct {
			// Indicates if the Diesel Exhaust Fluid level is low. True if level is low. Definition of low is vehicle dependent.
			IsLevelLow null.Bool `json:"isLevelLow,omitempty"`
			// Level of the Diesel Exhaust Fluid tank as percent of capacity. 0 = empty. 100 = full.
			Level null.Uint8 `json:"level,omitempty"`
			// Remaining range in meters of the Diesel Exhaust Fluid present in the vehicle.
			Range null.Uint32 `json:"range,omitempty"`
		} `json:"dieselExhaustFluid,omitempty"`
		// Diesel Particulate Filter signals.
		DieselParticulateFilter struct {
			// Delta Pressure of Diesel Particulate Filter.
			DeltaPressure null.Float64 `json:"deltaPressure,omitempty"`
			// Inlet temperature of Diesel Particulate Filter.
			InletTemperature null.Float64 `json:"inletTemperature,omitempty"`
			// Outlet temperature of Diesel Particulate Filter.
			OutletTemperature null.Float64 `json:"outletTemperature,omitempty"`
		} `json:"dieselParticulateFilter,omitempty"`
		// Engine coolant temperature.
		ECT null.Int16 `json:"eCT,omitempty"`
		// Engine oil pressure.
		EOP null.Uint16 `json:"eOP,omitempty"`
		// Engine oil temperature.
		EOT null.Int16 `json:"eOT,omitempty"`
		// Accumulated time during engine lifetime with 'engine speed (rpm) > 0'.
		EngineHours null.Float64 `json:"engineHours,omitempty"`
		// Engine oil level.
		// Must be one of ["CRITICALLY_LOW", "LOW", "NORMAL", "HIGH", "CRITICALLY_HIGH"]
		EngineOilLevel null.String `json:"engineOilLevel,omitempty"`
		// Accumulated idling time during engine lifetime. Definition of idling is not standardized.
		IdleHours null.Float64 `json:"idleHours,omitempty"`
		// Engine Running. True if engine is rotating (Speed > 0).
		IsRunning null.Bool `json:"isRunning,omitempty"`
		// Grams of air drawn into engine per second.
		MAF null.Uint16 `json:"mAF,omitempty"`
		// Manifold absolute pressure possibly boosted using forced induction.
		MAP null.Uint16 `json:"mAP,omitempty"`
		// Remaining engine oil life in seconds. Negative values can be used to indicate that lifetime has been exceeded.
		OilLifeRemaining null.Int32 `json:"oilLifeRemaining,omitempty"`
		// Current engine power output. Shall be reported as 0 during engine breaking.
		Power null.Uint16 `json:"power,omitempty"`
		// Engine speed measured as rotations per minute.
		Speed null.Uint16 `json:"speed,omitempty"`
		// Current throttle position.
		TPS null.Uint8 `json:"tPS,omitempty"`
		// Current engine torque. Shall be reported as 0 during engine breaking.
		Torque null.Uint16 `json:"torque,omitempty"`
	} `json:"combustionEngine,omitempty"`
	// Electric Motor specific data.
	ElectricMotor struct {
		// Motor coolant temperature (if applicable).
		CoolantTemperature null.Int16 `json:"coolantTemperature,omitempty"`
		// Current motor power output. Negative values indicate regen mode.
		Power null.Int16 `json:"power,omitempty"`
		// Motor rotational speed measured as rotations per minute. Negative values indicate reverse driving mode.
		Speed null.Int32 `json:"speed,omitempty"`
		// Motor temperature.
		Temperature null.Int16 `json:"temperature,omitempty"`
		// Current motor torque. Negative values indicate regen mode.
		Torque null.Int16 `json:"torque,omitempty"`
	} `json:"electricMotor,omitempty"`
	// Fuel system data.
	FuelSystem struct {
		// Average consumption in liters per 100 km.
		AverageConsumption null.Float64 `json:"averageConsumption,omitempty"`
		// Fuel amount in liters consumed since start of current trip.
		ConsumptionSinceStart null.Float64 `json:"consumptionSinceStart,omitempty"`
		// Current consumption in liters per 100 km.
		InstantConsumption null.Float64 `json:"instantConsumption,omitempty"`
		// Indicates whether eco start stop is currently enabled.
		IsEngineStopStartEnabled null.Bool `json:"isEngineStopStartEnabled,omitempty"`
		// Indicates that the fuel level is low (e.g. <50km range).
		IsFuelLevelLow null.Bool `json:"isFuelLevelLow,omitempty"`
		// Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
		Level null.Uint8 `json:"level,omitempty"`
		// Remaining range in meters using only liquid fuel.
		Range null.Uint32 `json:"range,omitempty"`
		// Time in seconds elapsed since start of current trip.
		TimeSinceStart null.Uint32 `json:"timeSinceStart,omitempty"`
	} `json:"fuelSystem,omitempty"`
	// Remaining range in meters using all energy sources available in the vehicle.
	Range null.Uint32 `json:"range,omitempty"`
	// Battery Management data.
	TractionBattery struct {
		// The accumulated energy delivered to the battery during charging over lifetime of the battery.
		AccumulatedChargedEnergy null.Float64 `json:"accumulatedChargedEnergy,omitempty"`
		// The accumulated charge throughput delivered to the battery during charging over lifetime of the battery.
		AccumulatedChargedThroughput null.Float64 `json:"accumulatedChargedThroughput,omitempty"`
		// The accumulated energy leaving HV battery for propulsion and auxiliary loads over lifetime of the battery.
		AccumulatedConsumedEnergy null.Float64 `json:"accumulatedConsumedEnergy,omitempty"`
		// The accumulated charge throughput leaving HV battery for propulsion and auxiliary loads over lifetime of the battery.
		AccumulatedConsumedThroughput null.Float64 `json:"accumulatedConsumedThroughput,omitempty"`
		// Properties related to battery charging.
		Charging struct {
			// Current charging current.
			ChargeCurrent struct {
				// Current DC charging current at inlet. Negative if returning energy to grid.
				DC null.Float64 `json:"dC,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 1. Negative if returning energy to grid.
				Phase1 null.Float64 `json:"phase1,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 2. Negative if returning energy to grid.
				Phase2 null.Float64 `json:"phase2,omitempty"`
				// Current AC charging current (rms) at inlet for Phase 3. Negative if returning energy to grid.
				Phase3 null.Float64 `json:"phase3,omitempty"`
			} `json:"chargeCurrent,omitempty"`
			// Target charge limit (state of charge) for battery.
			ChargeLimit null.Uint8 `json:"chargeLimit,omitempty"`
			// Status of the charge port cover, can potentially be controlled manually.
			// Must be one of ["OPEN", "CLOSED"]
			ChargePortFlap null.String `json:"chargePortFlap,omitempty"`
			// Current charging rate, as in kilometers of range added per hour.
			ChargeRate null.Float64 `json:"chargeRate,omitempty"`
			// Current charging voltage, as measured at the charging inlet.
			ChargeVoltage struct {
				// Current DC charging voltage at charging inlet.
				DC null.Float64 `json:"dC,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 1.
				Phase1 null.Float64 `json:"phase1,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 2.
				Phase2 null.Float64 `json:"phase2,omitempty"`
				// Current AC charging voltage (rms) at inlet for Phase 3.
				Phase3 null.Float64 `json:"phase3,omitempty"`
			} `json:"chargeVoltage,omitempty"`
			// True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
			IsCharging null.Bool `json:"isCharging,omitempty"`
			// Indicates if a charging cable is physically connected to the vehicle or not.
			IsChargingCableConnected null.Bool `json:"isChargingCableConnected,omitempty"`
			// Is charging cable locked to prevent removal.
			IsChargingCableLocked null.Bool `json:"isChargingCableLocked,omitempty"`
			// True if discharging (vehicle to grid) is ongoing. Discharging is considered to be ongoing if energy is flowing from vehicle to charger/grid.
			IsDischarging null.Bool `json:"isDischarging,omitempty"`
			// Maximum charging current that can be accepted by the system, as measured at the charging inlet.
			MaximumChargingCurrent struct {
				// Maximum DC charging current at inlet that can be accepted by the system.
				DC null.Float64 `json:"dC,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 1 that can be accepted by the system.
				Phase1 null.Float64 `json:"phase1,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 2 that can be accepted by the system.
				Phase2 null.Float64 `json:"phase2,omitempty"`
				// Maximum AC charging current (rms) at inlet for Phase 3 that can be accepted by the system.
				Phase3 null.Float64 `json:"phase3,omitempty"`
			} `json:"maximumChargingCurrent,omitempty"`
			// Control of the charge process. MANUAL means manually initiated (plug-in event, companion app, etc). TIMER means timer-based. GRID means grid-controlled (eg ISO 15118). PROFILE means controlled by profile download to vehicle.
			// Must be one of ["MANUAL", "TIMER", "GRID", "PROFILE"]
			Mode null.String `json:"mode,omitempty"`
			// Electrical energy lost by power dissipation to heat inside the AC/DC converter.
			PowerLoss null.Float64 `json:"powerLoss,omitempty"`
			// Start or stop the charging process.
			// Must be one of ["START", "STOP"]
			StartStopCharging null.String `json:"startStopCharging,omitempty"`
			// Current temperature of AC/DC converter converting grid voltage to battery voltage.
			Temperature null.Float64 `json:"temperature,omitempty"`
			// The time needed for the current charging process to reach Charging.ChargeLimit. 0 if charging is complete or no charging process is active or planned.
			TimeToComplete null.Uint32 `json:"timeToComplete,omitempty"`
			// Properties related to timing of battery charging sessions.
			Timer struct {
				// Defines timer mode for charging: INACTIVE - no timer set, charging may start as soon as battery is connected to a charger. START_TIME - charging shall start at Charging.Timer.Time. END_TIME - charging shall be finished (reach Charging.ChargeLimit) at Charging.Timer.Time. When charging is completed the vehicle shall change mode to 'inactive' or set a new Charging.Timer.Time. Charging shall start immediately if mode is 'starttime' or 'endtime' and Charging.Timer.Time is a time in the past.
				// Must be one of ["INACTIVE", "START_TIME", "END_TIME"]
				Mode null.String `json:"mode,omitempty"`
				// Time for next charging-related action, formatted according to ISO 8601 with UTC time zone. Value has no significance if Charging.Timer.Mode is 'inactive'.
				Time null.String `json:"time,omitempty"`
			} `json:"timer,omitempty"`
		} `json:"charging,omitempty"`
		// Current current flowing in/out of battery. Positive = Current flowing in to battery, e.g. during charging. Negative = Current flowing out of battery, e.g. during driving.
		CurrentCurrent null.Float64 `json:"currentCurrent,omitempty"`
		// Current electrical energy flowing in/out of battery. Positive = Energy flowing in to battery, e.g. during charging. Negative = Energy flowing out of battery, e.g. during driving.
		CurrentPower null.Float64 `json:"currentPower,omitempty"`
		// Current Voltage of the battery.
		CurrentVoltage null.Float64 `json:"currentVoltage,omitempty"`
		// Properties related to DC/DC converter converting high voltage (from high voltage battery) to vehicle low voltage (supply voltage, typically 12 Volts).
		DCDC struct {
			// Electrical energy lost by power dissipation to heat inside DC/DC converter.
			PowerLoss null.Float64 `json:"powerLoss,omitempty"`
			// Current temperature of DC/DC converter converting battery high voltage to vehicle low voltage (typically 12 Volts).
			Temperature null.Float64 `json:"temperature,omitempty"`
		} `json:"dCDC,omitempty"`
		// Indicating if the ground (negative terminator) of the traction battery is connected to the powertrain.
		IsGroundConnected null.Bool `json:"isGroundConnected,omitempty"`
		// Indicating if the power (positive terminator) of the traction battery is connected to the powertrain.
		IsPowerConnected null.Bool `json:"isPowerConnected,omitempty"`
		// Total net capacity of the battery considering aging.
		NetCapacity null.Uint16 `json:"netCapacity,omitempty"`
		// Electrical energy lost by power dissipation to heat inside the battery.
		PowerLoss null.Float64 `json:"powerLoss,omitempty"`
		// Remaining range in meters using only battery.
		Range null.Uint32 `json:"range,omitempty"`
		// Information on the state of charge of the vehicle's high voltage battery.
		StateOfCharge struct {
			// Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
			Current null.Float64 `json:"current,omitempty"`
			// State of charge displayed to the customer.
			Displayed null.Float64 `json:"displayed,omitempty"`
		} `json:"stateOfCharge,omitempty"`
		// Calculated battery state of health at standard conditions.
		StateOfHealth null.Float64 `json:"stateOfHealth,omitempty"`
		// Temperature Information for the battery pack.
		Temperature struct {
			// Current average temperature of the battery cells.
			Average null.Float64 `json:"average,omitempty"`
			// Current maximum temperature of the battery cells, i.e. temperature of the hottest cell.
			Max null.Float64 `json:"max,omitempty"`
			// Current minimum temperature of the battery cells, i.e. temperature of the coldest cell.
			Min null.Float64 `json:"min,omitempty"`
		} `json:"temperature,omitempty"`
	} `json:"tractionBattery,omitempty"`
	// Transmission-specific data, stopping at the drive shafts.
	Transmission struct {
		// Clutch engagement. 0% = Clutch fully disengaged. 100% = Clutch fully engaged.
		ClutchEngagement null.Float64 `json:"clutchEngagement,omitempty"`
		// Clutch wear as a percent. 0 = no wear. 100 = worn.
		ClutchWear null.Uint8 `json:"clutchWear,omitempty"`
		// The current gear. 0=Neutral, 1/2/..=Forward, -1/-2/..=Reverse.
		CurrentGear null.Int8 `json:"currentGear,omitempty"`
		// Front Diff Lock engagement. 0% = Diff lock fully disengaged. 100% = Diff lock fully engaged.
		DiffLockFrontEngagement null.Float64 `json:"diffLockFrontEngagement,omitempty"`
		// Rear Diff Lock engagement. 0% = Diff lock fully disengaged. 100% = Diff lock fully engaged.
		DiffLockRearEngagement null.Float64 `json:"diffLockRearEngagement,omitempty"`
		// Is the gearbox in automatic or manual (paddle) mode.
		// Must be one of ["MANUAL", "AUTOMATIC"]
		GearChangeMode null.String `json:"gearChangeMode,omitempty"`
		// Is electrical powertrain mechanically connected/engaged to the drivetrain or not. False = Disconnected/Disengaged. True = Connected/Engaged.
		IsElectricalPowertrainEngaged null.Bool `json:"isElectricalPowertrainEngaged,omitempty"`
		// Is gearbox in low range mode or not. False = Normal/High range engaged. True = Low range engaged.
		IsLowRangeEngaged null.Bool `json:"isLowRangeEngaged,omitempty"`
		// Is the transmission park lock engaged or not. False = Disengaged. True = Engaged.
		IsParkLockEngaged null.Bool `json:"isParkLockEngaged,omitempty"`
		// Current gearbox performance mode.
		// Must be one of ["NORMAL", "SPORT", "ECONOMY", "SNOW", "RAIN"]
		PerformanceMode null.String `json:"performanceMode,omitempty"`
		// The selected gear. 0=Neutral, 1/2/..=Forward, -1/-2/..=Reverse, 126=Park, 127=Drive.
		SelectedGear null.Int8 `json:"selectedGear,omitempty"`
		// The current gearbox temperature.
		Temperature null.Int16 `json:"temperature,omitempty"`
		// Torque distribution between front and rear axle in percent. -100% = Full torque to front axle, 0% = 50:50 Front/Rear, 100% = Full torque to rear axle.
		TorqueDistribution null.Float64 `json:"torqueDistribution,omitempty"`
		// Odometer reading, total distance travelled during the lifetime of the transmission.
		TravelledDistance null.Float64 `json:"travelledDistance,omitempty"`
	} `json:"transmission,omitempty"`
}

// Service data.
type ServiceVariables struct {
	// Remaining distance to service (of any kind). Negative values indicate service overdue.
	DistanceToService null.Float64 `json:"distanceToService,omitempty"`
	// Indicates if vehicle needs service (of any kind). True = Service needed now or in the near future. False = No known need for service.
	IsServiceDue null.Bool `json:"isServiceDue,omitempty"`
	// Remaining time to service (of any kind). Negative values indicate service overdue.
	TimeToService null.Int32 `json:"timeToService,omitempty"`
}

// Trailer signals.
type TrailerVariables struct {
	// Signal indicating if trailer is connected or not.
	IsConnected null.Bool `json:"isConnected,omitempty"`
}

type DataSchemaStruct struct {
	// High-level vehicle data.
	Vehicle struct {
		// All Advanced Driver Assist Systems data.
		ADAS AdvancedDriveAssistVariables `json:"advancedDriveAssist,omitempty"`
		// Spatial acceleration. Axis definitions according to ISO 8855.
		Acceleration AccelerationVariables `json:"acceleration,omitempty"`
		// Spatial rotation. Axis definitions according to ISO 8855.
		AngularVelocity AngularVelocityVariables `json:"angularVelocity,omitempty"`
		// Average speed for the current trip.
		AverageSpeed null.Float64 `json:"averageSpeed,omitempty"`
		// All body components.
		Body BodyVariables `json:"body,omitempty"`
		// All in-cabin components, including doors.
		Cabin CabinVariables `json:"cabin,omitempty"`
		// All data concerning steering, suspension, wheels, and brakes.
		Chassis ChassisVariables `json:"chassis,omitempty"`
		// Connectivity data.
		Connectivity ConnectivityVariables `json:"connectivity,omitempty"`
		// The current latitude and longitude of the vehicle.
		CurrentLocation CurrentLocationVariables `json:"currentLocation,omitempty"`
		// Current overall Vehicle weight. Including passengers, cargo and other load inside the car.
		CurrentOverallWeight null.Uint16 `json:"currentOverallWeight,omitempty"`
		// Driver data.
		Driver DriverVariables `json:"driver,omitempty"`
		// Information about exterior measured by vehicle.
		Exterior ExteriorVariables `json:"exterior,omitempty"`
		// Vehicle breakdown or any similar event causing vehicle to stop on the road, that might pose a risk to other road users. True = Vehicle broken down on the road, due to e.g. engine problems, flat tire, out of gas, brake problems. False = Vehicle not broken down.
		IsBrokenDown null.Bool `json:"isBrokenDown,omitempty"`
		// Indicates whether the vehicle is stationary or moving.
		IsMoving null.Bool `json:"isMoving,omitempty"`
		// Signals related to low voltage battery.
		LowVoltageBattery LowVoltageBatteryVariables `json:"lowVoltageBattery,omitempty"`
		// State of the supply voltage of the control units (usually 12V).
		// Must be one of ["UNDEFINED", "LOCK", "OFF", "ACC", "ON", "START"]
		LowVoltageSystemState null.String `json:"lowVoltageSystemState,omitempty"`
		// OBD data.
		OBD OBDVariables `json:"oBD,omitempty"`
		// Powertrain data for battery management, etc.
		Powertrain PowertrainVariables `json:"powertrain,omitempty"`
		// Service data.
		Service ServiceVariables `json:"service,omitempty"`
		// Vehicle speed.
		Speed null.Float64 `json:"speed,omitempty"`
		// Trailer signals.
		Trailer TrailerVariables `json:"trailer,omitempty"`
		// Odometer reading, total distance travelled during the lifetime of the vehicle.
		TravelledDistance null.Float64 `json:"travelledDistance,omitempty"`
		// Current trip meter reading.

		TripMeterReading null.Float64 `json:"tripMeterReading,omitempty"`
	} `json:"vehicle,omitempty"`
}

func NewVehicleStatus() DataSchemaStruct {
	return DataSchemaStruct{}
}

func NewVehicleAttributes() VehicleAttributes {
	return VehicleAttributes{}
}

func (d DataSchemaStruct) Marshal() ([]byte, error) {
	if valid, err := d.IsValid(); !valid {
		return []byte{}, err
	}
	return json.Marshal(d)
}

func (v VehicleAttributes) Marshal() ([]byte, error) {
	if valid, err := v.IsValid(); !valid {
		return []byte{}, err
	}
	return json.Marshal(v)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
