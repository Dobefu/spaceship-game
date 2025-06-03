package player

func (p *Player) Shoot() {
	for _, b := range *p.bulletPool {
		// Skip already active bullets.
		if b.GetIsActive() {
			continue
		}

		b.Fire(*p.GetPosition(), p.rotation)
		break
	}
}
