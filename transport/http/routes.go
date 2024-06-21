package http

import (
	"net/http"
)

func (t *Transport) routes() http.Handler {
	corsMiddleware := t.corsMiddleware(t.corsHandler("*", "*", "*", ""))

	defaultGroup := middlewareGroup{
		t.panicMiddleware,
		corsMiddleware,
	}

	signIn2FaMiddleware := middlewareGroup{
		t.panicMiddleware,
		corsMiddleware,
		t.authMiddleware(true),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/sign-up", defaultGroup.Apply(t.handlerSignUp))
	mux.HandleFunc("POST /v1/auth/verification", defaultGroup.Apply(t.handlerActivateAccount))
	mux.HandleFunc("POST /v1/auth/sign-in", defaultGroup.Apply(t.handlerSignIn))
	mux.HandleFunc("POST /v1/auth/sign-in/2fa", signIn2FaMiddleware.Apply(t.handlerSignIn2FA))
	mux.HandleFunc("POST /v1/auth/recovery", defaultGroup.Apply(t.handlerRecovery))
	mux.HandleFunc("POST /v1/auth/recovery/{code}", defaultGroup.Apply(t.handlerRecoveryCode))
	mux.HandleFunc("POST /v1/auth/refresh", defaultGroup.Apply(t.handlerRefresh))

	return mux
}
