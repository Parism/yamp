package databaseclients

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis"
)

/*
RedisClient struct holds the client
communicating with the redis server
that holds the session data
*/
type RedisClient struct {
	client *redis.Client
}

/*
GetClient function
returns the pointer to the actual client
*/
func (rc *RedisClient) GetClient() interface{} {
	return rc.client
}

/*
GetMysqlClient function
returns nil since the receiver of the function is a redis client
*/
func (rc *RedisClient) GetMysqlClient() *sql.DB {
	return nil
}

/*
GetRedisClient function
returns the reference to the actual redis client object
*/
func (rc *RedisClient) GetRedisClient() *redis.Client {
	return rc.client
}

/*
SetClient function
set the actual client
*/
func (rc *RedisClient) SetClient(client *redis.Client) {
	rc.client = client
}

/*
CheckConnection function
is used to check the client
is also the reference point for the storage interface
*/
func (rc *RedisClient) CheckConnection() bool {
	_, err := rc.client.Ping().Result()
	if err != nil {
		log.Println("Check redis connection", err)
		return false
	}
	return true
}

/*
CloseConnection function
closes a redis connection
*/
func (rc *RedisClient) CloseConnection() {
	rc.client.Close()
}
