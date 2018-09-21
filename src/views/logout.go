package views

import (
	"auth"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/logout",
		middleware.WithMiddleware(Logout,
			middleware.NeedsSession(),
			middleware.CsrfProtection(),
		))
}

/*
Logout function
offloads the logout functionality to the gatekeeper
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.GetGatekeeper().Logout(w, r)
}
