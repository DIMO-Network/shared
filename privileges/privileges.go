// Package privileges contains the different privileges that can be given to a user.
// Privilege names and values defined here come from values stored on chain.
package privileges

const (
	// VehicleNonLocationData provides access to all data excluding location data.
	VehicleNonLocationData int64 = 1

	// VehicleCommands provides access to all commands that can be sent to the device.
	VehicleCommands int64 = 2

	// VehicleCurrentLocation provides access to the current location of the device.
	VehicleCurrentLocation int64 = 3

	// VehicleAllTimeLocation provives access to current and historical location data.
	VehicleAllTimeLocation int64 = 4
)

// PrivilegeString returns the string representation of a privilege.
func PrivilegeString(p int64) string {
	switch p {
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
