package model

import (
	"time"

	alias "github.com/iamJune20/dds/helper"
)

// App struct
type App struct {
	Code      string
	Name      string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  alias.NullTime
	Publish   string
}

// Apps type App list
type Apps []App

//NewApp App's Constructor
func NewApp() *App {
	return &App{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateApp() *App {
	return &App{
		UpdatedAt: time.Now(),
	}
}
