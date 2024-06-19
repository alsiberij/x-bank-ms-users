package swissknife

type (
	storedUser struct {
		Login           string
		Email           string
		Password        []byte
		IsActivated     bool
		HasPersonalData bool
		TelegramId      *int64
	}
)
