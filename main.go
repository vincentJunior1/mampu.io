package main

import (
	"context"
	"core-system/core/utils/getenv"
	corebanking "core-system/service/core-banking"
	"os"
	"os/signal"
	"syscall"
	"time"

	logs "github.com/sirupsen/logrus"
)

func main() {
	if err := getenv.LoadEnv(".env"); err != nil {
		panic(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	coreBankingSvc := corebanking.InitServiceCoreBanking()
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- coreBankingSvc.InitRoutes()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(getenv.Getenv[int]("SHUTDOWN_TIMEOUT"))*time.Second)
		defer cancel()
		if err := coreBankingSvc.Stop(shutdownCtx); err != nil {
			logs.Println("error while stopping service:", err)
		}
		logs.Println("got an interrupt, exiting...")
	case err := <-serverErr:
		if err != nil {
			logs.Println("error while running api, exiting...", err)
		}
	}
}
