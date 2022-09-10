package redis

// Settings add these redis settings to your settings struct
type Settings struct {
	// URL eg. localhost:6379
	URL      string `yaml:"URL"`
	Password string `yaml:"PASSWORD"`
	// TLS enable tls
	TLS bool `yaml:"TLS"`
	// DB 0-16 db instance to use from redis
	DB int `yaml:"DB"`
}
