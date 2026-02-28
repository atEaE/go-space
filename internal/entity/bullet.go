package entity

type Bullet struct {
	Pos    Position
	VX, VY float64
	Radius float64
	Damage int
	Alive  bool
}

func (b *Bullet) Update() {
	b.Pos.X += b.VX
	b.Pos.Y += b.VY
}
