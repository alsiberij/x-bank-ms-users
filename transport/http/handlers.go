package http

import (
	"fmt"
	"net/http"
)

func (t *Transport) handlerSignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIGN UP")
	// TODO
}

func (t *Transport) handlerActivateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("VERIFICATION")
	// TODO
}
