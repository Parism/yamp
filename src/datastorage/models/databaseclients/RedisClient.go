package databaseclients

import (
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
func (rc *RedisClient) GetClient() *redis.Client {
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
