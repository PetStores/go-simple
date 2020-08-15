package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/PetStores/go-simple/internal/diagnostics"
	categoryc "github.com/PetStores/go-simple/internal/petstore/category"
	categorydp "github.com/PetStores/go-simple/internal/petstore/category/withdb"
	petc "github.com/PetStores/go-simple/internal/petstore/pet"
	petdp "github.com/PetStores/go-simple/internal/petstore/pet/withdb"
	"github.com/PetStores/go-simple/internal/resources"
	"github.com/PetStores/go-simple/internal/restapi"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")

	rsc, err := resources.New(slogger)
	if err != nil {
		slogger.Fatalw("Can't initialize resources.", "err", err)
	}
	defer func() {
		err = rsc.Release()
		if err != nil {
			slogger.Errorw("Got an error during resources release.", "err", err)
		}
	}()

	slogger.Info("Configuring the application units...")
	catdb := categorydp.New(rsc.DB)
	cc := categoryc.NewController(catdb)

	petdb := petdp.New(rsc.DB)
	pc := petc.NewController(petdb)

	slogger.Info("Starting the servers...")
	rapi := restapi.New(slogger, rsc.Config.RESTAPIPort, cc, pc)
	rapi.Start()

	diag := diagnostics.New(slogger, rsc.Config.DiagPort, rsc.Healthz)
	diag.Start()
	slogger.Info("The application is ready to serve requests.")

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
