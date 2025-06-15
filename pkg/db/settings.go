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
	User               string `env:"USER"                 yaml:"USER"`
	Password           string `env:"PASSWORD"             yaml:"PASSWORD"`
	Port               string `env:"PORT"                 yaml:"PORT"`
	Host               string `env:"HOST"                 yaml:"HOST"`
	Name               string `env:"NAME"                 yaml:"NAME"`
	MaxOpenConnections int    `env:"MAX_OPEN_CONNECTIONS" yaml:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int    `env:"MAX_IDLE_CONNECTIONS" yaml:"MAX_IDLE_CONNECTIONS"`
	SSLMode            string `env:"SSL_MODE"             yaml:"SSL_MODE"`
}

// BuildConnectionString builds the connection string to the database - for now same as reader
func (app *Settings) BuildConnectionString(withSearchPath bool) string {
	sslMode := SSLModeDisable
	if app.SSLMode != "" {
		sslMode = app.SSLMode
	}
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
