package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"

	"github.com/atEaE/go-space/internal/archetype"
	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/config"
	"github.com/atEaE/go-space/internal/layer"
	"github.com/atEaE/go-space/internal/system"
)

type Scene int

const (
	SceneTitle Scene = iota
	ScenePlaying
)

type Game struct {
	scene Scene
	ecs   *ecs.ECS
}

func New() *Game {
	return &Game{scene: SceneTitle}
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
	switch g.scene {
	case SceneTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.ecs = g.setupECS()
			g.scene = ScenePlaying
		}
	case ScenePlaying:
		gs := component.GameState.Get(component.GameState.MustFirst(g.ecs.World))
		if gs.GameOver {
			if inpututil.IsKeyJustPressed(ebiten.KeyR) {
				g.ecs = nil
				g.scene = SceneTitle
			}
			return nil
		}
		g.ecs.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.scene {
	case SceneTitle:
		screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})
		ebitenutil.DebugPrintAt(screen, "Vampire Survivors Mini",
			config.ScreenWidth/2-70, config.ScreenHeight/3)
		ebitenutil.DebugPrintAt(screen, "Press Space to Start",
			config.ScreenWidth/2-65, config.ScreenHeight*2/3)
	case ScenePlaying:
		g.ecs.Draw(screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}
