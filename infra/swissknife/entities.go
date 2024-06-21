package swissknife

import "time"

type (
	storedUser struct {
		Login           string
		Email           string
		Password        []byte
		IsActivated     bool
		HasPersonalData bool
		TelegramId      *int64
		CreatedAt       time.Time
	}
)
