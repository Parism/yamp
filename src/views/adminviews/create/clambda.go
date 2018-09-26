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
	views.GetMux().HandleFunc("/clambda", middleware.WithMiddleware(clambda,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsAdmin(),
	))
}

/*
clambda function
is responsible for creating objects of type lambda
*/
func clambda(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	stmt := datastorage.GetDataRouter().GetStmt("create_lambda")
	_, err := stmt.Exec(name)
	if err != nil {
		utils.RedirectWithError(w, r, "/listlambdas", "Δημιουργία lambda ανεπιτυχής", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής δημιουργία lambda")
	http.Redirect(w, r, "/lambdas", http.StatusMovedPermanently)
}
