package main

import (
	"flag"

	"github.com/wugalde19/pratik/mc1/config"
)

var (
	environment string
)

func main() {
	flag.StringVar(&environment, "env", "", "Set the environment (testing, development, production)")
	flag.Parse()

	_, err := config.Load(environment)
	if err != nil {
		panic(err.Error())
	}
}
