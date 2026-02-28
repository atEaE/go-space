package system

import (
	"github.com/yohamta/donburi/ecs"

	"github.com/atEaE/go-space/internal/component"
)

// UpdateGameState increments the game tick each frame.
// If the game is over, it returns early without updating.
func UpdateGameState(e *ecs.ECS) {
	gs := component.GameState.Get(component.GameState.MustFirst(e.World))
	if gs.GameOver {
		return
	}
	gs.Tick++
}
