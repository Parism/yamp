package create

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/cdelta", middleware.WithMiddleware(cdelta,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsAdmin(),
	))
}

/*
cdelta function
is responsible for creating objects of type delta
*/
func cdelta(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	stmt := datastorage.GetDataRouter().GetStmt("create_delta")
	_, err := stmt.Exec(name)
	if err != nil {
		utils.RedirectWithError(w, r, "/listdeltas", "Δημιουργία delta ανεπιτυχής", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής δημιουργία delta")
	http.Redirect(w, r, "/deltas", http.StatusMovedPermanently)
}
