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

	UserData struct {
		Id            int64
		PhoneNumber   string
		FirstName     string
		LastName      string
		FathersName   string
		DateOfBirth   string
		PassportId    string
		Address       string
		Gender        string
		LiveInCountry string
	}
)
