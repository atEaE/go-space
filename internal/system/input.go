package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
)

var queryPlayer = donburi.NewQuery(
	filter.Contains(
		component.PlayerTag,
		component.Velocity,
		component.Speed,
	),
)

// UpdateInput reads WASD / arrow key input and sets the player's velocity
// based on the direction and speed. Diagonal movement is normalized.
func UpdateInput(e *ecs.ECS) {
	queryPlayer.Each(e.World, func(entry *donburi.Entry) {
		var dx, dy float64

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			dy--
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			dy++
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			dx--
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			dx++
		}

		// Normalize diagonal movement.
		if length := math.Sqrt(dx*dx + dy*dy); length > 0 {
			dx /= length
			dy /= length
		}

		spd := component.Speed.GetValue(entry).Value
		vel := component.Velocity.Get(entry)
		vel.X = dx * spd
		vel.Y = dy * spd
	})
}
