package adminviews

import (
	"fmt"
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/getdyn", middleware.WithMiddleware(getdyn,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func getdyn(w http.ResponseWriter, r *http.Request) {
	label := r.URL.Query().Get("label")
	date := r.URL.Query().Get("date")
	fmt.Fprintf(w, "%s %s", label, date)
}
