package dao

import (
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"lesson/fourth/internal/dao/redis"
)

var Provider = wire.NewSet(New, redis.NewRedis)

type Dao interface {
}

type dao struct {
	redis *go_redis.ClusterClient
	cache *fanout.Fanout
}

func New(r *go_redis.ClusterClient) (d Dao, cf func(), err error) {
	return newDao(r)
}

func newDao(r *go_redis.ClusterClient) (d *dao, cf func(), err error) {
	d = &dao{
		redis: r,
		cache: fanout.New("cache"),
	}
	cf = d.Close
	return
}

func (d *dao) Close() {
	d.cache.Close()
}
