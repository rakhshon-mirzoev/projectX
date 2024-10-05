package api

type Config struct {
	Port       string `toml:"port"`
	InfoLevel  string `toml:"info_level"`
	DebugLevel string `toml:"debug_level"`
	ErrorLevel string `toml:"error_level"`
}

func NewConfig() *Config {
	return &Config{
		Port:       "8080",
		InfoLevel:  "info",
		DebugLevel: "debug",
		ErrorLevel: "error",
	}
}
