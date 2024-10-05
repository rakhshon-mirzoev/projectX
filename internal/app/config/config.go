package config

type Config struct {
	Port       string `toml:"port"`
	Host       string `toml:"host"`
	PortPgres  string `toml:"portpgres"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	Dbname     string `toml:"dbname"`
	TimeZone   string `toml:"TimeZone"`
	Sslmode    string `toml:"sslmode"`
	InfoLevel  string `toml:"info_level"`
	DebugLevel string `toml:"debug_level"`
	ErrorLevel string `toml:"error_level"`
}

func NewConfig() *Config {
	return &Config{
		Port:       "8089",
		InfoLevel:  "info",
		DebugLevel: "debug",
		ErrorLevel: "error",
	}
}
