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
	views.GetMux().HandleFunc("/setuserdb", middleware.WithMiddleware(setuserdb,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func setuserdb(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	db := r.PostFormValue("db")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("delete_user_db")
	_, err := stmt.Exec(idint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveuser?id="+id, "error deleting user db access", err)
		return
	}

	stmt = datastorage.GetDataRouter().GetStmt("create_db_access")
	_, err = stmt.Exec(id, db)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveuser?id="+id, "error creating user db access", err)
		return
	}
	messages.SetMessage(r, "Απόδωση δικαιωμάτων επιτυχής")
	http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
}
