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
	views.GetMux().HandleFunc("/dlambda", middleware.WithMiddleware(dlambda,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func dlambda(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	stmt := datastorage.GetDataRouter().GetStmt("delete_lambda")
	_, err := stmt.Exec(id)
	if err != nil {
		utils.RedirectWithError(w, r, "/lambdas", "Ανεπιτυχής διαγραφή lambda", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή lambda")
	http.Redirect(w, r, "/lambdas", http.StatusMovedPermanently)
}
