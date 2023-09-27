package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Character interface {
	GoRight()
	GoLeft()
	Run()
	SetRunning(pause bool)
	GetRunning() bool
	GetPosY() float32
	GetPosX() float32
}

type Player struct {
	posX, posY, direction float32
	running               bool
	pel                   *canvas.Image
	isFocused             bool
	keyPressTime          time.Time
	keyPressDelay         time.Duration
}

func NewPlayer(posx float32, posy float32, img *canvas.Image) Character {
	return &Player{
		posX:      posx,
		posY:      posy,
		running:   true,
		pel:       img,
		direction: 0,
	}
}

// Implementación de la interfaz fyne.Focusable

func (p *Player) FocusGained() {
	p.isFocused = true
}

func (p *Player) FocusLost() {
	p.isFocused = false
}

func (p *Player) TypedRune(r rune) {
	if p.isFocused {
		switch r {
		case 'a':
			p.GoLeft()
		case 'd':
			p.GoRight()
		}
	}
}

func (t *Player) GoRight() {
	t.direction = 1
}

func (t *Player) GoLeft() {
	t.direction = -1
}

func (t *Player) Run() {
	for t.running {
		var incX float32 = 50
		incX *= t.direction

		if t.posX < 0 {
			t.posX = 0
		} else if t.posX > 1300 {
			t.posX = 1300
		}

		t.posX += incX
		t.pel.Move(fyne.NewPos(t.posX, t.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (t *Player) SetRunning(pause bool) {
	t.running = pause
}

func (t *Player) GetRunning() bool {
	return t.running
}

func (t *Player) GetPosY() float32 {
	return t.posY
}

func (t *Player) GetPosX() float32 {
	return t.posX
}
