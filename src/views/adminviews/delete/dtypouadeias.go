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
	views.GetMux().HandleFunc("/dtyposadeias", middleware.WithMiddleware(dtypouadeias,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func dtypouadeias(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	stmt := datastorage.GetDataRouter().GetStmt("delete_typos_adeias")
	_, err := stmt.Exec(id)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την διαγραφή του τύπου άδειας")
		log.Println(err)
		http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
		return
	}
	messages.SetMessage(r, "Ο τύπος άδειας διαγράφηκε επιτυχώς")
	http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
}
