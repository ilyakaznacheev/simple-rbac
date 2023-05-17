package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/ilyakaznacheev/simple-rbac/internal/api"
	"github.com/ilyakaznacheev/simple-rbac/internal/service"
	"github.com/ilyakaznacheev/simple-rbac/internal/storage/inmemory"
)

var (
	port      = flag.String("port", "8080", "gRPC port")
	printHelp = flag.Bool("h", false, "Print help and exit")
)

func main() {
	flag.Parse()
	if *printHelp {
		flag.Usage()
		os.Exit(0)
	}

	l, _ := zap.NewProduction()
	zap.ReplaceGlobals(l)

	s := inmemory.New()

	app := service.New(s)

	srv := api.New(app)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		zap.L().Info("starting server", zap.String("port", *port))
		return srv.Serve(net.JoinHostPort("0.0.0.0", *port))
	})
	eg.Go(func() error {
		// shutdown server
		<-ctx.Done()
		zap.L().Info("stopping server")
		defer zap.L().Info("server stopped")
		return srv.Close()
	})

	_ = eg.Wait()
}
