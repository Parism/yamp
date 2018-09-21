package models

/*
User struct
is used to represent the user accounts
not to be confused with persons or session field user
*/
type User struct {
	Username string
	Role     string
	Db       string
}
