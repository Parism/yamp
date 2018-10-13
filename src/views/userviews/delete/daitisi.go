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
	views.GetMux().HandleFunc("/daitisi", middleware.WithMiddleware(daitisi,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func daitisi(w http.ResponseWriter, r *http.Request) {
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	adeiaid := r.PostFormValue("id")
	adeiaidint, _ := strconv.Atoi(adeiaid)
	if !utils.CanActOnPerson(r, personidint) {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Μη αυθεντικοποιημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("delete_aitisi")
	_, err := stmt.Exec(adeiaidint, personidint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα κατά την διαγραφή της αίτησης", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή αίτησης")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
}
