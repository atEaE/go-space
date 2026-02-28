package system

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/atEaE/go-space/internal/component"
	"github.com/atEaE/go-space/internal/config"
)

var queryRenderEnemy = donburi.NewQuery(
	filter.Contains(component.EnemyTag, component.Position, component.CircleCollider),
)

var queryRenderBullet = donburi.NewQuery(
	filter.Contains(component.BulletTag, component.Position, component.CircleCollider),
)

var queryRenderGem = donburi.NewQuery(
	filter.Contains(component.GemTag, component.Position, component.CircleCollider),
)

var queryRenderPlayer = donburi.NewQuery(
	filter.Contains(component.PlayerTag, component.Position, component.CircleCollider),
)

func DrawBackground(e *ecs.ECS, screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	cam := component.Camera.Get(component.Camera.MustFirst(e.World))
	gridSize := 50.0
	gridColor := color.RGBA{R: 50, G: 50, B: 50, A: 255}

	startX := math.Floor(cam.X/gridSize) * gridSize
	startY := math.Floor(cam.Y/gridSize) * gridSize

	for x := startX; x < cam.X+float64(config.ScreenWidth)+gridSize; x += gridSize {
		sx := float32(x - cam.X)
		vector.StrokeLine(screen, sx, 0, sx, float32(config.ScreenHeight), 1, gridColor, true)
	}
	for y := startY; y < cam.Y+float64(config.ScreenHeight)+gridSize; y += gridSize {
		sy := float32(y - cam.Y)
		vector.StrokeLine(screen, 0, sy, float32(config.ScreenWidth), sy, 1, gridColor, true)
	}
}

func DrawGems(e *ecs.ECS, screen *ebiten.Image) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))
	gemColor := color.RGBA{R: 50, G: 220, B: 80, A: 255}

	queryRenderGem.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		col := component.CircleCollider.GetValue(entry)
		sx := float32(pos.X - cam.X)
		sy := float32(pos.Y - cam.Y)
		vector.FillCircle(screen, sx, sy, float32(col.Radius), gemColor, true)
	})
}

func DrawEnemies(e *ecs.ECS, screen *ebiten.Image) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))
	enemyColor := color.RGBA{R: 220, G: 40, B: 40, A: 255}

	queryRenderEnemy.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		col := component.CircleCollider.GetValue(entry)
		sx := float32(pos.X - cam.X)
		sy := float32(pos.Y - cam.Y)
		vector.FillCircle(screen, sx, sy, float32(col.Radius), enemyColor, true)
	})
}

func DrawBullets(e *ecs.ECS, screen *ebiten.Image) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))
	bulletColor := color.RGBA{R: 255, G: 255, B: 50, A: 255}

	queryRenderBullet.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		col := component.CircleCollider.GetValue(entry)
		sx := float32(pos.X - cam.X)
		sy := float32(pos.Y - cam.Y)
		vector.FillCircle(screen, sx, sy, float32(col.Radius), bulletColor, true)
	})
}

func DrawPlayer(e *ecs.ECS, screen *ebiten.Image) {
	cam := component.Camera.Get(component.Camera.MustFirst(e.World))

	queryRenderPlayer.Each(e.World, func(entry *donburi.Entry) {
		pos := component.Position.GetValue(entry)
		col := component.CircleCollider.GetValue(entry)
		sx := float32(pos.X - cam.X)
		sy := float32(pos.Y - cam.Y)
		vector.FillCircle(screen, sx, sy, float32(col.Radius), color.White, true)
	})
}

func DrawHUD(e *ecs.ECS, screen *ebiten.Image) {
	gs := component.GameState.GetValue(component.GameState.MustFirst(e.World))

	queryRenderPlayer.Each(e.World, func(entry *donburi.Entry) {
		hp := component.Health.GetValue(entry)
		stats := component.PlayerStats.GetValue(entry)

		hpBar := fmt.Sprintf("HP: %d/%d", hp.HP, hp.MaxHP)
		lvl := fmt.Sprintf("Lv: %d  EXP: %d/%d", stats.Level, stats.EXP, stats.NextEXP)
		ebitenutil.DebugPrintAt(screen, hpBar, 8, 8)
		ebitenutil.DebugPrintAt(screen, lvl, 8, 24)
	})

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), 8, 40)

	if gs.GameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER - Press R to Return to Title",
			config.ScreenWidth/2-120, config.ScreenHeight/2)
	}
}
