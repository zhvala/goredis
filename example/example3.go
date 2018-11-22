package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zhvala/goredis"
)

func main() {
	conn, err := redis.NewConn(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
	}

	_, err = conn.Do("MSET", "myfoo1", "mybar1", "myfoo2", "mybar2", "myfoo3", "mybar3")
	if err != nil {
		log.Fatalf("MSET error: %s", err.Error())
	}

	values, err := redis.Strings(conn.Do("MGET", "myfoo1", "myfoo5", "myfoo2", "myfoo3", "myfoo4"))
	if err != nil {
		log.Fatalf("MGET error: %s", err.Error())
	}

	for i := range values {
		fmt.Printf("reply[%d]: %s\n", i, values[i])
	}
}
