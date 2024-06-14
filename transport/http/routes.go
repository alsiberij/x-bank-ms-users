package http

import (
	"net/http"
)

func (t *Transport) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/sign-up", t.handlerSignUp)
	mux.HandleFunc("POST /v1/auth/verification", t.handlerActivateAccount)

	return mux
}
