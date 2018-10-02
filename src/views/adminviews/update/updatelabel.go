package update

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"strconv"
	"utils"
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
	label := r.PostFormValue("label")
	labelint, _ := strconv.Atoi(label)
	stmt := datastorage.GetDataRouter().GetStmt("update_user_label")
	_, err := stmt.Exec(labelint, id)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveuser?id="+id, "Σφάλμα κατά τon προσδιορισμό ομάδος", err)
		return
	}
	messages.SetMessage(r, "Προσδιορισμός ομάδος επιτυχής")
	http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
}
