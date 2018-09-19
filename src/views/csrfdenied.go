package views

import (
	"fmt"
	"net/http"
)

func csrfdenied(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Csrf token invalid")
}
