package views

import (
	"fmt"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/secret",
		middleware.WithMiddleware(secret,
			middleware.WithLogin(),
		))
	GetMux().HandleFunc("/", index)
}

/*
index function
this function
secret function
and secret with role function
will be used to test the gatekeeper
and the role middleware
*/
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret view ok")
}
