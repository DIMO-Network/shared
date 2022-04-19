package shared

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadFromYaml(t *testing.T) {
	var data = `
PORT: 3000
DB_CONNECT_STRING: mydb.aws.net
ENV: dev
`
	settings, err := loadFromYaml[TestSettings]([]byte(data))
	assert.NoError(t, err, "no error expected")
	assert.NotNilf(t, settings, "settings not expected to be nil")
	assert.Equal(t, 3000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
	assert.Equal(t, "dev", settings.Env)
}

func Test_loadFromEnvVars(t *testing.T) {
	settings := TestSettings{
		Port:            3000,
		DbConnectString: "mydb.aws.net",
		Env:             "dev",
	}
	// these will now override the above
	os.Setenv("ENV", "test")
	os.Setenv("PORT", "5000")

	err := loadFromEnvVars(&settings) // b/c of type inference we don't need to specify the type
	assert.NoError(t, err)
	assert.NotNilf(t, settings, "expected not nil")
	assert.Equal(t, "test", settings.Env)
	assert.Equal(t, 5000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
}

func Test_loadFromEnvVars_errIfNotPointer(t *testing.T) {
	s := TestSettings{}
	err := loadFromEnvVars(s)
	assert.Error(t, err, "expected error if settings not a pointer")
}

type TestSettings struct {
	Port            int    `yaml:"PORT"`
	DbConnectString string `yaml:"DB_CONNECT_STRING"`
	Env             string `yaml:"ENV"`
}
