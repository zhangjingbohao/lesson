package main

import (
	"flag"
	"lesson/fourth/configs"
	"lesson/fourth/internal/di"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil)
	defer log.Close()
	configs.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			closeFunc()
			return
		default:
			return
		}
	}
}
