package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"

	"github.com/atEaE/go-space/internal/archetype"
	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/event"
)

func SetupEvents(w donburi.World) {
	event.EnemyDeath.Subscribe(w, handleEnemyDeath)
	event.GemPickup.Subscribe(w, handleGemPickup)
}

func handleEnemyDeath(w donburi.World, ev event.EnemyDeathEvent) {
	archetype.CreateGem(w, ev.X, ev.Y, ev.EXPAmount)
}

func handleGemPickup(w donburi.World, ev event.GemPickupEvent) {
	if !ev.PlayerEntry.Valid() {
		return
	}
	stats := component.PlayerStats.Get(ev.PlayerEntry)
	stats.EXP += ev.Amount
	for stats.EXP >= stats.NextEXP {
		stats.EXP -= stats.NextEXP
		stats.Level++
		stats.NextEXP = stats.Level * 10

		spd := component.Speed.Get(ev.PlayerEntry)
		spd.Value += 0.2

		hp := component.Health.Get(ev.PlayerEntry)
		hp.MaxHP += 10
		hp.HP = hp.MaxHP

		event.LevelUp.Publish(w, event.LevelUpEvent{Level: stats.Level})
	}
}

func ProcessEvents(e *ecs.ECS) {
	events.ProcessAllEvents(e.World)
}
