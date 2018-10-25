package views

import (
	"middleware"
	"net/http"
	"utils"
)

func init() {
	GetMux().HandleFunc("/redirectproswpiko", middleware.WithMiddleware(redirectproswpiko,
		middleware.Time(),
		middleware.NeedsSession(),
	))
}

func redirectproswpiko(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, utils.RedirectByRole(r), http.StatusFound)
}
