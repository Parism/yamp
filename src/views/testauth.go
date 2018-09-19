package views

import (
	"datastorage"
	"html/template"
	"log"
	"middleware"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	GetMux().HandleFunc("/secretadmin", middleware.WithMiddleware(secret,
		middleware.IsAdmin(),
	))
	GetMux().HandleFunc("/secret",
		middleware.WithMiddleware(secret,
			middleware.IsUser(),
		))
	GetMux().HandleFunc("/", index)
}

/*
index function
this function
secret function
and secret with role function
will be used to test the gatekeeper
and the role middleware
*/
func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, nil)
}

/*
login function
passes the credentials provided to the gatekeeper for validation
must implement a login function in the gatekeeper
*/
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postlogin(w, r)
	} else {
		t := template.Must(template.ParseFiles("./templates/login.html"))
		t.Execute(w, nil)
	}
}

/*
signup function
registers test users
*/
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postsignup(w, r)
	} else {
		t := template.Must(template.ParseFiles("./templates/signup.html"))
		t.Execute(w, nil)
	}
}

/*
postsignup function
if user gets succesfully registered
log him in and redirect. to be implemented
*/
func postsignup(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	dbclient, _ := datastorage.GetDataRouter().GetDb("common")
	db := dbclient.GetMysqlClient()
	stmt, _ := db.Prepare("INSERT INTO accounts VALUES(?,?);")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := stmt.Exec(username, hash)
	if err != nil {
		log.Println("post signup error", err)
		return
	}
}

/*
postlogin function
asks gatekeeper if the credentials provided are ok
if they are redirects to secret
if not must return message to session and redirect to login page
either way a session token must be generated
and then check whether the token is authenticated or not
*/
func postlogin(w http.ResponseWriter, r *http.Request) {
	//offload the login function to the gatekeeper
}

func secret(w http.ResponseWriter, r *http.Request) {
	t := template.New("secret")
	t, _ = t.ParseFiles("./templates/secret.html")
	t.Execute(w, nil)
}

func secretadmin(w http.ResponseWriter, r *http.Request) {
	t := template.New("secretadmin")
	t, _ = t.ParseFiles("./templates/secretadmin.html")
	t.Execute(w, nil)
}
