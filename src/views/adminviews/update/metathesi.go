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
	views.GetMux().HandleFunc("/metathesi", middleware.WithMiddleware(metathesi,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func metathesi(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	label := r.PostFormValue("label")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("metathesi")
	_, err := stmt.Exec(label, idint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+id, "Σφάλμα κατά την διαδικασία μετάθεσης", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής μετάθεση")
	http.Redirect(w, r, "/retrieveproswpiko?id="+id, http.StatusMovedPermanently)
}
