package xredis

import (
	"context"
	"fmt"
	"github.com/prolifera/kit/log"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var rdb = &redis.Client{}

func Default() *redis.Client {
	if rdb == nil {
		return &redis.Client{}
	}
	return rdb
}

func NewClient(svc string, db string) (*redis.Client, error) {
	var redisDB int64
	var err error
	if db == "" || db == "0" {
		redisDB = 0
	} else {
		redisDB, err = strconv.ParseInt(db, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("get redis db error: %s", err.Error())
		}
	}

	cli := redis.NewClient(&redis.Options{
		Addr: svc,
		DB:   int(redisDB),
	})
	_, err = cli.Ping(context.Background()).Result()
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "init redis client error", "error": err.Error()}).Send()
		return nil, err
	}

	return cli, nil
}

func InitRedis(svc string, db string) error {

	cli, err := NewClient(svc, db)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"action": "new client error", "error": err.Error()}).Send()
		return err
	}

	rdb = cli

	log.Info().Msgf("init redis client addr %s db %d success", svc, db)

	return nil
}
