package redis

import (
	"context"
	"core/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func withRedis(f func(c *redis.Client) error) error {

	cos, err := config.ServiceNaCos()
	if err != nil {
		return err
	}
	r := cos.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", r.Host, r.Port),
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	err = f(rdb)
	if err != nil {
		return err
	}
	defer func(rdb *redis.Client) {
		err = rdb.Close()
		if err != nil {

		}
	}(rdb)
	return nil

}

func GetNxRedis(ctx context.Context, name string) error {

	err := withRedis(func(c *redis.Client) error {
		result, err := c.SetNX(ctx, name, "20", 15).Result()
		if err != nil && !result {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil

}

func AddRedis(ctx context.Context, mobile string, read string) error {

	err := withRedis(func(c *redis.Client) error {
		_, err := c.Set(ctx, mobile, read, time.Second*120).Result()
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
func GetRedis(ctx context.Context, mobile string) error {

	err := withRedis(func(c *redis.Client) error {
		_, err := c.Get(ctx, mobile).Result()

		return err
	})
	return err

}
