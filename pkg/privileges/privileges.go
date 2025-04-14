// Package privileges contains the different privileges that can be given to a user.
// Privilege names and values defined here come from values stored on chain.
package privileges

// Privilege is an enum that represents the different privileges that can be given to a user.
type Privilege int64

const (
	// Manufacturer Privileges.

	// ManufacturerMintDevice provides access to minting a device.
	ManufacturerMintDevice Privilege = 1
	// ManufacturerDistributeDevice provides access to distributing a device.
	ManufacturerDistributeDevice Privilege = 2
	// ManufacturerFactoryReset provides access to factory resetting a device.
	ManufacturerFactoryReset Privilege = 3
	// ManufacturerDeviceReprovision provides access to force remint aftermarket device.
	ManufacturerDeviceReprovision Privilege = 4
	// ManufacturerDeviceDefinitionInsert provides access to add device definitions on chain.
	ManufacturerDeviceDefinitionInsert Privilege = 5
	// ManufacturerDeviceLastSeen provides access start of time block when device last transmitted data.
	ManufacturerDeviceLastSeen Privilege = 6

	// Vehicle Privileges.

	// VehicleNonLocationData provides access to all data excluding location data.
	VehicleNonLocationData Privilege = 1
	// VehicleCommands provides access to all commands that can be sent to the device.
	VehicleCommands Privilege = 2
	// VehicleCurrentLocation provides access to the current location of the device.
	VehicleCurrentLocation Privilege = 3
	// VehicleAllTimeLocation provives access to current and historical location data.
	VehicleAllTimeLocation Privilege = 4
	// VehicleVinCredential provides access to the VIN Verified Credential.
	VehicleVinCredential Privilege = 5
	// VehicleSubscribeLiveDataPrivilege allows a user to subscribe to live data from the vehicle.
	VehicleSubscribeLiveDataPrivilege = 6
	// VehicleRawData provides access to the raw data from the vehicle.
	VehicleRawData = 7
	// VehicleApproximateLocation provides access to the approximate location of the vehicle.
	VehicleApproximateLocation = 8
)
