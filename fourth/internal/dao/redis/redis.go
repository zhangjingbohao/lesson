package redis

import (
	"github.com/go-redis/redis/v8"
	"lesson/fourth/configs"
)

func NewRedis() (r *redis.ClusterClient, cf func(), err error) {
	r = redis.NewClusterClient(configs.GlobalConfig.RedisConfig)
	cf = func() { r.Close() }
	return
}
