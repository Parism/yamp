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
	views.GetMux().HandleFunc("/dlambda", middleware.WithMiddleware(dlambda,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func dlambda(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("delete_lambda")
	_, err := stmt.Exec(idint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrievelambda?id="+id, "Υπάρχουν ακόμα άτομα συνδεδεμένα με το Λάμδα", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή Λάμδα")
	http.Redirect(w, r, "/lambdas", http.StatusMovedPermanently)
}
