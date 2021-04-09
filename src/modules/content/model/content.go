package model

import (
	"time"

	alias "github.com/iamJune20/dds/helper"
)

type Content struct {
	Code         string
	Title        string
	Desc         string
	CategoryCode string
	ManualCode   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     alias.NullTime
	Publish      string
}

type Contents []Content

func NewContent() *Content {
	return &Content{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateContent() *Content {
	return &Content{
		UpdatedAt: time.Now(),
	}
}
