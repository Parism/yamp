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
	views.GetMux().HandleFunc("/cergasia", middleware.WithMiddleware(cergasia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func cergasia(w http.ResponseWriter, r *http.Request) {
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	perigrafi := r.PostFormValue("perigrafi")
	if !utils.CanActOnPerson(r, personidint) {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Μη αυθεντικοποιημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("create_ergasia")
	_, err := stmt.Exec(perigrafi, personidint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα καταχώρησης εργασίας", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής καταχώρηση εργασίας")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
}
