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
	views.GetMux().HandleFunc("/dpersonypiresia", middleware.WithMiddleware(dpersonypiresia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func dpersonypiresia(w http.ResponseWriter, r *http.Request) {
	idypiresias := r.PostFormValue("id")
	idypiresiasint, _ := strconv.Atoi(idypiresias)
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	if utils.CanActOnPerson(r, personidint) == false {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Μη εξουσιοδοτημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("delete_person_ypiresia")
	res, err := stmt.Exec(idypiresiasint)
	affected, _ := res.RowsAffected()
	if (err != nil) || affected < 1 {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα κατά την διαγραφή", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή υπηρεσίας")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)

}
