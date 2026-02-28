package system

import (
	"math"

	"github.com/atEaE/go-space/internal/entity"
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
		BulletSpeed: 5.0,
		Damage:      1,
	}
}

func (w *Weapon) Update(playerX, playerY float64, enemies []*entity.Enemy, bullets *[]*entity.Bullet, level int) {
	w.Timer--
	if w.Timer > 0 {
		return
	}

	var nearest *entity.Enemy
	minDist := math.MaxFloat64
	for _, e := range enemies {
		if !e.Alive {
			continue
		}
		dx := e.Pos.X - playerX
		dy := e.Pos.Y - playerY
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
	dx := nearest.Pos.X - playerX
	dy := nearest.Pos.Y - playerY
	dist := math.Sqrt(dx*dx + dy*dy)
	vx := (dx / dist) * w.BulletSpeed
	vy := (dy / dist) * w.BulletSpeed

	dmg := w.Damage + (level-1)/2

	*bullets = append(*bullets, &entity.Bullet{
		Pos:    entity.Position{X: playerX, Y: playerY},
		VX:     vx,
		VY:     vy,
		Radius: 3,
		Damage: dmg,
		Alive:  true,
	})
}
