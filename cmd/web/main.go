package main

import (
	"flag"
	"log"
	"practice/config"
	"practice/core/web"
	"practice/transport/http"
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
