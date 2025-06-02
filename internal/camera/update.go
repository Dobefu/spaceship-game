package camera

func (c *Camera) Update() {
	if c.target == nil {
		return
	}

	c.position.X += (c.target.GetPosition().X - c.position.X) / 10
	c.position.Y += (c.target.GetPosition().Y - c.position.Y) / 10
}
