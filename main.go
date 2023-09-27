package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"modulos/scenes"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Te odio teemo")

	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(1400, 617))

	scenes.NewMenuScene(window)
	window.ShowAndRun()
}
