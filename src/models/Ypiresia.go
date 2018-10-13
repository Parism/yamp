package models

/*
TyposYpiresias struct
holds info about the different types of ypiresia object
*/
type TyposYpiresias struct {
	ID        int
	Perigrafi string
}

/*
Ypiresia struct
holds the actual information
*/
type Ypiresia struct {
	ID        int    `json:"id"`
	PersonID  int    `json:"idperson"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Rank      string `json:"rank"`
	Perigrafi string `json:"perigrafi"`
	Date      string `json:"date"`
}
