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
	views.GetMux().HandleFunc("/dtypeypiresias", middleware.WithMiddleware(dtypeypiresias,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func dtypeypiresias(w http.ResponseWriter, r *http.Request) {
	idtypeypiresias := r.PostFormValue("id")
	idint, _ := strconv.Atoi(idtypeypiresias)
	label := r.PostFormValue("label")
	labelform, _ := strconv.Atoi(label)
	if utils.CheckLabelAuthed(r, labelform) == false {
		utils.RedirectWithError(w, r, "/ypiresies", "Μη εξουσιοδοτημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("delete_typos_ypiresias")
	res, err := stmt.Exec(idint, labelform)
	affected, _ := res.RowsAffected()
	if (err != nil) || affected < 1 {
		utils.RedirectWithError(w, r, "/ypiresies", "Σφάλμα κατά την διαγραφή", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής διαγραφή τύπου υπηρεσίας")
	http.Redirect(w, r, "/ypiresies", http.StatusMovedPermanently)

}
