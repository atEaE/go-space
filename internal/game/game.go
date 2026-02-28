package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"

	"github.com/atEaE/go-space/internal/archetype"
	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/config"
	"github.com/atEaE/go-space/internal/layer"
	"github.com/atEaE/go-space/internal/system"
)

type Game struct {
	ecs *ecs.ECS
}

func New() *Game {
	g := &Game{}
	g.ecs = g.setupECS()
	return g
}

func (g *Game) setupECS() *ecs.ECS {
	w := donburi.NewWorld()
	e := ecs.NewECS(w)

	// Update systems (順序が重要)
	e.AddSystem(system.UpdateGameState)
	e.AddSystem(system.UpdateInput)
	e.AddSystem(system.UpdateCamera)
	e.AddSystem(system.UpdateSpawner)
	e.AddSystem(system.UpdateEnemyAI)
	e.AddSystem(system.UpdateWeapon)
	e.AddSystem(system.UpdateMovement)
	e.AddSystem(system.UpdateBulletLifetime)
	e.AddSystem(system.UpdateCollision)
	e.AddSystem(system.ProcessEvents)

	// Renderers (レイヤー順)
	e.AddRenderer(layer.Background, system.DrawBackground)
	e.AddRenderer(layer.Gems, system.DrawGems)
	e.AddRenderer(layer.Enemies, system.DrawEnemies)
	e.AddRenderer(layer.Bullets, system.DrawBullets)
	e.AddRenderer(layer.Player, system.DrawPlayer)
	e.AddRenderer(layer.HUD, system.DrawHUD)

	// イベント購読セットアップ
	system.SetupEvents(w)

	// 初期エンティティ生成
	archetype.CreatePlayer(w)
	archetype.CreateGameWorld(w)

	return e
}

func (g *Game) Update() error {
	gs := component.GameState.Get(component.GameState.MustFirst(g.ecs.World))
	if gs.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.ecs = g.setupECS()
		}
		return nil
	}

	g.ecs.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ecs.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
