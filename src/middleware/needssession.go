package middleware

import (
	"auth"
	"net/http"
)

/*
NeedsSession function
is a middleware function
ensuring that a request will always have a session associated with it
*/
func NeedsSession() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie("sessionid")
			if err != nil { //cookie does not exist, create it
				http.SetCookie(w, auth.NewCookie())
				return
			}
		})
	}
}
