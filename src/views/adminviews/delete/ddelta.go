package delete

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
	views.GetMux().HandleFunc("/ddelta", middleware.WithMiddleware(ddelta,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func ddelta(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("delete_delta")
	_, err := stmt.Exec(idint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrievedelta?id="+id, "Υπάρχουν ακόμα άτομα συνδεδεμένα με το Δέλτα", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή Δέλτα")
	http.Redirect(w, r, "/deltas", http.StatusMovedPermanently)
}
