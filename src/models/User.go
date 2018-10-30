package models

import (
	"database/sql"
)

/*
User struct
is used to represent the user accounts
not to be confused with persons or session field user
*/
type User struct {
	ID        int64
	Username  string
	RealRole  string
	RealLabel sql.NullString
	TempRole  string
	TempLabel sql.NullString
}
