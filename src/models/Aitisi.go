package models

/*
Aitisi struct
holds information about aitiseis
*/
type Aitisi struct {
	ID        int    `json:"id"`
	Perigrafi string `json:"perigrafi"`
	IDPerson  int    `json:"idperson"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
}
