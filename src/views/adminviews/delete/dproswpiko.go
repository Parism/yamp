package delete

import (
	"datastorage"
	"log"
	"messages"
	"middleware"
	"net/http"
	"strconv"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/dproswpiko", middleware.WithMiddleware(dproswpiko,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func dproswpiko(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("delete_proswpiko")
	res, err := stmt.Exec(idint)
	log.Println(id)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+id, "Ανεπιτυχής διαγραφή προσωπικού", err)
		return
	}
	log.Println(res.RowsAffected())
	messages.SetMessage(r, "Επιτυχής διαγραφή προσωπικού")
	http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
}
