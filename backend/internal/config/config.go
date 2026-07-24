package config

type Config struct {
	Port        string
	DatabasePath string
	MockData    bool
}

func Load() *Config {
	return &Config{
		Port:         "8080",
		DatabasePath: "./data/lakenet.db",
		MockData:    true,
	}
}