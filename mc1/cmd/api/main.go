package main

import (
	"flag"

	"github.com/wugalde19/pratik/mc1/config"
	"github.com/wugalde19/pratik/mc1/pkg/api"
)

var (
	environment string
)

func main() {
	flag.StringVar(&environment, "env", "", "Set the environment (testing, development, production)")
	flag.Parse()

	config, err := config.Load(environment)
	if err != nil {
		panic(err.Error())
	}

	api.Start(config)
}
