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
		Id              int64
		PhoneNumber     string
		FirstName       string
		LastName        string
		FathersName     *string
		DateOfBirth     time.Time
		PassportId      string
		Address         string
		Gender          string
		LiveInCountry   string
		UserEmployments []UserEmployment
	}

	UserEmployment struct {
		UserId      int64
		WorkplaceId int64
		Position    string
		StartDate   time.Time
		EndDate     time.Time
	}

	Workplace struct {
		Id      int64
		Name    string
		Address string
	}
)
