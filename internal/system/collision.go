package system

func CircleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	dist := dx*dx + dy*dy
	rSum := r1 + r2
	return dist < rSum*rSum
}

func IsOffScreen(x, y, camX, camY float64, screenWidth, screenHeight int) bool {
	sx := x - camX
	sy := y - camY
	margin := 100.0
	return sx < -margin || sx > float64(screenWidth)+margin || sy < -margin || sy > float64(screenHeight)+margin
}
