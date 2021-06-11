package datatype

import "github.com/PetStores/go-simple/internal/petstore/category/datatype"

type Pet struct {
	ID       int64             `json:"id"`
	Name     string            `json:"name"`
	Category datatype.Category `json:"category"`
}
