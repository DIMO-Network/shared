package logfields

import "github.com/rs/zerolog"

// Logger is a zerolog Wrapper so that we can add some convenience methods
type Logger struct {
	zerolog.Logger
}

// WithVehicleTokenID adds the vehicle token id as a log field with correct data type
//
// Example:
//
// baseLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
//
//	logger := mylog.Logger{baseLogger}
//
//	logger.
//	    WithVehicleTokenID(uint64(1234)).
//	    Info().
//	    Msg("user logged in")
func (l Logger) WithVehicleTokenID(tokenID uint64) Logger {
	return Logger{l.Logger.With().Uint64(VehicleTokenID, tokenID).Logger()}
}
