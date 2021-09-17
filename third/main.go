package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", handleTest)

	server := http.Server{
		Addr: "addr",
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			if err := server.Shutdown(ctx); err != nil {
				log.Panicln("main.Shutdown")
			}
			cancel()
			return
		default:
			return
		}
	}
}

func handleTest(w http.ResponseWriter, r *http.Request) {

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return errors.New("1")
	})
	g.Go(func() error {
		return errors.New("2")
	})

	if err := g.Wait(); err != nil {
		log.Println("handleTest.Wait", err)
	}
	_, err := w.Write([]byte("response"))
	if err != nil {
		log.Println("handleTest.Write", err)
	}
}

/*
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
