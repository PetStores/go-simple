package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/PetStores/go-simple/internal/petstore/category"

	"github.com/PetStores/go-simple/internal/petstore/pet/datatype"
)

func addPet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pet := new(datatype.Pet)
		err := json.NewDecoder(r.Body).Decode(pet)
		if err != nil {
			resp := APIResponse{
				Code:    http.StatusBadRequest,
				Message: "Couldn't parse request body",
			}
			json.NewEncoder(w).Encode(resp)
		}

		// Data validation
		if pet.Category != nil {
			opts := []category.CategoryOption
			if pet.Category.ID != nil {
				opts = append(opts, category.WithID(*pet.Category.ID))
			}

			if pet.Category.Name != nil {
				opts = append(opts, category.WithID(*pet.Category.Name))
			}


		}

		// Pet creation
	}
}
