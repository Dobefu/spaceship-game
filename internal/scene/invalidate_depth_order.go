package scene

func (s *Scene) InvalidateDepthOrder() {
	s.isDepthOrderDirty = true
}
