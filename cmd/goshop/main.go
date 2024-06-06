package main

import (
	"context"
	"fmt"
	"github.com/iypetrov/goshop/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iypetrov/goshop/internal/config/api"
	"github.com/iypetrov/goshop/internal/config/db"
)

func main() {
	ctx := context.Background()
	if err := Run(ctx); err != nil {
		return
	}
}

func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	config.New()

	conn, err := db.InitDatabaseConnectionPool(ctx)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	if err := db.RunDBSchemaMigration(conn); err != nil {
		panic(err.Error())
	}

	s := api.NewServer(ctx, conn)
	fmt.Printf("server started on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	<-setupGracefulShutdown(cancel)
	return nil
}

func setupGracefulShutdown(cancel context.CancelFunc) (shutdownCompleteChan chan struct{}) {
	shutdownCompleteChan = make(chan struct{})
	isFirstShutdownSignal := true

	shutdownFunc := func() {
		if !isFirstShutdownSignal {
			log.Println("caught another exit signal, now hard dying")
			os.Exit(1)
		}

		isFirstShutdownSignal = false
		log.Println("starting graceful shutdown")

		cancel()

		close(shutdownCompleteChan)
	}

	go func(shutdownFunc func()) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			log.Println("caught exit signal", "signal", <-sigint)
			go shutdownFunc()
		}
	}(shutdownFunc)

	return shutdownCompleteChan
}
