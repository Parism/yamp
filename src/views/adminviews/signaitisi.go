package adminviews

import (
	"datastorage"
	"fmt"
	"log"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/signaitisi", middleware.WithMiddleware(signaitisi,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsCaptain(), //user must be at least captain to sign an aitisi object
	))
}

func signaitisi(w http.ResponseWriter, r *http.Request) {
	idaitisi := r.PostFormValue("id")
	sign := r.PostFormValue("sign")
	user := utils.GetSessionValue(r, "username")
	signedAs := utils.GetSessionValue(r, "role")
	stmt := datastorage.GetDataRouter().GetStmt("sign_aitisi")
	res, err := stmt.Exec(user, idaitisi, sign, signedAs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	aff, _ := res.RowsAffected()
	if aff < 1 {
		log.Println("Bad request", aff)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	fmt.Fprintf(w, "OK")
}
