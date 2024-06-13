package main

import (
	"flag"
	"log"
	"x-bank-users/config"
	"x-bank-users/core/web"
	"x-bank-users/transport/http"
)

var (
	addr       = flag.String("addr", ":8080", "")
	configFile = flag.String("config", "config.json", "")
)

func main() {
	flag.Parse()

	conf, err := config.Read(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	_ = conf

	service := web.NewService()
	transport := http.NewTransport(service)

	errCh := transport.Start(*addr)

	log.Println(<-errCh)
}
