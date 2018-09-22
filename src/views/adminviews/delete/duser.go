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
	username := r.PostFormValue("username")
	stmt := datastorage.GetDataRouter().GetStmt("delete_user")
	_, err := stmt.Exec(username)
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την διαγραφή του χρήστη")
		log.Println(err)
		http.Redirect(w, r, "/retrieveuser?username="+username, http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, "/users", http.StatusMovedPermanently)
}
