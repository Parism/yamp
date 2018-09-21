package create

import (
	"datastorage"
	"log"
	"messages"
	"middleware"
	"net/http"
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
	dbase := r.PostFormValue("db")
	if role == "admin" && dbase != "nil" {
		messages.SetMessage(r, "Οι διαχειριστές δεν ανήκουν σε διμοιρίες")
		http.Redirect(w, r, "/xrhstes", http.StatusMovedPermanently)
		return
	}
	if password1 != password2 {
		messages.SetMessage(r, "Οι 2 κωδικοί χρηστών δεν ταιριάζουν")
		http.Redirect(w, r, "/xrhstes", http.StatusMovedPermanently)
		return
	}
	dbclient, _ := datastorage.GetDataRouter().GetDb("common")
	db := dbclient.GetMysqlClient()
	stmt, _ := db.Prepare("INSERT INTO accounts (username,password,role,db) VALUES(?,?,?,?);")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	_, err := stmt.Exec(username, hash, role, dbase)
	if err != nil {
		log.Println(err)
		messages.SetMessage(r, "Το όνομα υπάρχει ήδη")
		http.Redirect(w, r, "/xrhstes", http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, "/xrhstes", http.StatusMovedPermanently)
}
