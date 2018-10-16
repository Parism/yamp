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
	ID       int64
	Username string
	Role     string
	Label    sql.NullString
}
