package web

import "x-bank-users/auth"

type (
	UserDataToSignIn struct {
		Id              int64
		PasswordHash    []byte
		TelegramId      *int64
		IsActivated     bool
		HasPersonalData bool
	}

	SignInResult struct {
		AccessClaims auth.Claims
		RefreshToken string
	}
)
