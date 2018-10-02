package create

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
	views.GetMux().HandleFunc("/cierarxia", middleware.WithMiddleware(cierarxia,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
cierarxia function
is responsible for
creating objects in the chain of command
*/
func cierarxia(w http.ResponseWriter, r *http.Request) {
	perigrafi := r.PostFormValue("perigrafi")
	parentid := r.PostFormValue("parentid")
	parentidint, _ := strconv.Atoi(parentid)
	stmt := datastorage.GetDataRouter().GetStmt("create_ierarxia")
	_, err := stmt.Exec(
		perigrafi,
		parentidint,
	)
	if err != nil {
		utils.RedirectWithError(w, r, "/diaxeiristiko", "Ανεπιτυχής δημιουργία ιεραρχίας", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής δημιουργία ιεραρχίας")
	http.Redirect(w, r, "/ierarxia", http.StatusMovedPermanently)
}
