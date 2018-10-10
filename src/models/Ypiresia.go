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
	ID        int
	Perigrafi string
	Date      string
}
