package frothy

import "fmt"

type App struct{}

func (app *App) AddQR() {
	fmt.Println("AddQR")
}

func (app *App) AddCode() {
	fmt.Println("AddCode")
}
