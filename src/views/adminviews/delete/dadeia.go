package delete

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
	views.GetMux().HandleFunc("/dadeia", middleware.WithMiddleware(dadeia,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func dadeia(w http.ResponseWriter, r *http.Request) {
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	if !utils.CanActOnPerson(r, personidint) {
		http.Redirect(w, r, "/notfound", http.StatusMovedPermanently)
		return
	}
	id := r.PostFormValue("id")
	stmt := datastorage.GetDataRouter().GetStmt("delete_adeia")
	_, err := stmt.Exec(id)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την διαγραφή της άδειας")
		log.Println(err)
		http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή άδειας")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)

}
