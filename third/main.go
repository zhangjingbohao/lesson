package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := http.Server{
		Addr: ":80",
	}
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return server.ListenAndServe()
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				fmt.Println("exit !!")
				return server.Shutdown(ctx)
			}
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Println("end !", err)
		log.Fatal(err)
	}
}

/*
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
