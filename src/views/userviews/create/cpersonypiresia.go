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
	views.GetMux().HandleFunc("/cpersonypiresia", middleware.WithMiddleware(cpersonypiresia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func cpersonypiresia(w http.ResponseWriter, r *http.Request) {
	ypiresiaid := r.PostFormValue("ypiresiaid")
	ypiresiaidint, _ := strconv.Atoi(ypiresiaid)
	personid := r.PostFormValue("personid")
	personidint, _ := strconv.Atoi(personid)
	if utils.CanActOnPerson(r, personidint) == false {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Μη εξουσιοδοτημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("create_person_ypiresia")
	_, err := stmt.Exec(personidint, ypiresiaidint)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα κατά την προσθήκη", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής προσθήκη υπηρεσίας")
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
}
