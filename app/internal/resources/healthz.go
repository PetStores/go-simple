package resources

import (
	"time"

	"github.com/PetStores/go-simple/internal/diagnostics/healthz"
)

func (r *R) Healthz() []healthz.Resource {
	var err error

	dbStatus := healthz.Ok
	dbMsg := "It works!"
	for i := 0; i < 5; i++ {
		_, err = r.DB.Query("SELECT 1")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		dbStatus = healthz.Fatal
		dbMsg = err.Error()
	}

	return []healthz.Resource{
		{
			Name:    "reformDB",
			Status:  dbStatus,
			Message: dbMsg,
		},
	}
}
