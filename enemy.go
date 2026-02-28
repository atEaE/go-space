package main

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	X, Y   float64
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
		X:      playerX + math.Cos(angle)*dist,
		Y:      playerY + math.Sin(angle)*dist,
		HP:     3,
		Speed:  1.0 + rand.Float64()*0.5,
		Radius: 6,
		Damage: 10,
		Alive:  true,
	}
}

func (e *Enemy) Update(playerX, playerY float64) {
	dx := playerX - e.X
	dy := playerY - e.Y
	dist := math.Sqrt(dx*dx + dy*dy)
	if dist > 0 {
		e.X += (dx / dist) * e.Speed
		e.Y += (dy / dist) * e.Speed
	}
}

func (e *Enemy) Draw(screen *ebiten.Image, cam *Camera) {
	sx, sy := cam.WorldToScreen(e.X, e.Y)
	vector.FillCircle(screen, sx, sy, float32(e.Radius), color.RGBA{R: 220, G: 40, B: 40, A: 255}, true)
}
