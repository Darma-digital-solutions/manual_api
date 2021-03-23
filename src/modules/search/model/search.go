package model

import (
	"github.com/iamJune20/dds/src/modules/content/model"
)

type SearchOne struct {
	AppCode      string
	AppName      string
	AppLogo      string
	CountContent int
	Content      model.Contents
}

type SearchAll []SearchOne
