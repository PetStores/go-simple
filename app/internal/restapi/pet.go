package restapi

import (
	"encoding/json"
	"fmt"
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
		if petItem.ID != 0 {
			ResponseBadRequest("The identifier mustn't be set when creating pet", w)
			return
		}

		/*	if petItem.Category == 0 {
			ResponseBadRequest("The category must be set", w)
			return
		}*/

		if len(petItem.Name) < 2 {
			ResponseBadRequest("The name must be at least two letters", w)
			return
		}

		var opts []category.CategoryOption
		if petItem.Category.ID != nil {
			opts = append(opts, category.WithID(*petItem.Category.ID))
		}

		if petItem.Category.Name != nil {
			opts = append(opts, category.WithName(*petItem.Category.Name))
		}

		categ, err := h.categoryController.Validate(opts...)
		if err != nil {
			ResponseInternalError(err.Error(), w)
			// Here we can log the internal error, but we don't show it to the user
			return
		}

		if categ == nil {
			ResponseBadRequest("The category with given parameters doesn't exist", w)
			return
		}
		petItem.Category = *categ

		// Pet creation
		h.petController.AddPet(petItem)

		resp := APIResponse{
			Code: http.StatusCreated,
		}
		w.Header().Add("id", fmt.Sprintf("%d", petItem.ID))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}
