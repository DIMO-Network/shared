package shared

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func Test_loadFromYaml(t *testing.T) {
	var data = `
PORT: 3000
DB_CONNECT_STRING: mydb.aws.net
ENV: dev
REDIS:
  URL: https://k8s-cluster-redis.dimo.zone:8776/redis
INLINE_URL: inline.bobby
IGNORE: ignoreme
ACTUAL_URL_OBJECT: https://identity-api.dimo.zone/query
MINT_VEHICLE_NFT_CONTRACT: 0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8
`
	settings, err := loadFromYaml[TestSettings]([]byte(data))
	require.NoError(t, err, "no error expected")
	require.NotNilf(t, settings, "settings not expected to be nil")
	assert.Equal(t, 3000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
	assert.Equal(t, "dev", settings.Env)
	assert.Equal(t, "https://k8s-cluster-redis.dimo.zone:8776/redis", settings.Redis.URL.String())
	assert.Equal(t, "inline.bobby", settings.Inline.URL)
	assert.Equal(t, "", settings.IGNORE)
	assert.Equal(t, "https://identity-api.dimo.zone/query", settings.ActualURLObject.String())
	assert.Equal(t, "/query", settings.ActualURLObject.Path)
	assert.Equal(t, "identity-api.dimo.zone", settings.ActualURLObject.Host)
	assert.Equal(t, "0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8", settings.MintVehicleNFTContract.String())
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
	t.Setenv("REDIS_URL", "https://k8s-cluster-redis.dimo.zone:8776/redis")
	t.Setenv("INLINE_URL", "inline.bobby")
	t.Setenv("IGNORE", "ignoreme")
	t.Setenv("ACTUAL_URL_OBJECT", "https://identity-api.dimo.zone/query")
	t.Setenv("MINT_VEHICLE_NFT_CONTRACT", "0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8")

	err := loadFromEnvVars(&settings) // b/c of type inference we don't need to specify the type
	require.NoError(t, err)
	assert.NotNilf(t, settings, "expected not nil")
	assert.Equal(t, "test", settings.Env)
	assert.Equal(t, 5000, settings.Port)
	assert.Equal(t, "mydb.aws.net", settings.DbConnectString)
	assert.Equal(t, "https://k8s-cluster-redis.dimo.zone:8776/redis", settings.Redis.URL.String())
	assert.Equal(t, "inline.bobby", settings.Inline.URL)
	assert.Equal(t, "", settings.IGNORE)
	assert.Equal(t, "https://identity-api.dimo.zone/query", settings.ActualURLObject.String())
	assert.Equal(t, "/query", settings.ActualURLObject.Path)
	assert.Equal(t, "identity-api.dimo.zone", settings.ActualURLObject.Host)
	assert.Equal(t, "0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8", settings.MintVehicleNFTContract.String())
}

func Test_loadFromEnvVars_errIfNotPointer(t *testing.T) {
	s := TestSettings{}
	err := loadFromEnvVars(s)
	assert.Error(t, err, "expected error if settings not a pointer")
}

func Test_matchEnvVarToField_PortInt(t *testing.T) {
	settings := &TestSettings{}
	t.Setenv("PORT", "5000")
	t.Setenv("ACTUAL_URL_OBJECT", "https://identity-api.dimo.zone/query")

	valueOfConfig := reflect.ValueOf(settings).Elem()
	field := valueOfConfig.Field(0)

	err := matchEnvVarToField("PORT", field)
	assert.NoError(t, err)

	assert.Equal(t, int64(5000), field.Int())
	assert.Equal(t, 5000, settings.Port)
}

func Test_matchEnvVarToField_Url(t *testing.T) {
	settings := &TestSettings{}
	t.Setenv("ACTUAL_URL_OBJECT", "https://identity-api.dimo.zone/query")

	valueOfConfig := reflect.ValueOf(settings).Elem()
	field := valueOfConfig.Field(6)

	err := matchEnvVarToField("ACTUAL_URL_OBJECT", field)
	assert.NoError(t, err)

	//assert.Equal(t, "https://identity-api.dimo.zone/query", field.String())
	assert.Equal(t, "https://identity-api.dimo.zone/query", settings.ActualURLObject.String())
	assert.Equal(t, "/query", settings.ActualURLObject.Path)
	assert.Equal(t, "identity-api.dimo.zone", settings.ActualURLObject.Host)
}

type TestSettings struct {
	Port                   int            `yaml:"PORT"`
	DbConnectString        string         `yaml:"DB_CONNECT_STRING"`
	Env                    string         `yaml:"ENV"`
	Redis                  RedisSubProp   `yaml:"REDIS"`
	Inline                 InlineSubProp  `yaml:",inline"`
	IGNORE                 string         `yaml:"-"`
	ActualURLObject        url.URL        `yaml:"ACTUAL_URL_OBJECT"`
	MintVehicleNFTContract common.Address `yaml:"MINT_VEHICLE_NFT_CONTRACT"`
}

type RedisSubProp struct {
	URL url.URL `yaml:"URL"`
}

type InlineSubProp struct {
	URL string `yaml:"INLINE_URL"`
}
