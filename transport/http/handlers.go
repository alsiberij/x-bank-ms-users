package http

import (
	"net/http"
)

func (t *Transport) handlerSignUp(w http.ResponseWriter, r *http.Request) {
	// TODO Алёна
	t.service.SignUp()
}

func (t *Transport) handlerActivateAccount(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	err := t.service.ActivateAccount(r.Context(), code)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
