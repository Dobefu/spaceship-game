package camera

func (c *Camera) Update() {
	c.position = c.target.GetPosition()
}
