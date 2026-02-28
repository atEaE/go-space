package system

import (
	"math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
)

var queryEnemyAIPlayer = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Position,
	),
)

var queryEnemy = donburi.NewQuery(
	filter.Contains(
		component.EnemyTag,
		component.Position,
		component.Velocity,
		component.Speed,
	),
)

func UpdateEnemyAI(e *ecs.ECS) {
	var playerX, playerY float64

	queryEnemyAIPlayer.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		playerX = pos.X
		playerY = pos.Y
	})

	queryEnemy.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		vel := component.Velocity.Get(entry)
		spd := component.Speed.GetValue(entry).Value

		dx := playerX - pos.X
		dy := playerY - pos.Y

		if length := math.Sqrt(dx*dx + dy*dy); length > 0 {
			dx /= length
			dy /= length
		}

		vel.X = dx * spd
		vel.Y = dy * spd
	})
}
