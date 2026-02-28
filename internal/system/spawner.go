package system

import "github.com/atEaE/go-space/internal/entity"

type EnemySpawner struct {
	Timer int
	Rate  int
}

func NewEnemySpawner() *EnemySpawner {
	return &EnemySpawner{
		Rate: 60,
	}
}

func (s *EnemySpawner) Update(tick int, playerX, playerY float64) *entity.Enemy {
	s.Timer++
	currentRate := max(s.Rate-tick/600, 15)
	if s.Timer >= currentRate {
		s.Timer = 0
		return entity.NewEnemy(playerX, playerY)
	}
	return nil
}
