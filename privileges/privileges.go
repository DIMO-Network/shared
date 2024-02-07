// Package privileges contains the different privileges that can be given to a user.
package privileges

// Privilege is a type for the different privileges that can be given.
type Privilege int64

const (
	// NonLocationData provides access to all data excluding location data.
	NonLocationData Privilege = 1

	// Commands provides access to all commands that can be sent to the device.
	Commands Privilege = 2

	// CurrentLocation provides access to the current location of the device.
	CurrentLocation Privilege = 3

	// AllTimeLocation provives access to current and historical location data.
	AllTimeLocation Privilege = 4
)

// String returns the string representation of a privilege.
func (p Privilege) String() string {
	switch p {
	case NonLocationData:
		return "NonLocationData"
	case Commands:
		return "Commands"
	case CurrentLocation:
		return "CurrentLocation"
	case AllTimeLocation:
		return "AllTimeLocation"
	default:
		return "Unknown"
	}
}
