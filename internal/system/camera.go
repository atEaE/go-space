package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/config"
)

var queryCamera = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Position,
	),
)

func UpdateCamera(e *ecs.ECS) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))

	queryCamera.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		cam.X = pos.X - float64(config.ScreenWidth)/2
		cam.Y = pos.Y - float64(config.ScreenHeight)/2
	})
}
