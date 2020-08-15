package withdb

import (
	"fmt"

	"github.com/PetStores/go-simple/internal/petstore/pet/datatype"
	"gopkg.in/reform.v1"
)

type WithDB struct {
	db *reform.DB
}

func New(db *reform.DB) *WithDB {
	return &WithDB{
		db: db,
	}
}

func (wdb *WithDB) ReadPet(id int64) (*datatype.Pet, error) {
	rec, err := wdb.db.FindByPrimaryKeyFrom(petTable, id)
	if err != nil {
		return nil, fmt.Errorf("couldn't find pet by primary key: %w", err)
	}

	internal := rec.(*pet)
	external := datatype.Pet{
		ID: internal.ID,
	}

	return &external, nil
}

func (wdb *WithDB) WritePet(pet *datatype.Pet) error {
	return fmt.Errorf("the method is not implemented")
}
