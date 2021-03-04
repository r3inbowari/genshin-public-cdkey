package genshin_public_cdkey

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
	println("[INFO] redis client init")
	rdb = redis.NewClient(&redis.Options{
		Addr:     "119.23.175.142:6379",
		Password: "159463",
		DB:       9,
	})
}

func PushOK(id, md5 string) {
	Set("pushed:"+id+":"+md5, "ok") //bitmap......
}

func OK(id, md5 string) bool {
	return Get("pushed:"+id+":"+md5) == "ok"
}

func Get(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

func Set(key, val string) {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		println("error set")
	}
}
