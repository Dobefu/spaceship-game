package bullet

func (b *Bullet) Update() (err error) {
	b.GetPosition().Add(b.velocity)

	return nil
}
