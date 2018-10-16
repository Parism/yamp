package middleware

import (
	"net/http"
)

/*
NoCache function of middleware package
prevents the requesting agent from storing
the response.
Middleware for better cache control can be added.
*/
func NoCache() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "max-age:5, public")
			h.ServeHTTP(w, r)
		})
	}
}
