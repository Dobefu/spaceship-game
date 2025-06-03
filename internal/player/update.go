package player

func (p *Player) Update() (err error) {
	p.HandleMovement()

	return nil
}
