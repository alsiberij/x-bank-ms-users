package swissknife

type (
	storedUser struct {
		Login       string
		Email       string
		Password    []byte
		IsActivated bool
	}
)
