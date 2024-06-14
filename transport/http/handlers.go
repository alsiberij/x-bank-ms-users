package http

import (
	"net/http"
)

func (t *Transport) handlerSignUp(w http.ResponseWriter, r *http.Request) {
	// TODO Алёна
	t.service.SignUp()
}

func (t *Transport) handlerActivateAccount(w http.ResponseWriter, r *http.Request) {
	// TODO Игорь
	t.service.ActivateAccount()
}
