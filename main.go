package main

import (
	corebanking "core-system/service/core-banking"
	"os"
	"os/signal"
	"syscall"

	logs "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	coreBankingSvc := corebanking.InitServiceCoreBanking()
	serverErr := make(chan error, 1)
	go func() {
		serverErr <- coreBankingSvc.InitRoutes()
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		logs.Println("got an interrupt, exiting...")
	case err := <-serverErr:
		if err != nil {
			logs.Println("error while running api, exiting...", err)
		}
	}
}
