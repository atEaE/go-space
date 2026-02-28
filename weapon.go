package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Weapon struct {
	Cooldown    int
	Timer       int
	BulletSpeed float64
	Damage      int
}

func NewWeapon() *Weapon {
	return &Weapon{
		Cooldown:    20,
		Timer:       0,
		BulletSpeed: 5.0,
		Damage:      1,
	}
}

func (w *Weapon) Update(playerX, playerY float64, enemies []*Enemy, bullets *[]*Bullet, level int) {
	w.Timer--
	if w.Timer > 0 {
		return
	}

	var nearest *Enemy
	minDist := math.MaxFloat64
	for _, e := range enemies {
		if !e.Alive {
			continue
		}
		dx := e.X - playerX
		dy := e.Y - playerY
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist < minDist {
			minDist = dist
			nearest = e
		}
	}

	if nearest == nil {
		return
	}

	w.Timer = w.Cooldown
	dx := nearest.X - playerX
	dy := nearest.Y - playerY
	dist := math.Sqrt(dx*dx + dy*dy)
	vx := (dx / dist) * w.BulletSpeed
	vy := (dy / dist) * w.BulletSpeed

	dmg := w.Damage + (level-1)/2

	*bullets = append(*bullets, &Bullet{
		X:      playerX,
		Y:      playerY,
		VX:     vx,
		VY:     vy,
		Radius: 3,
		Damage: dmg,
		Alive:  true,
	})
}

type Bullet struct {
	X, Y   float64
	VX, VY float64
	Radius float64
	Damage int
	Alive  bool
}

func (b *Bullet) Update(camX, camY float64) {
	b.X += b.VX
	b.Y += b.VY

	sx := b.X - camX
	sy := b.Y - camY
	margin := 100.0
	if sx < -margin || sx > float64(screenWidth)+margin || sy < -margin || sy > float64(screenHeight)+margin {
		b.Alive = false
	}
}

func (b *Bullet) Draw(screen *ebiten.Image, cam *Camera) {
	sx, sy := cam.WorldToScreen(b.X, b.Y)
	vector.FillCircle(screen, sx, sy, float32(b.Radius), color.RGBA{R: 255, G: 255, B: 50, A: 255}, true)
}
