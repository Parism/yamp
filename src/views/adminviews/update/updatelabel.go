package update

package update

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"views"

	"golang.org/x/crypto/bcrypt"
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
	stmt := datastorage.GetDataRouter().GetStmt("update_user_label")
	_, err := stmt.Exec(label, id)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά τon προσδιορισμό ομάδος")
		http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Προσδιορισμός ομάδος, ")
	http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
}
