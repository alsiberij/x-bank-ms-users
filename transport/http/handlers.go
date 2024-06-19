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
	// TODO Алёна
	// 1. Парсим тело запроса (структура UserDataToSignIn)
	// 2. Вызываем бизнес логику
	// 3. Формируем токен с помощью t.authorizer.Authorize
	// 4. Формируем ответ (структура SignInResponse).
	// 4.1. Если claims.Is2FAToken == true, то сохраняем токен в поле SignInResponse.TwoFaDemand, остальное пустое.
	// 4.2. Иначе заполняем структуру TokenPair
}

func (t *Transport) handlerSignIn2FA(w http.ResponseWriter, r *http.Request) {
	// TODO Игорь
	// 1. Парсим тело запроса (структура UserDataToSignIn2FA)
	// 2. Получаем из контекста *auth.Claims (см. свое другое TO DO). Ключ - t.claimsCtxKey
	// 2. Вызываем бизнес логику
	// 3. Формируем токен с помощью t.authorizer.Authorize
	// 4. Формируем ответ (структура TokenPair).
}

func (t *Transport) handlerRecovery(w http.ResponseWriter, r *http.Request) {
	var request RecoveryRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	err = t.service.Recovery(r.Context(), request.Login, request.Email)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *Transport) handlerRecoveryCode(w http.ResponseWriter, r *http.Request) {
	var request RecoveryCodeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	code := r.PathValue("code")

	err = t.service.RecoveryCode(r.Context(), code, request.Password)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
