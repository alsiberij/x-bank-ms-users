package http

type (
	UserData struct {
		Email    string `json:"email"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)
