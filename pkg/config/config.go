package config

type Config struct {
	DatabaseURI  string
	DatabaseName string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURI:  "mongodb://mongodb:27017/",
		DatabaseName: "feature_toggle",
	}
}
