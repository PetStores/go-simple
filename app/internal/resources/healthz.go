package resources

import "github.com/PetStores/go-simple/internal/diagnostics/healthz"

func (r *R) Healthz() []healthz.Resource {
	return []healthz.Resource{}
}
