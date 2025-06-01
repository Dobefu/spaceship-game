package camera

func (c *Camera) Update() {
	if c.target == nil {
		return
	}

	c.position = c.target.GetPosition()
}
