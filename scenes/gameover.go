package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
	window fyne.Window
}

func NewGameOverScene(fyneWindow fyne.Window) *GameOverScene {
	scene := &GameOverScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *GameOverScene) Render() {
	gameOverImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/League-defeat.jpg"))
	gameOverImage.Resize(fyne.NewSize(1400, 617))
	gameOverImage.Move(fyne.NewPos(0, 0))

	btnGoMenu := widget.NewButton("Menu", s.goMenu)
	btnGoMenu.Resize(fyne.NewSize(150, 30))
	btnGoMenu.Move(fyne.NewPos(670, 50))

	btnRestart := widget.NewButton("Restart", s.restart)
	btnRestart.Resize(fyne.NewSize(150, 30))
	btnRestart.Move(fyne.NewPos(670, 10))

	s.window.SetContent(container.NewWithoutLayout(gameOverImage, btnGoMenu, btnRestart))
}

func (s *GameOverScene) goMenu() {
	NewMenuScene(s.window)
}
func (s *GameOverScene) restart() {
	NewGameScene(s.window)
}
