package http

import (
	"net/http"
	"regexp"
)

var (
	isValidEmail = regexp.MustCompile("^.+@.+\\..+$").MatchString
	isValidLogin = regexp.MustCompile("^[a-z0-9_-]{6,32}$").MatchString
)

func (t *Transport) validate(w http.ResponseWriter, v validatable) bool {
	ve := v.validate()
	if len(ve) > 0 {
		t.errorHandler.setUnprocessableEntityError(w, ve)
		return false
	}

	return true
}

func (u *UserDataToSignUp) validate() (ve validationErrors) {
	ve = make(validationErrors, 0, 3)

	if !isValidEmail(u.Email) {
		ve.Add("Неверный адрес электронной почты")
	}

	if !isValidLogin(u.Login) {
		ve.Add("Неверный логин")
	}

	if len(u.Password) < 6 {
		ve.Add("Слишком короткий пароль")
	} else if len(u.Password) > 16 {
		ve.Add("Слишком длинный пароль")
	}

	return
}
