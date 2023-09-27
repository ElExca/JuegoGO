package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var t *Player

type Meteoro struct {
	posX, posY, direction float32
	running               bool
	pel                   *canvas.Image
}

func NewMeteoro(posx float32, posy float32, img *canvas.Image, tux *Player) *Meteoro {
	t = tux
	return &Meteoro{
		posX:    posx,
		posY:    posy,
		running: true,
		pel:     img,
	}
}

func (w *Meteoro) Run() {
	for w.running {
		var inc float32 = 50

		if w.posY > 450 {
			w.posY = -50
			w.posX = float32(rand.Intn(1400))
		}

		w.posY += inc
		w.pel.Move(fyne.NewPos(w.posX, w.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (w *Meteoro) SetRunning(pause bool) {
	w.running = pause
}
func (w *Meteoro) GetRunning() bool {
	return w.running
}
func (w *Meteoro) GetPosY() float32 {
	return w.posY
}
func (w *Meteoro) GetPosX() float32 {
	return w.posX
}
