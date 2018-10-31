package databaseclients

import (
	"database/sql"

	"github.com/go-redis/redis"
)

/*
DbClient interface
is used to have multiple types of databases
in the datarouter
e.g. mysql,redis and so on
Each dbclient object must implement all methods BUT!
A mysqlclient will return nil when it is asked for a redis client
and vice-versa
Obviously a hack. Reflection is costly. Performance uber alles
*/
type DbClient interface {
	CheckConnection() bool
	GetRedisClient() *redis.Client
	GetMysqlClient() *sql.DB
	CloseConnection()
}
