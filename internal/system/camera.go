package system

type Camera struct {
	X, Y float64
}

func (c *Camera) Update(playerX, playerY float64, screenWidth, screenHeight int) {
	c.X = playerX - float64(screenWidth)/2
	c.Y = playerY - float64(screenHeight)/2
}

func (c *Camera) WorldToScreen(wx, wy float64) (float32, float32) {
	return float32(wx - c.X), float32(wy - c.Y)
}
