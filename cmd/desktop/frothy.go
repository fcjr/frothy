package main

import (
	"github.com/fcjr/alert"
	"github.com/fcjr/frothy/desktop"
)

func main() {
	app, err := desktop.NewApp()
	if err != nil {
		alert.Error("Error", err.Error())
	}
	if err := app.RunUI(); err != nil {
		alert.Error("Error", err.Error())
	}
}
