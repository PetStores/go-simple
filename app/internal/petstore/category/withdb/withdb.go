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

func (wdb *WithDB) FindCategory(params map[string]interface{}) (*datatype.Category, error) {
	intersect := make([]string, len(params))
	args := make([]interface{}, len(params))

	i := 0
	for key, arg := range params {
		intersect[i] = fmt.Sprintf("%s = $%d", key, i)
		args[i] = arg
	}

	tail := "WHERE " + strings.Join(intersect, " AND ")

	sts, err := wdb.db.SelectAllFrom(categoryTable, tail, args...)
	if err != nil {
		//
	} else if len(sts) == 0 {
		return nil, fmt.Errorf("couldn't find category with the given parameters")
	} else if len(sts) > 1 {

	} else {

	}

	//
}
