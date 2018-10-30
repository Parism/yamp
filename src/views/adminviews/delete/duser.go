package delete

import (
	"datastorage"
	"log"
	"messages"
	"middleware"
	"net/http"
	"strconv"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/deleteuser", middleware.WithMiddleware(Duser,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
Duser function
deletes a user account
*/
func Duser(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	idint, _ := strconv.Atoi(id)
	stmt := datastorage.GetDataRouter().GetStmt("delete_user")
	_, err := stmt.Exec(idint)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την διαγραφή του χρήστη")
		log.Println(err)
		http.Redirect(w, r, "/retrieveuser?id="+id, http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, "/listusers", http.StatusMovedPermanently)
}
