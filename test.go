package main

import (
	"testing"

	_ "github.com/iamJune20/dds/models"
)

func modelApp(t *testing.T) {
	t.Logf("Out:", models.getAllAppModel())
}
