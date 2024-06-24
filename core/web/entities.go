package web

import (
	"time"
	"x-bank-users/auth"
)

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

	UserPersonalData struct {
		Id            int64
		PhoneNumber   string
		FirstName     string
		LastName      string
		FathersName   *string
		DateOfBirth   time.Time
		PassportId    string
		Address       string
		Gender        string
		LiveInCountry string
		// TODO Добавить список usersEmployments
	}

	UserData struct {
		Id           int64
		UUID         string
		Login        string
		Email        string
		PasswordHash []byte
		TelegramId   *int64
		IsActivated  bool
		CreatedAt    time.Time
	}
)
