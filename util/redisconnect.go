package util

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// func RedisConnect is return connector for Redis.
func RedisConnect() redis.Conn {
	// Get the connection object
	conn, err := redis.Dial("tcp", "localhost:6380")
	if err != nil {
		log.Fatal("Error connect to Redis")
	}

	return conn
}
