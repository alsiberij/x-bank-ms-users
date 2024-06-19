package http

import "net/http"

// TODO Алёна
// Написать метод принимающий corsHandler, который будет отдавать мидлвару использующую этот хендлер.

func (t *Transport) corsMiddleware(cors http.HandlerFunc) middleware {
	panic("implement me")
}
