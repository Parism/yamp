package update

import (
	"database/sql"
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/updateuserlabel", middleware.WithMiddleware(updateUserLabel,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsAdmin(),
	))
}

/*
updateUserLabel function
is used to change the password of an account
if it ever gets lost.
*/
func updateUserLabel(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	var label interface{}
	label = r.PostFormValue("label")
	if label == "nil" {
		label = sql.NullString{}
	}
	stmt := datastorage.GetDataRouter().GetStmt("update_user_label")
	_, err := stmt.Exec(label, id)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά τon προσδιορισμό ομάδος")
		http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Προσδιορισμός ομάδος επιτυχής")
	http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
}
