package component

import "github.com/yohamta/donburi"

type PositionData struct {
	X, Y float64
}

type VelocityData struct {
	X, Y float64
}

type HealthData struct {
	HP, MaxHP int
}

type CircleColliderData struct {
	Radius float64
}

type SpeedData struct {
	Value float64
}

type PlayerStatsData struct {
	Level   int
	EXP     int
	NextEXP int
}

type DamageData struct {
	Value int
}

type EXPValueData struct {
	Amount int
}

type WeaponData struct {
	Cooldown    int
	Timer       int
	BulletSpeed float64
	BaseDamage  int
}

type CameraData struct {
	X, Y float64
}

type SpawnerData struct {
	Timer int
	Rate  int
}

type GameStateData struct {
	GameOver bool
	Tick     int
}

var (
	Position       = donburi.NewComponentType[PositionData]()
	Velocity       = donburi.NewComponentType[VelocityData]()
	Health         = donburi.NewComponentType[HealthData]()
	CircleCollider = donburi.NewComponentType[CircleColliderData]()
	Speed          = donburi.NewComponentType[SpeedData]()
	PlayerStats    = donburi.NewComponentType[PlayerStatsData]()
	Damage         = donburi.NewComponentType[DamageData]()
	EXPValue       = donburi.NewComponentType[EXPValueData]()
	Weapon         = donburi.NewComponentType[WeaponData]()
	Camera         = donburi.NewComponentType[CameraData]()
	Spawner        = donburi.NewComponentType[SpawnerData]()
	GameState      = donburi.NewComponentType[GameStateData]()
)
