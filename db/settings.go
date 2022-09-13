package db

import "fmt"

// Settings connection settings to postgres db
type Settings struct {
	DBUser               string `yaml:"USER"`
	DBPassword           string `yaml:"PASSWORD"`
	DBPort               string `yaml:"PORT"`
	DBHost               string `yaml:"HOST"`
	DBName               string `yaml:"NAME"`
	DBMaxOpenConnections int    `yaml:"MAX_OPEN_CONNECTIONS"`
	DBMaxIdleConnections int    `yaml:"MAX_IDLE_CONNECTIONS"`
}

// BuildConnectionString builds the connection string to the database - for now same as reader
func (app *Settings) BuildConnectionString(withSearchPath bool) string {
	cs := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		app.DBUser,
		app.DBPassword,
		app.DBName,
		app.DBHost,
		app.DBPort,
	)
	if withSearchPath {
		cs = fmt.Sprintf("%s search_path=%s", cs, app.DBName) // assumption is schema has same name as dbname
	}
	return cs
}
