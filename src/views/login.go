package views

import (
	"auth"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/login", middleware.WithMiddleware(login,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
	))
}

/*
login function
sole purpose of this function
is to off load the writer and the reader to the
gatekeeper
*/
func login(w http.ResponseWriter, r *http.Request) {
	auth.GetGatekeeper().Login(w, r)
}
