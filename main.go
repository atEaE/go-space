package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type EXPGem struct {
	X, Y   float64
	Amount int
	Radius float64
}

type Game struct {
	Player     *Player
	Enemies    []*Enemy
	Bullets    []*Bullet
	EXPGems    []*EXPGem
	Weapon     *Weapon
	Camera     *Camera
	SpawnTimer int
	SpawnRate  int
	GameOver   bool
	Tick       int
}

func NewGame() *Game {
	return &Game{
		Player:    NewPlayer(),
		Weapon:    NewWeapon(),
		Camera:    &Camera{},
		SpawnRate: 60,
	}
}

func (g *Game) Update() error {
	if g.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return nil
	}

	g.Tick++

	// プレイヤー移動
	g.Player.Update()

	// カメラ更新
	g.Camera.Update(g.Player.X, g.Player.Y)

	// 敵スポーン（時間経過で頻度増加）
	g.SpawnTimer++
	currentRate := max(g.SpawnRate-g.Tick/600, 15)
	if g.SpawnTimer >= currentRate {
		g.SpawnTimer = 0
		g.Enemies = append(g.Enemies, NewEnemy(g.Player.X, g.Player.Y))
	}

	// 敵更新
	for _, e := range g.Enemies {
		if e.Alive {
			e.Update(g.Player.X, g.Player.Y)
		}
	}

	// 武器＆弾更新
	g.Weapon.Update(g.Player.X, g.Player.Y, g.Enemies, &g.Bullets, g.Player.Level)
	for _, b := range g.Bullets {
		if b.Alive {
			b.Update(g.Camera.X, g.Camera.Y)
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
			if circleCollision(b.X, b.Y, b.Radius, e.X, e.Y, e.Radius) {
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
		if circleCollision(e.X, e.Y, e.Radius, g.Player.X, g.Player.Y, g.Player.Radius) {
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
			g.EXPGems = append(g.EXPGems, &EXPGem{
				X:      e.X,
				Y:      e.Y,
				Amount: 1,
				Radius: 4,
			})
		}
	}

	// 衝突判定: プレイヤー↔ジェム
	pickupRange := 30.0
	remaining := g.EXPGems[:0]
	for _, gem := range g.EXPGems {
		dx := gem.X - g.Player.X
		dy := gem.Y - g.Player.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist < pickupRange {
			g.Player.AddEXP(gem.Amount)
		} else {
			remaining = append(remaining, gem)
		}
	}
	g.EXPGems = remaining

	// 死んだオブジェクト除去
	g.Enemies = filterAlive(g.Enemies)
	g.Bullets = filterAliveBullets(g.Bullets)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 背景
	screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	// グリッド線
	gridSize := 50.0
	gridColor := color.RGBA{R: 50, G: 50, B: 50, A: 255}
	startX := math.Floor(g.Camera.X/gridSize) * gridSize
	startY := math.Floor(g.Camera.Y/gridSize) * gridSize
	for x := startX; x < g.Camera.X+float64(screenWidth)+gridSize; x += gridSize {
		sx, _ := g.Camera.WorldToScreen(x, 0)
		vector.StrokeLine(screen, sx, 0, sx, float32(screenHeight), 1, gridColor, true)
	}
	for y := startY; y < g.Camera.Y+float64(screenHeight)+gridSize; y += gridSize {
		_, sy := g.Camera.WorldToScreen(0, y)
		vector.StrokeLine(screen, 0, sy, float32(screenWidth), sy, 1, gridColor, true)
	}

	// ジェム描画
	gemColor := color.RGBA{R: 50, G: 220, B: 80, A: 255}
	for _, gem := range g.EXPGems {
		sx, sy := g.Camera.WorldToScreen(gem.X, gem.Y)
		vector.FillCircle(screen, sx, sy, float32(gem.Radius), gemColor, true)
	}

	// 敵描画
	for _, e := range g.Enemies {
		e.Draw(screen, g.Camera)
	}

	// 弾描画
	for _, b := range g.Bullets {
		if b.Alive {
			b.Draw(screen, g.Camera)
		}
	}

	// プレイヤー描画
	g.Player.Draw(screen, g.Camera)

	// UI
	hpBar := fmt.Sprintf("HP: %d/%d", g.Player.HP, g.Player.MaxHP)
	lvl := fmt.Sprintf("Lv: %d  EXP: %d/%d", g.Player.Level, g.Player.EXP, g.Player.NextEXP)
	ebitenutil.DebugPrintAt(screen, hpBar, 8, 8)
	ebitenutil.DebugPrintAt(screen, lvl, 8, 24)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), 8, 40)

	if g.GameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER - Press R to Restart", screenWidth/2-100, screenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func circleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	dist := dx*dx + dy*dy
	rSum := r1 + r2
	return dist < rSum*rSum
}

func filterAlive(enemies []*Enemy) []*Enemy {
	result := enemies[:0]
	for _, e := range enemies {
		if e.Alive {
			result = append(result, e)
		}
	}
	return result
}

func filterAliveBullets(bullets []*Bullet) []*Bullet {
	result := bullets[:0]
	for _, b := range bullets {
		if b.Alive {
			result = append(result, b)
		}
	}
	return result
}

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("Vampire Survivors Mini")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
