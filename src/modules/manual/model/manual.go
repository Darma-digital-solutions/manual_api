package model

import (
	"time"

	alias "github.com/iamJune20/dds/helper"
)

type Manual struct {
	Code      string
	Name      string
	Desc      string
	Icon      string
	AppCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  alias.NullTime
	Publish   string
}

type Manuals []Manual

func NewManual() *Manual {
	return &Manual{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateManual() *Manual {
	return &Manual{
		UpdatedAt: time.Now(),
	}
}
