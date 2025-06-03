package bullet

func (b *Bullet) Update() (err error) {
	// TODO: Move this to the GameObject struct.
	if !b.isActive {
		return nil
	}

	return nil
}
