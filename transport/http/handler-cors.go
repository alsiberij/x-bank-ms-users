package http

import "net/http"

// TODO Алёна
// Написать метод, который будет принимать значения CORS заголовков и отдавать http.HandlerFunc, в котором
// будут проставляться эти заголовки. Разобраться что это и для чего нужно.

func (t *Transport) corsHandler(allowOrigin, allowHeaders, allowMethods, exposeHeaders string) http.HandlerFunc {
	panic("implement me")
}
