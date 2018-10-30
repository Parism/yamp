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

/*
ctyposadeias function
is responsible for creating leave objects
*/
func ctyposadeias(w http.ResponseWriter, r *http.Request) {
	typosadeias := r.PostFormValue("typosadeias")
	categoryadeias := r.PostFormValue("category")
	stmt := datastorage.GetDataRouter().GetStmt("create_typos_adeias")
	_, err := stmt.Exec(typosadeias, categoryadeias)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την δημιουργία του τύπου μεταβολής")
		log.Println(err)
	} else {
		messages.SetMessage(r, "Ο τύπος μεταβολής δημιουργήθηκε επιτυχώς")
	}
	http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
}
