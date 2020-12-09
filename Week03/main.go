package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	errg, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second * 10)
		fmt.Fprintln(resp,"Hello GoGoGo!")
	})
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
	}

	//监听信号
	errg.Go(func() error {
		sign := make(chan os.Signal)
		signal.Notify(sign,syscall.SIGINT,syscall.SIGHUP,syscall.SIGTERM,syscall.SIGQUIT)
		select {
		case <- sign:
			fmt.Println("监听到信号；开始退出...")
			_ = server.Shutdown(context.Background())
			return errors.New("信号-结束")
		case <-ctx.Done():
			return nil
		}
	})

	//开启服务，监听端口
	errg.Go(func() error {
		return server.ListenAndServe()
	})

	//等待结束
	err := errg.Wait()
	fmt.Println("End...",err)
}
