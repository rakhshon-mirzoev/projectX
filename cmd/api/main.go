package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"main.go/internal/app/api"
	"main.go/logger"
)

var (
	configPathToml string = "configs/api.toml"
	// configPathEnv  string = "configs/.env"
)

func init() {
	flag.StringVar(&configPathToml, "path", "configs/api.toml", "values written while starting service in cmd.")
	// flag.StringVar(&configPathEnv, "path", "configs/.env", "values written while starting service in cmd.")
}

func main() {
	logger.Init()

	flag.Parse()

	config := api.NewConfig()
	_, err := toml.DecodeFile(configPathToml, config)
	if err != nil {
		logger.Warn.Println("Using default values of configs.", err)
	}

	server := api.New(config)

	logger.Error.Fatal(server.Start())

	logger.Info.Println("We started!", config.Port)
}
