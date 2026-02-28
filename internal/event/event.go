package event

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
)

type EnemyDeathEvent struct {
	X, Y      float64
	EXPAmount int
}

type GemPickupEvent struct {
	PlayerEntry *donburi.Entry
	Amount      int
}

type LevelUpEvent struct {
	Level int
}

var (
	EnemyDeath = events.NewEventType[EnemyDeathEvent]()
	GemPickup  = events.NewEventType[GemPickupEvent]()
	LevelUp    = events.NewEventType[LevelUpEvent]()
)
