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
	"x-bank-users/transport/http/jwt"
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

	knife := swissknife.NewService()
	passwordHasher := hasher.NewService()

	jwtHs512, err := jwt.NewHS512(conf.Hs512SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	randomGenerator := random.NewService()

	service := web.NewService(&knife, &randomGenerator, &knife, &knife, &passwordHasher, &knife, &knife, &knife, &knife)

	transport := http.NewTransport(service, &jwtHs512)

	// TODO Алёна.
	// Сделать graceful shutdown. Ловим SIGINT и SIGTERM, используем метод Stop у транспорта, таймаут на остановку - 30 сек.
	errCh := transport.Start(*addr)

	log.Println(<-errCh)
}
