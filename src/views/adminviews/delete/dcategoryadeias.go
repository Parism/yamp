package delete

import (
	"datastorage"
	"log"
	"messages"
	"middleware"
	"net/http"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/dcategoryadeias", middleware.WithMiddleware(dcategoryadeias,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func dcategoryadeias(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	stmt := datastorage.GetDataRouter().GetStmt("delete_category_adeias")
	_, err := stmt.Exec(id)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την διαγραφή της κατηγορίας μεταβολής")
		log.Println(err)
		http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Η κατηγορία μεταβολής διαγράφηκε επιτυχώς")
	http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
}
