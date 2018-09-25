package delete

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/ddelta", middleware.WithMiddleware(ddelta,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func ddelta(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	stmt := datastorage.GetDataRouter().GetStmt("delete_delta")
	_, err := stmt.Exec(id)
	if err != nil {
		utils.RedirectWithError(w, r, "/deltas", "Ανεπιτυχής διαγραφή delta", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή delta")
	http.Redirect(w, r, "/deltas", http.StatusMovedPermanently)
}
