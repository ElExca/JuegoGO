package scenes

import (
	"modulos/driverColision"
	"modulos/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
}

var player *models.Player
var meteoro *models.Meteoro
var collisionDriver *driverColision.CollisionDriver
var bird *canvas.Image // Agregamos una variable para el pájaro

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
	scene.Render()
	scene.StartGame()
	return scene
}

func (s *GameScene) Render() {
	// Cargar imagen de fondo
	backgroundImage := loadImageFromURI("./assets/img.jpeg", 1400, 617)
	backgroundImage.Move(fyne.NewPos(0, 0))

	// Renderizar jugador
	playerImage := loadImageFromURI("./assets/teemo.png", 100, 100)
	player = models.NewPlayer(350, 450, playerImage).(*models.Player)

	// Renderizar meteoro
	meteoroImage := loadImageFromURI("./assets/meteoro.png", 100, 100)
	meteoro = models.NewMeteoro(350, 600, meteoroImage, player)

	// Inicializar CollisionDriver
	collisionDriver = driverColision.NewCollisionDriver(player, meteoro)

	// Renderizar el pájaro
	birdImage := loadImageFromURI("./assets/Anivia.png", 100, 100) // Asegúrate de tener una imagen del pájaro
	bird = birdImage
	go moveBird() // Lanzamos una goroutine para mover el pájaro

	s.window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if player.GetRunning() {
			switch ev.Name {
			case "Left":
				player.GoLeft()
			case "Right":
				player.GoRight()
			}
		}
	})

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, playerImage, meteoroImage, birdImage))
}

func moveBird() {
	screenWidth := float32(1400) // Ancho de la pantalla
	birdWidth := float32(50)     // Ancho del pájaro

	for {
		x := -birdWidth // Inicialmente, establece la posición X del pájaro fuera de la pantalla a la izquierda
		for x <= screenWidth {
			bird.Move(fyne.NewPos(x, 100)) // Mueve el pájaro a la nueva posición
			x += 10                        // Ajusta la velocidad de movimiento del pájaro aquí
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(2000 * time.Millisecond) // Espera antes de reiniciar el vuelo
	}
}

func (s *GameScene) StartGame() {
	go player.Run()
	go meteoro.Run()
	go collisionDriver.Run()
	go s.checkGameOver()
}

func (s *GameScene) StopGame() {
	player.SetRunning(!player.GetRunning())
	meteoro.SetRunning(!meteoro.GetRunning())
}

func (s *GameScene) checkGameOver() {
	running := true
	for running {
		if collisionDriver.GetGameOver() {
			running = false
			time.Sleep(1000 * time.Millisecond)
			NewGameOverScene(s.window)
		}
	}
}

func loadImageFromURI(fileURI string, sizeX, sizeY float32) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileURI))
	image.Resize(fyne.NewSize(sizeX, sizeY))
	return image
}
