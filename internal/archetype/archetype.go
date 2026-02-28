package archetype

import (
	"math"
	"math/rand/v2"

	"github.com/yohamta/donburi"

	"github.com/atEaE/go-space/internal/component"
)

// CreatePlayer : プレイヤーエンティティを生成する (HP:100, Speed:3.0, Weapon cooldown:20)。
func CreatePlayer(w donburi.World) *donburi.Entry {
	entity := w.Create(
		component.PlayerTag,
		component.Position,
		component.Velocity,
		component.Health,
		component.CircleCollider,
		component.Speed,
		component.PlayerStats,
		component.Weapon,
	)
	entry := w.Entry(entity)
	component.Health.SetValue(entry, component.HealthData{HP: 100, MaxHP: 100})
	component.CircleCollider.SetValue(entry, component.CircleColliderData{Radius: 8})
	component.Speed.SetValue(entry, component.SpeedData{Value: 3.0})
	component.PlayerStats.SetValue(entry, component.PlayerStatsData{Level: 1, NextEXP: 10})
	component.Weapon.SetValue(entry, component.WeaponData{
		Cooldown:    20,
		BulletSpeed: 5.0,
		BaseDamage:  1,
	})
	return entry
}

// CreateEnemy : プレイヤーから300〜400距離のランダムな位置に敵を生成する。
func CreateEnemy(w donburi.World, playerX, playerY float64) *donburi.Entry {
	angle := rand.Float64() * 2 * math.Pi
	dist := 300.0 + rand.Float64()*100.0

	entity := w.Create(
		component.EnemyTag,
		component.Position,
		component.Velocity,
		component.Health,
		component.CircleCollider,
		component.Speed,
		component.Damage,
		component.EXPValue,
	)
	entry := w.Entry(entity)
	component.Position.SetValue(entry, component.PositionData{
		X: playerX + math.Cos(angle)*dist,
		Y: playerY + math.Sin(angle)*dist,
	})
	component.Health.SetValue(entry, component.HealthData{HP: 3, MaxHP: 3})
	component.CircleCollider.SetValue(entry, component.CircleColliderData{Radius: 6})
	component.Speed.SetValue(entry, component.SpeedData{Value: 1.0 + rand.Float64()*0.5})
	component.Damage.SetValue(entry, component.DamageData{Value: 10})
	component.EXPValue.SetValue(entry, component.EXPValueData{Amount: 1})
	return entry
}

// CreateBullet : 指定座標(x,y)に、指定された速度とダメージを持つ弾を生成する。
func CreateBullet(w donburi.World, x, y, vx, vy float64, damage int) *donburi.Entry {
	entity := w.Create(
		component.BulletTag,
		component.Position,
		component.Velocity,
		component.CircleCollider,
		component.Damage,
	)
	entry := w.Entry(entity)
	component.Position.SetValue(entry, component.PositionData{X: x, Y: y})
	component.Velocity.SetValue(entry, component.VelocityData{X: vx, Y: vy})
	component.CircleCollider.SetValue(entry, component.CircleColliderData{Radius: 3})
	component.Damage.SetValue(entry, component.DamageData{Value: damage})
	return entry
}

// CreateGem : 指定座標(x,y)に経験値ジェムを生成する。
func CreateGem(w donburi.World, x, y float64, amount int) *donburi.Entry {
	entity := w.Create(
		component.GemTag,
		component.Position,
		component.CircleCollider,
		component.EXPValue,
	)
	entry := w.Entry(entity)
	component.Position.SetValue(entry, component.PositionData{X: x, Y: y})
	component.CircleCollider.SetValue(entry, component.CircleColliderData{Radius: 4})
	component.EXPValue.SetValue(entry, component.EXPValueData{Amount: amount})
	return entry
}

// CreateGameWorld : Camera、Spawner、GameStateのシングルトンエンティティを生成する。
func CreateGameWorld(w donburi.World) {
	// Camera singleton
	camEntity := w.Create(component.Camera)
	_ = w.Entry(camEntity)

	// Spawner singleton
	spawnEntity := w.Create(component.Spawner)
	spawnEntry := w.Entry(spawnEntity)
	component.Spawner.SetValue(spawnEntry, component.SpawnerData{Rate: 60})

	// GameState singleton
	gsEntity := w.Create(component.GameState)
	_ = w.Entry(gsEntity)
}
