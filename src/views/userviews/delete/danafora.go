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
	views.GetMux().HandleFunc("/danafora", middleware.WithMiddleware(danafora,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func danafora(w http.ResponseWriter, r *http.Request) {
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	anaforaid := r.PostFormValue("id")
	anaforaidint, _ := strconv.Atoi(anaforaid)
	if !utils.CanActOnPerson(r, personidint) {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Μη αυθεντικοποιημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("delete_anafora")
	_, err := stmt.Exec(anaforaidint, personidint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα κατά την διαγραφή της αίτησης", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή αίτησης")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
}
