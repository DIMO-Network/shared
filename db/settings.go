package db

import "fmt"

const (
	SSLModeDisable    = "disable"
	SSLModeAllow      = "allow"
	SSLModePrefer     = "prefer"
	SSLModeRequire    = "require"
	SSLModeVerifyCA   = "verify-ca"
	SSLModeVerifyFull = "verify-full"
)

// Settings connection settings to postgres db
type Settings struct {
	User               string `yaml:"USER"`
	Password           string `yaml:"PASSWORD"`
	Port               string `yaml:"PORT"`
	Host               string `yaml:"HOST"`
	Name               string `yaml:"NAME"`
	MaxOpenConnections int    `yaml:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `yaml:"MAX_IDLE_CONNECTIONS"`
	SSLMode            string `yaml:"SSL_MODE"`
}

// BuildConnectionString builds the connection string to the database - for now same as reader
func (app *Settings) BuildConnectionString(withSearchPath bool) string {
	sslMode := SSLModeDisable
	if app.SSLMode != "" {
		sslMode = app.SSLMode
	}
	cs := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslMode=%s",
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
