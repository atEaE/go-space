package system

import (
	"math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/event"
)

var queryCollisionBullet = donburi.NewQuery(
	filter.Contains(
		component.BulletTag,
		component.Position,
		component.CircleCollider,
		component.Damage,
	),
)

var queryCollisionEnemy = donburi.NewQuery(
	filter.Contains(
		component.EnemyTag,
		component.Position,
		component.CircleCollider,
		component.Health,
	),
)

var queryCollisionPlayer = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Position,
		component.CircleCollider,
		component.Health,
	),
)

var queryCollisionGem = donburi.NewQuery(
	filter.Contains(
		component.GemTag,
		component.Position,
		component.EXPValue,
	),
)

func circleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	rSum := r1 + r2
	return dx*dx+dy*dy < rSum*rSum
}

func UpdateCollision(e *ecs.ECS) {
	gs := component.GameState.Get(component.GameState.MustFirst(e.World))

	// 弾 ↔ 敵
	var bulletsToRemove []donburi.Entity
	var enemiesToRemove []donburi.Entity

	queryCollisionBullet.Each(e.World, func(bullet *donburi.Entry) {
		bPos := component.Position.GetValue(bullet)
		bCol := component.CircleCollider.GetValue(bullet)
		bDmg := component.Damage.GetValue(bullet)

		queryCollisionEnemy.Each(e.World, func(enemy *donburi.Entry) {
			ePos := component.Position.GetValue(enemy)
			eCol := component.CircleCollider.GetValue(enemy)

			if circleCollision(bPos.X, bPos.Y, bCol.Radius, ePos.X, ePos.Y, eCol.Radius) {
				hp := component.Health.Get(enemy)
				hp.HP -= bDmg.Value
				bulletsToRemove = append(bulletsToRemove, bullet.Entity())

				if hp.HP <= 0 {
					exp := component.EXPValue.GetValue(enemy)
					enemiesToRemove = append(enemiesToRemove, enemy.Entity())
					event.EnemyDeath.Publish(e.World, event.EnemyDeathEvent{
						X: ePos.X, Y: ePos.Y, EXPAmount: exp.Amount,
					})
				}
			}
		})
	})

	for _, entity := range bulletsToRemove {
		if e.World.Valid(entity) {
			e.World.Remove(entity)
		}
	}
	for _, entity := range enemiesToRemove {
		if e.World.Valid(entity) {
			e.World.Remove(entity)
		}
	}

	// 敵 ↔ プレイヤー
	var enemyContactRemove []donburi.Entity

	queryCollisionPlayer.Each(e.World, func(player *donburi.Entry) {
		pPos := component.Position.GetValue(player)
		pCol := component.CircleCollider.GetValue(player)
		pHP := component.Health.Get(player)

		queryCollisionEnemy.Each(e.World, func(enemy *donburi.Entry) {
			ePos := component.Position.GetValue(enemy)
			eCol := component.CircleCollider.GetValue(enemy)
			eDmg := component.Damage.GetValue(enemy)

			if circleCollision(pPos.X, pPos.Y, pCol.Radius, ePos.X, ePos.Y, eCol.Radius) {
				pHP.HP -= eDmg.Value
				if pHP.HP < 0 {
					pHP.HP = 0
				}
				enemyContactRemove = append(enemyContactRemove, enemy.Entity())
				if pHP.HP <= 0 {
					gs.GameOver = true
				}
			}
		})
	})

	for _, entity := range enemyContactRemove {
		if e.World.Valid(entity) {
			e.World.Remove(entity)
		}
	}

	// プレイヤー ↔ ジェム
	var gemsToRemove []donburi.Entity
	pickupRange := 30.0

	queryCollisionPlayer.Each(e.World, func(player *donburi.Entry) {
		pPos := component.Position.GetValue(player)

		queryCollisionGem.Each(e.World, func(gem *donburi.Entry) {
			gPos := component.Position.GetValue(gem)
			gExp := component.EXPValue.GetValue(gem)

			dx := gPos.X - pPos.X
			dy := gPos.Y - pPos.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < pickupRange {
				gemsToRemove = append(gemsToRemove, gem.Entity())
				event.GemPickup.Publish(e.World, event.GemPickupEvent{
					PlayerEntry: player,
					Amount:      gExp.Amount,
				})
			}
		})
	})

	for _, entity := range gemsToRemove {
		if e.World.Valid(entity) {
			e.World.Remove(entity)
		}
	}
}
