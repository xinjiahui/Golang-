package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

//redis

var redisdb *redis.Client
var ctx = context.Background()

func initRedis() (err error) {
	ctx := context.Background()
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping(ctx).Result()
	return

}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed,err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功")
	//zset
	key := "rank"
	redisdb.ZAdd(ctx, key, redis.Z{
		Score: 90, Member: "PHP",
	})

}
