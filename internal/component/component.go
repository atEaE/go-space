package component

import "github.com/yohamta/donburi"

// PositionData : ワールド座標を表す。
type PositionData struct {
	X, Y float64
}

// VelocityData : フレームあたりの移動量。
type VelocityData struct {
	X, Y float64
}

// HealthData : 現在HPと最大HP。
type HealthData struct {
	HP, MaxHP int
}

// CircleColliderData : 円形衝突判定の半径。
type CircleColliderData struct {
	Radius float64
}

// SpeedData : エンティティの移動速度。
type SpeedData struct {
	Value float64
}

// PlayerStatsData : プレイヤーのレベル・経験値情報。
type PlayerStatsData struct {
	Level   int
	EXP     int
	NextEXP int
}

// DamageData : 接触時に与えるダメージ量。
type DamageData struct {
	Value int
}

// EXPValueData : 取得時に得られる経験値量。
type EXPValueData struct {
	Amount int
}

// WeaponData : 武器のクールダウン・弾速・ダメージ。
type WeaponData struct {
	Cooldown    int
	Timer       int
	BulletSpeed float64
	BaseDamage  int
}

// CameraData : カメラのワールド座標(左上)。
type CameraData struct {
	X, Y float64
}

// SpawnerData : 敵スポーンのタイマーと間隔。
type SpawnerData struct {
	Timer int
	Rate  int
}

// GameStateData : ゲーム全体の状態(ゲームオーバー判定、経過ティック)。
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
