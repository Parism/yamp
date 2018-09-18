package middleware

import (
	//"auth"
	"auth"
	"net/http"
)

/*
WithLogin middleware
checks if a request is authenticated
will use additional middleware to check role
*/
func WithLogin() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("sessionid")
			if err != nil { //cookie does not exist, redirect to /login
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
				return
			}
			if auth.GetGatekeeper().Checkauth(cookie.Value) {
				h.ServeHTTP(w, r)
				/*
					cookie exists and is authenticated
					proceed to next middleware
				*/
			} else {
				/*
					cookie exists but is not authenticated
					redirect to /login
				*/
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			}
		})
	}
}
