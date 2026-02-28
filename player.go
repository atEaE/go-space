package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y    float64
	HP      int
	MaxHP   int
	Level   int
	EXP     int
	NextEXP int
	Speed   float64
	Radius  float64
}

func NewPlayer() *Player {
	return &Player{
		X:       0,
		Y:       0,
		HP:      100,
		MaxHP:   100,
		Level:   1,
		EXP:     0,
		NextEXP: 10,
		Speed:   3.0,
		Radius:  8,
	}
}

func (p *Player) Update() {
	dx, dy := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		dx = 1
	}

	if dx != 0 && dy != 0 {
		len := math.Sqrt(dx*dx + dy*dy)
		dx /= len
		dy /= len
	}

	p.X += dx * p.Speed
	p.Y += dy * p.Speed
}

func (p *Player) Draw(screen *ebiten.Image, cam *Camera) {
	sx, sy := cam.WorldToScreen(p.X, p.Y)
	vector.FillCircle(screen, sx, sy, float32(p.Radius), color.White, true)
}

func (p *Player) TakeDamage(dmg int) {
	p.HP -= dmg
	if p.HP < 0 {
		p.HP = 0
	}
}

func (p *Player) AddEXP(amount int) {
	p.EXP += amount
	for p.EXP >= p.NextEXP {
		p.EXP -= p.NextEXP
		p.LevelUp()
	}
}

func (p *Player) LevelUp() {
	p.Level++
	p.NextEXP = p.Level * 10
	p.Speed += 0.2
	p.MaxHP += 10
	p.HP = p.MaxHP
}
