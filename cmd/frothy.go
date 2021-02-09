package main

import (
	"github.com/fcjr/frothy"
	"github.com/sqweek/dialog"
)

func main() {
	app, err := frothy.NewApp()
	if err != nil {
		dialog.Message(err.Error()).Error()
	}
	if err := app.RunUI(); err != nil {
		dialog.Message(err.Error()).Error()
	}
}