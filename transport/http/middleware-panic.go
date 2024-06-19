package http

import "net/http"

// TODO Игорь
// Написать мидлвару которая будет ловить панику и отдавать на фронт ошибку с помощью setFatalError.
// Написать проверку что panicMiddleware приводится к типу middleware

func (t *Transport) panicMiddleware(h http.HandlerFunc) http.HandlerFunc {
	panic("implement me")
}
