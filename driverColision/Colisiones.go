package driverColision

import "modulos/models"

type CollisionDriver struct {
	player   *models.Player
	meteoro  *models.Meteoro
	gameOver bool
}

func NewCollisionDriver(player *models.Player, meteoro *models.Meteoro) *CollisionDriver {
	return &CollisionDriver{
		player:   player,
		meteoro:  meteoro,
		gameOver: false,
	}
}

func (c *CollisionDriver) Run() {
	for !c.gameOver {
		if c.meteoro.GetPosY() >= 400 {
			if c.meteoro.GetPosX() >= c.player.GetPosX()-50 && c.meteoro.GetPosX() <= c.player.GetPosX()+50 {
				c.meteoro.SetRunning(false)
				c.player.SetRunning(false)
				c.gameOver = true
			}
		}
	}
}

func (c *CollisionDriver) GetGameOver() bool {
	return c.gameOver
}
