package adminviews

import (
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/epopteia", middleware.WithMiddleware(epopteia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.NoCache(),
		middleware.IsAdmin(),
	))
}

func epopteia(w http.ResponseWriter, r *http.Request) {

}
