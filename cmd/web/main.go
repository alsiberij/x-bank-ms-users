package main

import (
	"flag"
	"log"
	"x-bank-users/config"
	"x-bank-users/core/web"
	"x-bank-users/infra/hasher"
	"x-bank-users/infra/random"
	"x-bank-users/infra/swissknife"
	"x-bank-users/transport/http"
)

var (
	addr       = flag.String("addr", ":8080", "")
	configFile = flag.String("config", "config.json", "")
)

func main() {
	flag.Parse()

	_, err := config.Read(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	knife := swissknife.NewService()
	passwordHasher := hasher.NewService()
	randomGenerator := random.NewService()
	service := web.NewService(&knife, &randomGenerator, &knife, &knife, &passwordHasher)
	transport := http.NewTransport(service)

	errCh := transport.Start(*addr)

	log.Println(<-errCh)
}
