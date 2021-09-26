// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"lesson/fourth/internal/dao"
	"lesson/fourth/internal/server/grpc"
	"lesson/fourth/internal/server/http"
	"lesson/fourth/internal/service"

	"github.com/google/wire"
)

//go:generate wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
