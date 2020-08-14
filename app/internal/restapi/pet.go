package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/PetStores/go-simple/internal/petstore/category"
	"github.com/PetStores/go-simple/internal/petstore/pet/datatype"
)

type petHandlers struct {
	categoryController category.Controller
}

func (h *petHandlers) addPet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pet := new(datatype.Pet)
		err := json.NewDecoder(r.Body).Decode(pet)
		if err != nil {
			ResponseBadRequest("Couldn't parse request body", w)
			return
		}

		// Data validation
		if pet.Category != nil {
			var opts []category.CategoryOption
			if pet.Category.ID != nil {
				opts = append(opts, category.WithID(*pet.Category.ID))
			}

			if pet.Category.Name != nil {
				opts = append(opts, category.WithName(*pet.Category.Name))
			}

			categ, err := h.categoryController.Validate(opts...)
			if err != nil {
				ResponseBadRequest("Couldn't validate category", w)
				return
			}

			if categ == nil {
				ResponseBadRequest("The category with given parameters doesn't exist", w)
				return
			}
		}

		// Pet creation
	}
}
