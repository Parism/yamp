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
	views.GetMux().HandleFunc("/ccategoryadeias", middleware.WithMiddleware(ccategoryadeias,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
ccategoryadeias function
is responsible for creating leave categories
*/
func ccategoryadeias(w http.ResponseWriter, r *http.Request) {
	categoryadeias := r.PostFormValue("categoryadeias")
	stmt := datastorage.GetDataRouter().GetStmt("create_category_adeias")
	_, err := stmt.Exec(categoryadeias)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την δημιουργία της κατηγορίας μεταβολής")
		log.Println(err)
	} else {
		messages.SetMessage(r, "H κατηγορία μεταβολής δημιουργήθηκε επιτυχώς")
	}
	http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
}
