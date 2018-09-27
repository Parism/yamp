package update

import (
	"datastorage"
	"log"
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
	lambda := r.PostFormValue("lambda")
	delta := r.PostFormValue("delta")
	idint, _ := strconv.Atoi(id)
	lambdaint, _ := strconv.Atoi(lambda)
	deltaint, _ := strconv.Atoi(delta)
	log.Println(idint, lambdaint, deltaint)
	stmt := datastorage.GetDataRouter().GetStmt("metathesi")
	_, err := stmt.Exec(lambdaint, deltaint, idint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+id, "Σφάλμα κατά την διαδικασία μετάθεσης", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής μετάθεση")
	http.Redirect(w, r, "/retrieveproswpiko?id="+id, http.StatusMovedPermanently)
}
