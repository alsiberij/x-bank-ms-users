package http

import "net/http"

type (
	middleware func(http.HandlerFunc) http.HandlerFunc
)
