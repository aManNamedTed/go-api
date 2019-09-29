package app

import (
	"net/http"
	u "go-api/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resource was not found."))
		next.ServeHTTP(w, r)
	})
}