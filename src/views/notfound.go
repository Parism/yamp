package views

import (
	"fmt"
	"middleware"
	"net/http"
)

func init() {
	GetMux().HandleFunc("/notfound", middleware.WithMiddleware(notfound,
		middleware.Time(),
	))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Object not found, http error 404")
}
