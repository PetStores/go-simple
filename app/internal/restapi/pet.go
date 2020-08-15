package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/PetStores/go-simple/internal/petstore/category"
	"github.com/PetStores/go-simple/internal/petstore/pet"
	"github.com/PetStores/go-simple/internal/petstore/pet/datatype"
)

type petHandlers struct {
	categoryController *category.Controller
	petController      *pet.Controller
}

func (h *petHandlers) addPet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		petItem := new(datatype.Pet)
		err := json.NewDecoder(r.Body).Decode(petItem)
		if err != nil {
			ResponseBadRequest("Couldn't parse request body", w)
			return
		}

		// Data validation
		if petItem.Category != nil {
			var opts []category.CategoryOption
			if petItem.Category.ID != nil {
				opts = append(opts, category.WithID(*petItem.Category.ID))
			}

			if petItem.Category.Name != nil {
				opts = append(opts, category.WithName(*petItem.Category.Name))
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
		h.petController.AddPet(petItem)

		resp := APIResponse{
			Code: http.StatusCreated,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}
