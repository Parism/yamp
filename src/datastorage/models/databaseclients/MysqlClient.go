package databaseclients

import "database/sql"

/*
MysqlClient struct
holds the db connection
and the configuration
*/
type MysqlClient struct {
	db *sql.DB
}

/*
GetClient function
returns the pointer to the actual client
*/
func (mc *MysqlClient) GetClient() *sql.DB {
	return mc.db
}

/*
SetClient function
set the actual client
*/
func (mc *MysqlClient) SetClient(db *sql.DB) {
	mc.db = db
}

/*
CheckConnection function
is used to check the client
is also the reference point for the storage interface
*/
func (mc *MysqlClient) CheckConnection() bool {
	err := mc.db.Ping()
	if err != nil {
		return false
	}
	return true

}
