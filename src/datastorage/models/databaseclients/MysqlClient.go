package databaseclients

import (
	"database/sql"

	"github.com/go-redis/redis"
)

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
func (mc *MysqlClient) GetClient() interface{} {
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
GetMysqlClient function
returns a pointer to a sql.DB object
*/
func (mc *MysqlClient) GetMysqlClient() *sql.DB {
	return mc.db
}

/*
GetRedisClient function
returns nil, since the receiver of the function is a mysqlclient object
*/
func (mc *MysqlClient) GetRedisClient() *redis.Client {
	return nil
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

/*
CloseConnection function
closes a mysqlconnection
*/
func (mc *MysqlClient) CloseConnection() {
	mc.db.Close()
}
