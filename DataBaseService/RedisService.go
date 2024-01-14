package DataBaseService

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SearchKey(key string) (string, error) {
	ctx := context.Background()
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exist
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		// Other error
		return "", err
	}
	return value, nil
}

func SetKey(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func CheckToken(UserId string, UserToken string) bool {
	ans, err := SearchKey(UserId)
	fmt.Println("Redis value==", ans)
	fmt.Println("UserToken==", UserToken)
	if err != nil {

	}
	return ans == UserToken
}
