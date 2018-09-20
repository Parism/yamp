package views

import (
	"auth"
	"datastorage"
	"html/template"
	"middleware"
	"net/http"
	"utils"

	"messages"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	GetMux().HandleFunc("/secretadmin", middleware.WithMiddleware(secret,
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
	GetMux().HandleFunc("/secret",
		middleware.WithMiddleware(secret,
			middleware.NeedsSession(),
			middleware.IsUser(),
		))
	GetMux().HandleFunc("/dosignup",
		middleware.WithMiddleware(postsignup,
			middleware.NeedsSession(),
			middleware.CsrfProtection(),
		))
	GetMux().HandleFunc("/dologin",
		middleware.WithMiddleware(postlogin,
			middleware.NeedsSession(),
			middleware.CsrfProtection(),
		))
	GetMux().HandleFunc("/csrfdenied", csrfdenied)
	GetMux().HandleFunc("/",
		middleware.WithMiddleware(index,
			middleware.Time(),
			middleware.NeedsSession(),
		))
}

/*
Context struct will hold all common files among the
data objects
All the data objects will also have a context object
*/
type Context struct {
	Csrftoken string
	Message   string
}

/*
Data struct is a dummy
struct holding data to test context
*/
type Data struct {
	Context Context
	Data    string
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
	csrftoken := utils.GetSessionValue(r, "csrftoken")
	message := messages.GetMessage(r)
	var context Context
	context.Csrftoken = csrftoken
	context.Message = message
	data := Data{}
	data.Context = context
	data.Data = "testdata"
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, data)
}

/*
postsignup function
if user gets succesfully registered
log him in and redirect. to be implemented
*/
func postsignup(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	role := r.PostFormValue("role")
	dbclient, _ := datastorage.GetDataRouter().GetDb("common")
	db := dbclient.GetMysqlClient()
	stmt, _ := db.Prepare("INSERT INTO accounts (username,password,role) VALUES(?,?,?);")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := stmt.Exec(username, hash, role)
	if err != nil {
		messages.SetMessage(r, "Το όνομα υπάρχει ήδη")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	cookie, _ := r.Cookie("sessionid")
	sessionid := cookie.Value
	auth.GetGatekeeper().Login(sessionid, role, w, r)
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
