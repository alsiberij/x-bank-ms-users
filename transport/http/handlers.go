package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"x-bank-users/auth"
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
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}
	signInResponse := SignInResponse{}

	if signInResult.AccessClaims.Is2FAToken {
		signInResponse.TwoFaDemand = string(token)
	} else {
		signInResponse.Tokens.AccessToken = string(token)
		signInResponse.Tokens.RefreshToken = signInResult.RefreshToken
	}

	err = json.NewEncoder(w).Encode(signInResponse)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *Transport) handlerSignIn2FA(w http.ResponseWriter, r *http.Request) {
	userDataToSignIn2FA := UserDataToSignIn2FA{}

	err := json.NewDecoder(r.Body).Decode(&userDataToSignIn2FA)
	if err != nil {
		t.errorHandler.setBadRequestError(w, err)
		return
	}

	claims, ok := r.Context().Value(t.claimsCtxKey).(*auth.Claims)
	if !ok {
		t.errorHandler.setError(w, errors.New("отсутствуют claims в контексте"))
		return
	}

	code := userDataToSignIn2FA.Code

	signInResult, err := t.service.SignIn2FA(r.Context(), *claims, code)

	token, err := t.authorizer.Authorize(r.Context(), signInResult.AccessClaims)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	signInResponse := SignInResponse{}

	signInResponse.Tokens.AccessToken = string(token)
	signInResponse.Tokens.RefreshToken = signInResult.RefreshToken

	err = json.NewEncoder(w).Encode(signInResponse)
	if err != nil {
		t.errorHandler.setError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
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
