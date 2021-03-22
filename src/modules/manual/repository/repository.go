package repository

import (
	"github.com/iamJune20/dds/src/modules/manual/model"
)

type ManualRepository interface {
	Save(*model.Manual) (string, error)
	Update(string, *model.Manual) (string, error)
	Delete(string) (string, error)
	FindByID(string) (*model.Manual, error)
	FindByAppCode(string) (*model.Manuals, error)
	FindAll() (model.Manuals, error)
}
