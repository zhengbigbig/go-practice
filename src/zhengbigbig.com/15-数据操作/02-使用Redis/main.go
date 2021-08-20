package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Silvanware*123",
		DB:       0,
		PoolSize: 100, // 连接池数量
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return err
}

func main() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		fmt.Println(err)
		return
	}
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2,err := rdb.Get(ctx,"key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	}else if err != nil {
		panic(err)
	}else {
		fmt.Println("key2",val2)
	}
	// Output: key value
	// key2 does not exist
}
