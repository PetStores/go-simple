package withdb

import (
	"fmt"
	"strings"

	"github.com/PetStores/go-simple/internal/petstore/category/datatype"

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

func (wdb *WithDB) FindCategory(params map[string]interface{}) (*datatype.Category, error) {
	intersect := make([]string, len(params))
	args := make([]interface{}, len(params))

	i := 0
	for key, arg := range params {
		intersect[i] = fmt.Sprintf("%s = $%d", key, i+1)
		args[i] = arg
		i++
	}

	tail := "WHERE " + strings.Join(intersect, " AND ")

	sts, err := wdb.db.SelectAllFrom(categoryTable, tail, args...)
	if err != nil {
		return nil, err
	} else if len(sts) > 1 {

	} else if len(sts) == 0 {
		return nil, nil
	}

	internal := sts[0].(*category)
	external := datatype.Category{
		ID:   &internal.ID,
		Name: &internal.Name,
	}

	return &external, nil
}
