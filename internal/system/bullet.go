package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/config"
)

var queryBullet = donburi.NewQuery(
	filter.Contains(
		component.BulletTag,
		component.Position,
	),
)

func UpdateBulletLifetime(e *ecs.ECS) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))

	var toRemove []donburi.Entity
	queryBullet.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		sx := pos.X - cam.X
		sy := pos.Y - cam.Y
		margin := 100.0
		if sx < -margin || sx > float64(config.ScreenWidth)+margin ||
			sy < -margin || sy > float64(config.ScreenHeight)+margin {
			toRemove = append(toRemove, entry.Entity())
		}
	})

	for _, entity := range toRemove {
		e.World.Remove(entity)
	}
}
