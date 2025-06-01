package player

import "github.com/hajimehoshi/ebiten/v2"

func (p *Player) Update() (err error) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= .1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += .1
	}

	p.position.Add(p.velocity)

	return nil
}
