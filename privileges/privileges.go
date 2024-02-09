// Package privileges contains the different privileges that can be given to a user.
// Privilege names and values defined here come from values stored on chain.
package privileges

// Privilege is an enum that represents the different privileges that can be given to a user.
type Privilege int64

const (
	// ManufacureMintDevice provides access to minting a device.
	ManufacureMintDevice Privilege = 1
	// ManufactureDistributeDevice provides access to distributing a device.
	ManufactureDistributeDevice Privilege = 2
	// ManufactureFactoryReset provides access to factory resetting a device.
	ManufactureFactoryReset Privilege = 3

	// Vehicle Privileges.

	// VehicleNonLocationData provides access to all data excluding location data.
	VehicleNonLocationData Privilege = 1
	// VehicleCommands provides access to all commands that can be sent to the device.
	VehicleCommands Privilege = 2
	// VehicleCurrentLocation provides access to the current location of the device.
	VehicleCurrentLocation Privilege = 3
	// VehicleAllTimeLocation provives access to current and historical location data.
	VehicleAllTimeLocation Privilege = 4
)

// String returns the string representation of a VehiclePrivilege.
func (v Privilege) String() string {
	switch v {
	case VehicleNonLocationData:
		return "NonLocationData"
	case VehicleCommands:
		return "Commands"
	case VehicleCurrentLocation:
		return "CurrentLocation"
	case VehicleAllTimeLocation:
		return "AllTimeLocation"
	default:
		return "Unknown"
	}
}
