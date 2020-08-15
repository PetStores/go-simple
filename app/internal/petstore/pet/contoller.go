package pet

import (
	"fmt"

	"github.com/PetStores/go-simple/internal/petstore/pet/datatype"
)

type Controller struct {
	rw PetReadWriter
}

type PetReadWriter interface {
	ReadPet(id int64) (*datatype.Pet, error)
	WritePet(pet *datatype.Pet) error
}

func NewController(rw PetReadWriter) *Controller {
	return &Controller{
		rw: rw,
	}
}

func (c *Controller) AddPet(pet *datatype.Pet) error {
	return fmt.Errorf("the method is not implemented")
}
