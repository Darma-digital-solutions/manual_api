package repository

import (
	"github.com/iamJune20/dds/src/modules/app/model"
)

//AppRepository
type AppRepository interface {
	Save(*model.App) (string, error)
	Update(string, *model.App) (string, error)
	Delete(string) (string, error)
	FindByID(string) (*model.App, error)
	FindAll() (model.Apps, error)
}
