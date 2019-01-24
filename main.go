package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/akerl/go-lambda/apigw/events"
	"github.com/akerl/go-lambda/mux"
	"github.com/akerl/go-lambda/s3"
)

var (
	smsRegex     = regexp.MustCompile(`^/sms$`)
	userRegex    = regexp.MustCompile(`^/user/([\w-]+)$`)
	defaultRegex = regexp.MustCompile(`^/.*$`)
)

type config struct {
	DefaultUser string `json:"default_user"`
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
	d := mux.NewDispatcher(
		mux.NewRoute(smsRegex, smsHandler),
		mux.NewRoute(userRegex, userHandler),
		mux.NewRoute(defaultRegex, defaultHandler),
	)
	mux.Start(d)
}

func defaultHandler(req events.Request) (events.Response, error) {
	host := req.Headers["Host"]
	target := fmt.Sprintf("https://%s/user/%s", host, c.DefaultUser)
	return events.Redirect(target, 307)
}

func userHandler(req events.Request) (events.Response, error) {
	return events.Succeed("Placeholder")
}

func smsHandler(req events.Request) (events.Response, error) {
	return events.Succeed("Placeholder")
}
