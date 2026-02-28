package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/archetype"
	"github.com/atEaE/go-space/internal/component"
)

var querySpawnerPlayer = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Position,
	),
)

func UpdateSpawner(e *ecs.ECS) {
	spawner := component.Spawner.Get(component.Spawner.MustFirst(e.World))
	gs := component.GameState.GetValue(component.GameState.MustFirst(e.World))

	spawner.Timer++

	currentRate := max(spawner.Rate-gs.Tick/600, 15)

	if spawner.Timer >= currentRate {
		spawner.Timer = 0

		querySpawnerPlayer.Each(e.World, func(entry *donburi.Entry) {
			pos := component.Position.GetValue(entry)
			archetype.CreateEnemy(e.World, pos.X, pos.Y)
		})
	}
}
