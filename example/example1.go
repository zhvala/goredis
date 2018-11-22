package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/zhvala/goredis"
)

const kNumOfRoutine = 50

func main() {
	conn, err := redis.NewConn(
		&redis.Options{
			StartNodes:   []string{"10.168.66.193:7001", "10.168.66.193:7000", "10.168.66.187:7002"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
	}

	chann := make(chan int, kNumOfRoutine)
	for i := 0; i < kNumOfRoutine; i++ {
		go redisTest(conn, i*100000, (i+1)*100000, chann)
	}

	for i := 0; i < kNumOfRoutine; i++ {
		_ = <-chann
	}
}

func redisTest(conn *redis.Conn, begin, end int, done chan int) {
	prefix := "mykey"
	for i := begin; i < end; i++ {
		key := prefix + strconv.Itoa(i)

		_, err := conn.Do("set", key, i*10)
		if err != nil {
			fmt.Printf("-set %s: %s\n", key, err.Error())
			time.Sleep(100 * time.Millisecond)
			continue
		}
		value, err := redis.Int(conn.Do("GET", key))
		if err != nil {
			fmt.Printf("-get %s: %s\n", key, err.Error())
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if value != i*10 {
			fmt.Printf("-mismatch %s: %d\n", key, value)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Printf("+set %s\n", key)
		time.Sleep(50 * time.Millisecond)
	}

	done <- 1
}
