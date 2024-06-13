package http

import (
	"net/http"
	"x-bank-users/cerrors"
	"x-bank-users/core/web"
)

type (
	Transport struct {
		service      web.Service
		errorHandler errorHandler
	}
)

func NewTransport(service web.Service) Transport {
	t := Transport{
		service: service,
		errorHandler: errorHandler{
			defaultStatusCode: http.StatusBadRequest,
			statusCodes:       map[cerrors.Code]int{},
		},
	}

	return t
}

func (t *Transport) Start(addr string) chan error {
	srv := &http.Server{Addr: addr, Handler: t.routes()}
	ch := make(chan error)

	go func() {
		ch <- srv.ListenAndServe()
	}()

	return ch
}
