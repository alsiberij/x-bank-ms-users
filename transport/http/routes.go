package http

import (
	"net/http"
)

func (t *Transport) routes() http.Handler {
	corsHandler := t.corsHandler("*", "*", "*", "")
	corsMiddleware := t.corsMiddleware(corsHandler)

	defaultMiddlewareGroup := middlewareGroup{
		t.panicMiddleware,
		corsMiddleware,
	}

	signIn2FaMiddlewareGroup := middlewareGroup{
		t.panicMiddleware,
		corsMiddleware,
		t.authMiddleware(true),
	}

	telegramMiddlewareGroup := middlewareGroup{
		t.panicMiddleware,
		corsMiddleware,
		t.authMiddleware(false),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultMiddlewareGroup.Apply(t.handlerNotFound))
	mux.HandleFunc("OPTIONS /", corsHandler)

	mux.HandleFunc("POST /v1/auth/sign-up", defaultMiddlewareGroup.Apply(t.handlerSignUp))
	mux.HandleFunc("POST /v1/auth/verification", defaultMiddlewareGroup.Apply(t.handlerActivateAccount))
	mux.HandleFunc("POST /v1/auth/sign-in", defaultMiddlewareGroup.Apply(t.handlerSignIn))
	mux.HandleFunc("POST /v1/auth/sign-in/2fa", signIn2FaMiddlewareGroup.Apply(t.handlerSignIn2FA))
	mux.HandleFunc("POST /v1/auth/recovery", defaultMiddlewareGroup.Apply(t.handlerRecovery))
	mux.HandleFunc("POST /v1/auth/recovery/{code}", defaultMiddlewareGroup.Apply(t.handlerRecoveryCode))
	mux.HandleFunc("POST /v1/auth/refresh", defaultMiddlewareGroup.Apply(t.handlerRefresh))
	mux.HandleFunc("POST /v1/auth/telegram", telegramMiddlewareGroup.Apply(t.handlerTelegramBind))
	mux.HandleFunc("DELETE /v1/auth/telegram", telegramMiddlewareGroup.Apply(t.handlerTelegramDelete))

	return mux
}
