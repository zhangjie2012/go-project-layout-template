package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zhangjie2012/go-project-layout-template/cmd/options"
	"github.com/zhangjie2012/go-project-layout-template/pkg/cache"
	"github.com/zhangjie2012/go-project-layout-template/pkg/store"
)

type Server struct {
	HttpServer *http.Server

	Store *store.Store
	Cache *cache.Cache
}

type ServerOption func(*Server)

func WithStore(storeX *store.Store) ServerOption {
	return func(s *Server) {
		s.Store = storeX
	}
}

func WithCache(cacheX *cache.Cache) ServerOption {
	return func(s *Server) {
		s.Cache = cacheX
	}
}

func NewServer(config *options.AppOption, opts ...ServerOption) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{
		HttpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: router,
		},
	}
	for _, opt := range opts {
		opt(s)
	}
	s.RegisterRouter(router)

	return s
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go s.Run(ctx, wg)
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// graceful close
	defer func() {
		waiting := 5 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), waiting)
		defer cancel()

		if err := s.HttpServer.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failure, error=%s", err)
		}

		log.Infof("server exiting")
	}()

	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Infof("server run on: %s", s.HttpServer.Addr)

	select {
	case <-ctx.Done():
		return
	}
}
