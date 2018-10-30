package middleware

import (
	"auth"
	"net/http"
	"utils"
	"variables"
)

/*
IsYditis middleware function
passes the cookie value to the according gatekeeper function
if the cookie is authenticated and the role of the user is yditis
then proceed to the next middleware
Else, provide a message to the session
and redirect to login page
*/
func IsYditis() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("sessionid")
			if err != nil { //cookie does not exist, redirect to /login
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			}
			if auth.GetGatekeeper().CheckRoleAndAuth(cookie.Value) >= variables.YDITIS {
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
				http.Redirect(w, r, utils.RedirectByRole(r), http.StatusMovedPermanently)
			}
		})
	}
}
