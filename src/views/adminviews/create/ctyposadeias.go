package create

import (
	"datastorage"
	"log"
	"messages"
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/ctyposadeias", middleware.WithMiddleware(ctyposadeias,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func ctyposadeias(w http.ResponseWriter, r *http.Request) {
	typosadeias := r.PostFormValue("typosadeias")
	stmt := datastorage.GetDataRouter().GetStmt("create_typos_adeias")
	_, err := stmt.Exec(typosadeias)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την δημιουργία του τύπου άδειας")
		log.Println(err)
	} else {
		messages.SetMessage(r, "Ο τύπος άδειας δημιουργήθηκε επιτυχώς")
	}
	http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
}
