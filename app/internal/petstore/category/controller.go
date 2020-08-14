package category

import (
	"fmt"

	"github.com/PetStores/go-simple/internal/petstore/category/datatype"
)

type Controller struct {
	cf CategoryFinder
}

type CategoryFinder interface {
	FindCategory(params map[string]interface{}) (*datatype.Category, error)
}

func NewController(cf CategoryFinder) *Controller {
	return &Controller{
		cf: cf,
	}
}

func (c *Controller) Validate(opts ...CategoryOption) (*datatype.Category, error) {
	o := &CategoryOptions{
		queryParams: make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(o)
	}

	if len(o.queryParams) == 0 {
		return nil, fmt.Errorf("can't validate category: none of the parameters are set")
	}

	category, err := c.cf.FindCategory(o.queryParams)
	if err != nil {
		return nil, fmt.Errorf("can't validate category: %w", err)
	}

	return category, nil
}
