package repository

import (
	"github.com/iamJune20/dds/src/modules/search/model"
)

type Repository interface {
	FindAll(string) (model.SearchAll, error)
	FindOne(string, string) (model.SearchOne, error)
}
