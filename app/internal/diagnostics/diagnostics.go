package diagnostics

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/PetStores/go-simple/internal/diagnostics/healthz"
)

// Diagnostics represents a diagnostics server.
type Diagnostics struct {
	server http.Server
	errors chan error
	logger *zap.SugaredLogger
}

// New returns a new instance of the Diagnostics server.
func New(logger *zap.SugaredLogger, port int, hc healthz.Check) *Diagnostics {
	http.HandleFunc("/healthz", healthzHandler(logger, hc))
	return &Diagnostics{
		server: http.Server{
			Addr:    net.JoinHostPort("", strconv.Itoa(port)),
			Handler: nil,
		},
		errors: make(chan error, 1),
		logger: logger,
	}
}

// Start diagnostics server.
func (d *Diagnostics) Start() {
	go func() {
		d.errors <- d.server.ListenAndServe()
		close(d.errors)
	}()
}

// Stop diagnostics server.
func (d *Diagnostics) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return d.server.Shutdown(ctx)
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel diagnostic you should stop the application.
func (d *Diagnostics) Notify() <-chan error {
	return d.errors
}

func healthzHandler(logger *zap.SugaredLogger, c healthz.Check) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resources := c()
		status := http.StatusOK
		for _, r := range resources {
			if r.Status == healthz.Fatal {
				status = http.StatusInternalServerError
				break
			}
		}

		err := json.NewEncoder(w).Encode(resources)
		if err != nil {
			logger.Errorw("Couldn't encode resources", "error", err, "data", resources)
		}

		w.WriteHeader(status)
	}
}
