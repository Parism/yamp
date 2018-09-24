package adminviews

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
	views.GetMux().HandleFunc("/setcaptaindb", middleware.WithMiddleware(setcaptaindb,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func setcaptaindb(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idint, _ := strconv.Atoi(id)
	r.ParseForm()
	selectedvalues := r.Form["db"]
	stmt := datastorage.GetDataRouter().GetStmt("create_db_access")
	for _, value := range selectedvalues {
		_, err := stmt.Exec(value, idint)
		if err != nil {
			utils.RedirectWithError(w, r, "/retrieveuser?id="+id, "Error writing db access", err)
			return
		}
	}
	messages.SetMessage(r, "Απόδωση δικαιωμάτων επιτυχής")
	http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
}
