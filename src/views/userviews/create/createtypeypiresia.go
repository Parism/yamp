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
	views.GetMux().HandleFunc("/ctypeypiresias", middleware.WithMiddleware(ctypeypiresias,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func ctypeypiresias(w http.ResponseWriter, r *http.Request) {
	perigrafi := r.PostFormValue("perigrafi")
	label := r.PostFormValue("label")
	labelform, _ := strconv.Atoi(label)
	if utils.CheckLabelAuthed(r, labelform) == false {
		utils.RedirectWithError(w, r, "/ypiresies", "Μη εξουσιοδοτημένο αίτημα", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("create_typos_ypiresias")
	_, err := stmt.Exec(perigrafi, labelform)
	if err != nil {
		utils.RedirectWithError(w, r, "/ypiresies", "Σφάλμα κατά την καταχώρηση", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής δημιουργία τύπου υπηρεσίας")
	http.Redirect(w, r, "/ypiresies", http.StatusMovedPermanently)

}
