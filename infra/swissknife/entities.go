package swissknife

import (
	"time"
	"x-bank-users/core/web"
)

type (
	storedUser struct {
		Login           string
		Email           string
		Password        []byte
		IsActivated     bool
		HasPersonalData *web.UserPersonalData
		TelegramId      *int64
		CreatedAt       time.Time
	}
)
