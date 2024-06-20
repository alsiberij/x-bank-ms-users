package http

import "net/http"

func (t *Transport) corsHandler(allowOrigin, allowHeaders, allowMethods, exposeHeaders string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
	}
}
