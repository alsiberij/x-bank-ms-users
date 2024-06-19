package http

type (
	UserDataToSignUp struct {
		Email    string `json:"email"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	UserDataToSignIn struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		TwoFaDemand string    `json:"2FA"`
		Tokens      TokenPair `json:"tokens"`
	}

	TokenPair struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	UserDataToSignIn2FA struct {
		Code string `json:"code"`
	}

	RecoveryRequest struct {
		Login string `json:"login"`
		Email string `json:"email"`
	}

	RecoveryCodeRequest struct {
		Password string `json:"password"`
	}
)
