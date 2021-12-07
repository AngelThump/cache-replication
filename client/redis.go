package client

import (
	"context"

	utils "github.com/angelthump/cache-replication/utils"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func Initalize() {
	Rdb = redis.NewClient(&redis.Options{
		Network: utils.Config.Redis.Network,
		Addr:    utils.Config.Redis.Unix,
		DB:      0,
	})
}
