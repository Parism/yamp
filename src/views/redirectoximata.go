package views

import (
	"fmt"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/redirectoximata", middleware.WithMiddleware(redirectoximata,
		middleware.Time(),
		middleware.NeedsSession(),
	))
}

func redirectoximata(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not yet implemented")
}
