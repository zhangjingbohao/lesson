// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package dao

import (
	"github.com/google/wire"
	"lesson/fourth/internal/dao/redis"
)

//go:generate wire
func newTestDao() (*dao, func(), error) {
	panic(wire.Build(newDao, redis.NewRedis))
}
