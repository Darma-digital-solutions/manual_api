package repository

import (
	"github.com/iamJune20/dds/src/modules/content/model"
)

type ContentRepository interface {
	Save(*model.Content) (string, error)
	Update(string, *model.Content) (string, error)
	Delete(string) (string, error)
	FindByID(string) (*model.Content, error)
	FindAll() (model.Contents, error)
	FindOne(string, string) (model.Contents, error)
}
