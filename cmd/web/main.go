package main

import (
	"flag"
	"log"
	"x-bank-users/config"
	"x-bank-users/core/web"
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

	service := web.NewService(&knife, &knife, &knife, &knife)
	transport := http.NewTransport(service)

	errCh := transport.Start(*addr)

	log.Println(<-errCh)
}
