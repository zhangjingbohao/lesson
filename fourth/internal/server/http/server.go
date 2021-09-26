package http

import (
	pb "lesson/fourth/api"
	"lesson/fourth/configs"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.DemoServer

func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	svc = s
	engine = bm.DefaultServer(configs.GlobalConfig.HttpConfig)
	pb.RegisterDemoBMServer(engine, s)
	//initRouter(engine)
	err = engine.Start()
	return
}
