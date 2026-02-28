package system

import (
	"math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/archetype"
	"github.com/atEaE/go-space/internal/component"
)

var queryWeaponPlayer = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Position,
		component.Weapon,
		component.PlayerStats,
	),
)

var queryWeaponEnemy = donburi.NewQuery(
	filter.Contains(
		component.EnemyTag,
		component.Position,
	),
)

func UpdateWeapon(e *ecs.ECS) {
	queryWeaponPlayer.Each(e.World, func(player *donburi.Entry) {
		wpn := component.Weapon.Get(player)
		wpn.Timer--
		if wpn.Timer > 0 {
			return
		}

		playerPos := component.Position.GetValue(player)
		stats := component.PlayerStats.GetValue(player)

		var nearest *donburi.Entry
		minDist := math.MaxFloat64

		queryWeaponEnemy.Each(e.World, func(enemy *donburi.Entry) {
			ePos := component.Position.GetValue(enemy)
			dx := ePos.X - playerPos.X
			dy := ePos.Y - playerPos.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
				nearest = enemy
			}
		})

		if nearest == nil {
			return
		}

		wpn.Timer = wpn.Cooldown
		ePos := component.Position.GetValue(nearest)
		dx := ePos.X - playerPos.X
		dy := ePos.Y - playerPos.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		vx := (dx / dist) * wpn.BulletSpeed
		vy := (dy / dist) * wpn.BulletSpeed
		dmg := wpn.BaseDamage + (stats.Level-1)/2

		archetype.CreateBullet(e.World, playerPos.X, playerPos.Y, vx, vy, dmg)
	})
}
