package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/wugalde19/pratik/mc2/config"
	"github.com/wugalde19/pratik/mc2/pkg/api"
)

var (
	environment string
)

func main() {
	flag.StringVar(&environment, "env", "", "Set the environment (testing, development, production)")
	flag.Parse()

	envVariablesFileName := fmt.Sprintf("./%s.env", environment)
	if err := godotenv.Load(envVariablesFileName); err != nil {
		panic(fmt.Errorf("no .env file found: %s", err.Error()))
	}

	config, err := config.Load(environment)
	if err != nil {
		panic(err.Error())
	}

	api.Start(config)
}
