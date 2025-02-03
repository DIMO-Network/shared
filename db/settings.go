package db

import "fmt"

// Settings connection settings to postgres db
type Settings struct {
	User               string `yaml:"USER"`
	Password           string `yaml:"PASSWORD"`
	Port               string `yaml:"PORT"`
	Host               string `yaml:"HOST"`
	Name               string `yaml:"NAME"`
	MaxOpenConnections int    `yaml:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `yaml:"MAX_IDLE_CONNECTIONS"`
}

// BuildConnectionString builds the connection string to the database - for now same as reader
func (app *Settings) BuildConnectionString(withSearchPath bool, sslMode SSLMode) string {
	cs := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		app.User,
		app.Password,
		app.Name,
		app.Host,
		app.Port,
		sslMode,
	)
	if withSearchPath {
		cs = fmt.Sprintf("%s search_path=%s", cs, app.Name) // assumption is schema has same name as dbname
	}
	return cs
}

// SSLMode represents the different PostgreSQL SSL modes
type SSLMode string

const (
	SSLModeDisable SSLMode = "disable"
	SSLModeAllow   SSLMode = "allow"
	// SSLModePrefer falls back to no SSL if can't connect with SSL
	SSLModePrefer     SSLMode = "prefer"
	SSLModeRequire    SSLMode = "require"
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)
