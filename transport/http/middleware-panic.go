package http

import "net/http"

var (
	_ middleware = (&Transport{}).panicMiddleware
)

func (t *Transport) panicMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				t.errorHandler.setFatalError(w, err)
			}
		}()
		h(w, r)
	}
}
