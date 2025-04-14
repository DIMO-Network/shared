package redis

// Settings add these redis settings to your settings struct
type Settings struct {
	// URL eg. localhost:6379
	URL      string `yaml:"URL"`
	Password string `yaml:"PASSWORD"`
	// TLS enable tls
	TLS bool `yaml:"TLS"`
	// KeyPrefix prepends this string to all keys. Makes it easier with many services sharing redis cluster. If empty no prefix.
	KeyPrefix string `yaml:"KEY_PREFIX"`
}
