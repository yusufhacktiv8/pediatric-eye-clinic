package controllers

import (
	"os"
	"testing"

	"github.com/appleboy/gofight"
)

var a AppTest

func TestMain(m *testing.M) {
	a = AppTest{}
	a.Initialize("", "", "")
	defer a.DB.Close()

	a.GoFight = gofight.New()

	theRun := m.Run()

	os.Exit(theRun)
}

func GetAppTest() *AppTest {
	return &a
}
