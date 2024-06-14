package http

import (
	"encoding/json"
	"net/http"
)

func (t *Transport) handlerSignUp(w http.ResponseWriter, r *http.Request) {
	userData := struct {
		Email    string
		Login    string
		Password string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	if err := t.service.SignUp(r.Context(), userData.Login, userData.Password, userData.Login); err != nil {
		t.errorHandler.setError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *Transport) handlerActivateAccount(w http.ResponseWriter, r *http.Request) {
	// TODO Игорь
	t.service.ActivateAccount()
}
