package models

/*
Anafora struct
holds all the required information
for the anafora object
*/
type Anafora struct {
	ID        int    `json:"id"`
	IDPerson  int    `json:"idperson"`
	Perigrafi string `json:"perigrafi"`
	Date      string `json:"date"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
}
