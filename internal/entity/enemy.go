package entity

import (
	"math"
	"math/rand/v2"
)

type Enemy struct {
	Pos    Position
	HP     int
	Speed  float64
	Radius float64
	Damage int
	Alive  bool
}

func NewEnemy(playerX, playerY float64) *Enemy {
	angle := rand.Float64() * 2 * math.Pi
	dist := 300.0 + rand.Float64()*100.0
	return &Enemy{
		Pos:    Position{X: playerX + math.Cos(angle)*dist, Y: playerY + math.Sin(angle)*dist},
		HP:     3,
		Speed:  1.0 + rand.Float64()*0.5,
		Radius: 6,
		Damage: 10,
		Alive:  true,
	}
}

func (e *Enemy) Update(playerX, playerY float64) {
	dx := playerX - e.Pos.X
	dy := playerY - e.Pos.Y
	dist := math.Sqrt(dx*dx + dy*dy)
	if dist > 0 {
		e.Pos.X += (dx / dist) * e.Speed
		e.Pos.Y += (dy / dist) * e.Speed
	}
}
