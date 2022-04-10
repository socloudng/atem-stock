package cache

import (
	"atem-stock/configs"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func getRedisClient() *redis.Client {
	config := configs.AppConfigInstance.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pwd,
		DB:       config.DB,
	})
	return client
}

func redisHMSet(key string, stockMaps map[string]interface{}) {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := client.HMSet(ctx, key, stockMaps)
	log.Println(cmd.Result())
}

func redisHMGet(key string, fileds []string) []interface{} {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := client.HMGet(ctx, key, fileds...).Result()
	if err != nil {
		log.Println("获取Value失败")
		return nil
	}
	return result
}

func redisHGetAll(key string) map[string]string {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		log.Println("获取Field失败")
		return nil
	}
	return result
}

func redisHGetAllFields(key string) []string {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := client.HKeys(ctx, key).Result()
	if err != nil {
		log.Println("获取Field失败")
		return nil
	}
	return result
}

func redisMSet(stockMaps map[string]interface{}) {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := client.MSet(ctx, stockMaps)
	log.Println(cmd.Result())
}

func redisMGet(key string) {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := client.MGet(ctx, key)
	log.Println(cmd.Result())
}

func redisMDel(keys []string) {
	client := getRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client.Del(ctx, keys...)
}
