package service

import (
	pb "lesson/fourth/api"

	"github.com/google/wire"
	"lesson/fourth/internal/dao"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

type Service struct {
	dao dao.Dao
}

func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = func() {}
	return
}
