package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
)

var queryMovement = donburi.NewQuery(
	filter.Contains(
		component.Position,
		component.Velocity,
	),
)

// UpdateMovement applies velocity to position for all entities
// that have both Position and Velocity components.
func UpdateMovement(e *ecs.ECS) {
	queryMovement.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.Get(entry)
		vel := component.Velocity.GetValue(entry)
		pos.X += vel.X
		pos.Y += vel.Y
	})
}
