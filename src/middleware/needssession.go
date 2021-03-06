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
			cookie, err := r.Cookie("sessionid")
			if err != nil {
				cookie := auth.NewCookie() //cookie does not exist, create it
				r.AddCookie(cookie)
				http.SetCookie(w, cookie)
				http.Redirect(w, r, r.URL.String(), http.StatusFound)
			} else if !auth.GetGatekeeper().SessionExists(cookie.Value) {
				cookie := auth.NewCookie()
				r.AddCookie(cookie)
				http.SetCookie(w, cookie)
				http.Redirect(w, r, r.URL.String(), http.StatusFound)
			}
			//check if cookie exists in db. if not, create a new one
			h.ServeHTTP(w, r)
		})
	}
}
