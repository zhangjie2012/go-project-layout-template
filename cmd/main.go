package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/zhangjie2012/go-project-layout-template/cmd/options"
	"github.com/zhangjie2012/go-project-layout-template/pkg/api"
	"github.com/zhangjie2012/go-project-layout-template/pkg/cache"
	"github.com/zhangjie2012/go-project-layout-template/pkg/store"
)

var (
	config = "/etc/config.yaml"
)

func init() {
	flag.StringVar(&config, "conf", config, "the server configure file")

	callerPrettyfier := func(frame *runtime.Frame) (function string, file string) {
		ss := strings.Split(frame.Function, ".")
		function = ss[len(ss)-1]
		file = fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		return function, file
	}

	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: callerPrettyfier,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		FullTimestamp:    true,
	})
	log.SetLevel(log.TraceLevel)
}

func main() {
	flag.Parse()

	_, err := options.ParseOption(config)
	if err != nil {
		log.Fatalf("parse config file failure, error=%s", err)
	}

	Serve()
}

func Serve() {
	opt := options.GetOption()

	// init/release resource
	storeX, err := store.NewStore(opt)
	if err != nil {
		log.Errorf("new store failure, error=%s", err)
		return
	}
	defer storeX.Close()

	cacheX, err := cache.NewCache(opt)
	if err != nil {
		log.Errorf("new cache failure, error=%s", err)
		return
	}
	defer cacheX.Close()

	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start server
	server := api.NewServer(opt)
	server.Start(ctx, &wg)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Infof("receiver shutdown signal")
}
