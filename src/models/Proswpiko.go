package models

/*
Proswpiko struct
holds the representation of personel
*/
type Proswpiko struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Rank    string `json:"rank"`
	Label   string `json:"label"`
}
