package http

import (
	"net/http"
)

func (t *Transport) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/sign-up", t.handlerSignUp)
	mux.HandleFunc("POST /v1/auth/verification", t.handlerActivateAccount)
	mux.HandleFunc("POST /v1/auth/sign-in", t.handlerSignIn)
	mux.HandleFunc("POST /v1/auth/sign-in/2fa", t.handlerSignIn2FA) // TODO Игорь. Тут будет вызов authMiddleware

	return mux
}
