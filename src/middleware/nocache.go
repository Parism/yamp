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
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
			h.ServeHTTP(w, r)
		})
	}
}
