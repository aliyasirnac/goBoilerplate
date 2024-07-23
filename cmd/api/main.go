package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aliyasirnac/goBackendBoilerplate/app"
	"github.com/aliyasirnac/goBackendBoilerplate/internal/config"
	"github.com/aliyasirnac/goBackendBoilerplate/internal/loggerx"
)

func main() {
	parentCtx := context.Background()
	closeChan := make(chan os.Signal, 2)
	signal.Notify(closeChan, syscall.SIGTERM, syscall.SIGINT)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("hata:", err)
	}
	loggerx.ExitOnError(err, "failed to load config")

	a := app.New(cfg)
	go func() {
		if err := a.Start(parentCtx); err != nil {
			loggerx.ExitOnError(err, "failed to start app")
		}
	}()

	// Sinyal yakalandığında Stop fonksiyonunu çağır
	sig := <-closeChan
	log.Printf("Caught signal %s: shutting down.", sig)

	if err := a.Stop(parentCtx); err != nil {
		loggerx.ExitOnError(err, "failed to stop app")
	}
}

