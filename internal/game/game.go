package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/atEaE/go-space/internal/entity"
	"github.com/atEaE/go-space/internal/system"
)

type Game struct {
	Player   *entity.Player
	Enemies  []*entity.Enemy
	Bullets  []*entity.Bullet
	EXPGems  []*entity.EXPGem
	Weapon   *system.Weapon
	Camera   *system.Camera
	Spawner  *system.EnemySpawner
	GameOver bool
	Tick     int
}

func New() *Game {
	return &Game{
		Player:  entity.NewPlayer(),
		Weapon:  system.NewWeapon(),
		Camera:  &system.Camera{},
		Spawner: system.NewEnemySpawner(),
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *New()
		}
		return nil
	}

	g.Tick++

	// プレイヤー移動
	g.Player.Update()

	// カメラ更新
	g.Camera.Update(g.Player.Pos.X, g.Player.Pos.Y, ScreenWidth, ScreenHeight)

	// 敵スポーン
	if enemy := g.Spawner.Update(g.Tick, g.Player.Pos.X, g.Player.Pos.Y); enemy != nil {
		g.Enemies = append(g.Enemies, enemy)
	}

	// 敵更新
	for _, e := range g.Enemies {
		if e.Alive {
			e.Update(g.Player.Pos.X, g.Player.Pos.Y)
		}
	}

	// 武器＆弾更新
	g.Weapon.Update(g.Player.Pos.X, g.Player.Pos.Y, g.Enemies, &g.Bullets, g.Player.Level)
	for _, b := range g.Bullets {
		if b.Alive {
			b.Update()
			if system.IsOffScreen(b.Pos.X, b.Pos.Y, g.Camera.X, g.Camera.Y, ScreenWidth, ScreenHeight) {
				b.Alive = false
			}
		}
	}

	// 衝突判定: 弾↔敵
	for _, b := range g.Bullets {
		if !b.Alive {
			continue
		}
		for _, e := range g.Enemies {
			if !e.Alive {
				continue
			}
			if system.CircleCollision(b.Pos.X, b.Pos.Y, b.Radius, e.Pos.X, e.Pos.Y, e.Radius) {
				e.HP -= b.Damage
				b.Alive = false
				if e.HP <= 0 {
					e.Alive = false
				}
				break
			}
		}
	}

	// 衝突判定: 敵↔プレイヤー
	for _, e := range g.Enemies {
		if !e.Alive {
			continue
		}
		if system.CircleCollision(e.Pos.X, e.Pos.Y, e.Radius, g.Player.Pos.X, g.Player.Pos.Y, g.Player.Radius) {
			g.Player.TakeDamage(e.Damage)
			e.Alive = false
			if g.Player.HP <= 0 {
				g.GameOver = true
			}
		}
	}

	// 死んだ敵からジェム生成
	for _, e := range g.Enemies {
		if !e.Alive && e.HP <= 0 {
			g.EXPGems = append(g.EXPGems, &entity.EXPGem{
				Pos:    entity.Position{X: e.Pos.X, Y: e.Pos.Y},
				Amount: 1,
				Radius: 4,
			})
		}
	}

	// 衝突判定: プレイヤー↔ジェム
	pickupRange := 30.0
	remaining := g.EXPGems[:0]
	for _, gem := range g.EXPGems {
		dx := gem.Pos.X - g.Player.Pos.X
		dy := gem.Pos.Y - g.Player.Pos.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist < pickupRange {
			g.Player.AddEXP(gem.Amount)
		} else {
			remaining = append(remaining, gem)
		}
	}
	g.EXPGems = remaining

	// 死んだオブジェクト除去
	g.Enemies = filterAliveEnemies(g.Enemies)
	g.Bullets = filterAliveBullets(g.Bullets)

	return nil
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func filterAliveEnemies(enemies []*entity.Enemy) []*entity.Enemy {
	result := enemies[:0]
	for _, e := range enemies {
		if e.Alive {
			result = append(result, e)
		}
	}
	return result
}

func filterAliveBullets(bullets []*entity.Bullet) []*entity.Bullet {
	result := bullets[:0]
	for _, b := range bullets {
		if b.Alive {
			result = append(result, b)
		}
	}
	return result
}
