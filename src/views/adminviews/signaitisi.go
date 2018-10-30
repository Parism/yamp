package adminviews

import (
	"fmt"
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/signaitisi", middleware.WithMiddleware(signaitisi,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsCaptain(), //use must be at least captain to sign an aitisi object
	))
}

func signaitisi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
