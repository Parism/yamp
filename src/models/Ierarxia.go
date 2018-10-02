package models

/*
Ierarxia struct represents
the chain of command
*/
type Ierarxia struct {
	ID              int
	Perigrafi       string
	ParentID        int
	ParentPerigrafi string
}
