package adminviews

import (
	"bytes"
	"datastorage"
	"encoding/json"
	"fmt"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/getaitiseis", middleware.WithMiddleware(getaitiseis,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func getaitiseis(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	lastinpage := r.URL.Query().Get("maxid")
	var buffer bytes.Buffer
	buffer.WriteString("SELECT aitiseis.id, idperson, aitiseis.perigrafi, date,")
	buffer.WriteString("name, surname, ranks.rank, ierarxia.perigrafi FROM aitiseis ")
	buffer.WriteString("JOIN proswpiko on idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia on proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE aitiseis.id > ? ")
	buffer.WriteString("ORDER BY aitiseis.id ASC, proswpiko.rank DESC, date ASC LIMIT 4;")
	res, err := dbc.Query(buffer.String(), lastinpage)
	defer res.Close()
	if err != nil {
		utils.RedirectWithError(w, r, utils.RedirectByRole(r), "Σφάλμα ανάκτησης αιτήσεων", err)
		return
	}
	var aitisi models.Aitisi
	var aitiseis []models.Aitisi
	for res.Next() {
		_ = res.Scan(
			&aitisi.ID,
			&aitisi.IDPerson,
			&aitisi.Perigrafi,
			&aitisi.Date,
			&aitisi.Name,
			&aitisi.Surname,
			&aitisi.Rank,
			&aitisi.Monada,
		)
		aitisi.Date = models.DateBuilder(aitisi.Date)
		aitiseis = append(aitiseis, aitisi)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonString, err := json.MarshalIndent(aitiseis, "", " ")
	fmt.Fprintf(w, string(jsonString[:]))
}
