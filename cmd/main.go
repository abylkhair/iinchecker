package main

import (
	"context"
	server "github.com/wildegor/kaspi-rest/internal"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer done()

	srv, _ := server.NewServer()
	srv.Run(ctx)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	srv.Shutdown(ctx)
}
