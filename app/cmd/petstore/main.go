package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PetStores/go-simple/internal/diagnostics"
	"github.com/PetStores/go-simple/internal/resources"
	"github.com/PetStores/go-simple/internal/restapi"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	slogger := logger.Sugar()

	// Timeout context to shutdown resources and servers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsc, err := resources.New(ctx, logger)
	if err != nil {
		slogger.Fatalw("Can't initialize resources.", "err", err)
	}

	rapi := restapi.New(slogger, rsc.Config.RESTAPIPort)
	rapi.Start()

	diag := diagnostics.New(slogger, rsc.Config.DiagPort, rsc.Healthz)
	diag.Start()

	slogger.Info("The servers are ready.")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		slogger.Infow("Received a signal.", "signal", x.String())
	case err := <-diag.Notify():
		slogger.Errorw("Received an error from the diagnostics server.", "err", err)
	case err := <-rapi.Notify():
		slogger.Errorw("Received an error from the business logic server.", "err", err)
	}

	slogger.Info("Stopping the servers...")
	err = rapi.Stop()
	if err != nil {
		slogger.Error("Got an error while stopping the business logic server.", "err", err)
	}

	err = diag.Stop()
	if err != nil {
		slogger.Error("Got an error while stopping the diag logic server.", "err", err)
	}

	slogger.Info("The app is calling the last defers and will be stopped.")
}
