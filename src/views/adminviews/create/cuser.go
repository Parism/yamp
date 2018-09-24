package create

import (
	"datastorage"
	"messages"
	"middleware"
	"net/http"
	"utils"
	"views"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	views.GetMux().HandleFunc("/cuser", middleware.WithMiddleware(user,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsAdmin(),
	))
}

/*
User function
creates either an admin user
or a simple user
must provide group for simple users
*/
func user(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password1 := r.PostFormValue("password1")
	password2 := r.PostFormValue("password2")
	role := r.PostFormValue("role")
	if password1 != password2 {
		utils.RedirectWithError(w, r, "listusers", "Οι 2 κωδικοί χρηστών δεν ταιριάζουν", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("insert_new_user")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	_, err := stmt.Exec(username, hash, role)
	if err != nil {
		utils.RedirectWithError(w, r, "listusers", "Το όνομα υπάρχει ήδη", err)
		return
	}
	messages.SetMessage(r, "Επιτυχής εισαγωγή χρήστη")
	http.Redirect(w, r, "retrieveuser?username="+username, http.StatusMovedPermanently)
}
