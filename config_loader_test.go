package shared

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadFromYaml(t *testing.T) {
	var data = `
PORT: 3000
DB_CONNECT_STRING: mydb.aws.net
ENV: dev
REDIS:
  URL: redis.bobby
INLINE_URL: inline.bobby
`
	settings, err := loadFromYaml[TestSettings]([]byte(data))
	assert.NoError(t, err, "no error expected")
	assert.NotNilf(t, settings, "settings not expected to be nil")
	assert.Equal(t, 3000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
	assert.Equal(t, "dev", settings.Env)
	assert.Equal(t, "redis.bobby", settings.Redis.URL)
	assert.Equal(t, "inline.bobby", settings.Inline.URL)
}

func Test_loadFromEnvVars(t *testing.T) {
	settings := TestSettings{
		Port:            3000,
		DbConnectString: "mydb.aws.net",
		Env:             "dev",
	}
	// these will now override the above
	t.Setenv("ENV", "test")
	t.Setenv("PORT", "5000")
	t.Setenv("REDIS_URL", "redis.bobby")
	t.Setenv("INLINE_URL", "inline.bobby")

	err := loadFromEnvVars(&settings) // b/c of type inference we don't need to specify the type
	assert.NoError(t, err)
	assert.NotNilf(t, settings, "expected not nil")
	assert.Equal(t, "test", settings.Env)
	assert.Equal(t, 5000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
	assert.Equal(t, "redis.bobby", settings.Redis.URL)
	assert.Equal(t, "inline.bobby", settings.Inline.URL)
}

func Test_loadFromEnvVars_errIfNotPointer(t *testing.T) {
	s := TestSettings{}
	err := loadFromEnvVars(s)
	assert.Error(t, err, "expected error if settings not a pointer")
}

func Test_matchEnvVarToField(t *testing.T) {
	settings := &TestSettings{}
	t.Setenv("PORT", "5000")

	valueOfConfig := reflect.ValueOf(settings).Elem()
	field := valueOfConfig.Field(0)

	err := matchEnvVarToField("PORT", field)
	assert.NoError(t, err)

	assert.Equal(t, int64(5000), field.Int())
	assert.Equal(t, 5000, settings.Port)
}

type TestSettings struct {
	Port            int           `yaml:"PORT"`
	DbConnectString string        `yaml:"DB_CONNECT_STRING"`
	Env             string        `yaml:"ENV"`
	Redis           RedisSubProp  `yaml:"REDIS"`
	Inline          InlineSubProp `yaml:",inline"`
}

type RedisSubProp struct {
	URL string `yaml:"URL"`
}

type InlineSubProp struct {
	URL string `yaml:"INLINE_URL"`
}
