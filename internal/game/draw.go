package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	drawGrid(screen, g)

	// ジェム描画
	gemColor := color.RGBA{R: 50, G: 220, B: 80, A: 255}
	for _, gem := range g.EXPGems {
		sx, sy := g.Camera.WorldToScreen(gem.Pos.X, gem.Pos.Y)
		vector.FillCircle(screen, sx, sy, float32(gem.Radius), gemColor, true)
	}

	// 敵描画
	enemyColor := color.RGBA{R: 220, G: 40, B: 40, A: 255}
	for _, e := range g.Enemies {
		sx, sy := g.Camera.WorldToScreen(e.Pos.X, e.Pos.Y)
		vector.FillCircle(screen, sx, sy, float32(e.Radius), enemyColor, true)
	}

	// 弾描画
	bulletColor := color.RGBA{R: 255, G: 255, B: 50, A: 255}
	for _, b := range g.Bullets {
		if b.Alive {
			sx, sy := g.Camera.WorldToScreen(b.Pos.X, b.Pos.Y)
			vector.FillCircle(screen, sx, sy, float32(b.Radius), bulletColor, true)
		}
	}

	// プレイヤー描画
	px, py := g.Camera.WorldToScreen(g.Player.Pos.X, g.Player.Pos.Y)
	vector.FillCircle(screen, px, py, float32(g.Player.Radius), color.White, true)

	drawHUD(screen, g)
}

func drawGrid(screen *ebiten.Image, g *Game) {
	gridSize := 50.0
	gridColor := color.RGBA{R: 50, G: 50, B: 50, A: 255}
	startX := math.Floor(g.Camera.X/gridSize) * gridSize
	startY := math.Floor(g.Camera.Y/gridSize) * gridSize
	for x := startX; x < g.Camera.X+float64(ScreenWidth)+gridSize; x += gridSize {
		sx, _ := g.Camera.WorldToScreen(x, 0)
		vector.StrokeLine(screen, sx, 0, sx, float32(ScreenHeight), 1, gridColor, true)
	}
	for y := startY; y < g.Camera.Y+float64(ScreenHeight)+gridSize; y += gridSize {
		_, sy := g.Camera.WorldToScreen(0, y)
		vector.StrokeLine(screen, 0, sy, float32(ScreenWidth), sy, 1, gridColor, true)
	}
}

func drawHUD(screen *ebiten.Image, g *Game) {
	hpBar := fmt.Sprintf("HP: %d/%d", g.Player.HP, g.Player.MaxHP)
	lvl := fmt.Sprintf("Lv: %d  EXP: %d/%d", g.Player.Level, g.Player.EXP, g.Player.NextEXP)
	ebitenutil.DebugPrintAt(screen, hpBar, 8, 8)
	ebitenutil.DebugPrintAt(screen, lvl, 8, 24)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), 8, 40)

	if g.GameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER - Press R to Restart", ScreenWidth/2-100, ScreenHeight/2)
	}
}
