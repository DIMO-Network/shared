package shared

// isOdometerValid encapsulates logic to decide whether to return odometer
func isOdometerValid(odometer float64) bool {
	if odometer <= 100 {
		return false
	}
	if odometer == 65539 || odometer == 65538 {
		return false
	}
	return true
}
