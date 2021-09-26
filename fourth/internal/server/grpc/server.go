package grpc

import (
	pb "lesson/fourth/api"
	"lesson/fourth/configs"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.DemoServer) (ws *warden.Server, err error) {
	ws = warden.NewServer(configs.GlobalConfig.GrpcConfig)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
