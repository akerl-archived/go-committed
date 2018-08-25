package main

import (
	"log"

	"github.com/akerl/go-lambda/s3"
	"github.com/aws/aws-lambda-go/lambda"
)

type config struct {
}

var c config

func loadConfig() {
	cf, err := s3.GetConfigFromEnv(&c)
	if err != nil {
		panic(err)
	}
	log.Print("Loaded config")
	cf.OnError = func(_ *s3.ConfigFile, err error) {
		log.Printf("Error reloading config: %s", err)
	}
	cf.Autoreload(60)
}

func main() {
	loadConfig()
	lambda.Start(handler)
}

func handler() error {
	return nil
}
