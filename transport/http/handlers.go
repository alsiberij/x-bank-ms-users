package http

import (
	"encoding/json"
	"net/http"
)

func (t *Transport) handlerSignUp(w http.ResponseWriter, r *http.Request) {
	userData := UserDataToSignUp{}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	if err := t.service.SignUp(r.Context(), userData.Login, userData.Password, userData.Email); err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
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

func (t *Transport) handlerSignIn(w http.ResponseWriter, r *http.Request) {
	userDataToSignIn := UserDataToSignIn{}
	if err := json.NewDecoder(r.Body).Decode(&userDataToSignIn); err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	signInResult, err := t.service.SignIn(r.Context(), userDataToSignIn.Login, userDataToSignIn.Password)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	token, err := t.authorizer.Authorize(r.Context(), signInResult.AccessClaims)
	signInResponse := SignInResponse{}

	if signInResult.AccessClaims.Is2FAToken {
		signInResponse.TwoFaDemand = string(token)
		return
	}

	signInResponse.Tokens.AccessToken = string(token)
	signInResponse.Tokens.RefreshToken = signInResult.RefreshToken
}

func (t *Transport) handlerSignIn2FA(w http.ResponseWriter, r *http.Request) {
	// TODO Игорь
	// 1. Парсим тело запроса (структура UserDataToSignIn2FA)
	// 2. Получаем из контекста *auth.Claims (см. свое другое TO DO). Ключ - t.claimsCtxKey
	// 2. Вызываем бизнес логику
	// 3. Формируем токен с помощью t.authorizer.Authorize
	// 4. Формируем ответ (структура TokenPair).
}
