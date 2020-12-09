package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, _ := errgroup.WithContext(ctx)

	http1 := NewHttpServer(":9190")
	http2 := NewHttpServer(":9191")
	http3 := NewHttpServer(":9192")

	// 启动http1
	g.Go(func() error {
		return http1.Start()
	})

	// 启动http2
	g.Go(func() error {
		return http2.Start()
	})

	// 启动http3
	g.Go(func() error {
		return http3.Start()
	})

	// 监听signal信号
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for {
			select {
			// 收到signal信号
			case s := <-c:
				log.Printf("get a signal %s, trigger shutdown.", s)
				cancel()
			case <-ctx.Done():
				ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Microsecond)
				defer cancel()
				if err := http1.Shutdown(ctx); err != nil {
					return err
				}
				if err := http2.Shutdown(ctx); err != nil {
					return err
				}
				if err := http3.Shutdown(ctx); err != nil {
					return err
				}
				return context.Canceled
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}

type httpServer struct {
	server http.Server
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		server: http.Server{
			Addr: addr,
		},
	}
}

func (s *httpServer) Start() error {
	log.Printf("server 0.0.0.0%s start.", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	log.Printf("server 0.0.0.0%s shutdown.", s.server.Addr)
	return s.server.Shutdown(ctx)
}


