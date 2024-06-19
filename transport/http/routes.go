package http

import (
	"net/http"
)

func (t *Transport) routes() http.Handler {
	// TODO Алёна, Игорь. Пример использования мидлвары. Потом сделаем красивее. Плюс побегайте дебагером и разберитесь что кого возвращает и вызывает.
	// mux.HandleFunc("POST /v1/test", t.corsMiddleware(t.corsHandler("", "", "", ""))(t.handlerActivateAccount))

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/sign-up", t.handlerSignUp)
	mux.HandleFunc("POST /v1/auth/verification", t.handlerActivateAccount)
	mux.HandleFunc("POST /v1/auth/sign-in", t.handlerSignIn)
	mux.HandleFunc("POST /v1/auth/sign-in/2fa", t.handlerSignIn2FA) // TODO Игорь. Тут обязательно нужна authMiddleware пропускающая втч и 2FA токены

	return mux
}
