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
	views.GetMux().HandleFunc("/cproswpiko", middleware.WithMiddleware(cproswpiko,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
cproswpiko function
is responsible for registering personel in the application
*/
func cproswpiko(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	surname := r.PostFormValue("surname")
	rank := r.PostFormValue("rank")
	lambda := r.PostFormValue("lambda")
	delta := r.PostFormValue("delta")
	stmt := datastorage.GetDataRouter().GetStmt("create_proswpiko")
	_, err := stmt.Exec(
		name,
		surname,
		rank,
		lambda,
		delta,
	)
	if err != nil {
		utils.RedirectWithError(w, r, "/proswpiko", "Ανεπιτυχής δημιουργία προσωπικού", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής δημιουργία προσωπικού")
	http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
}
