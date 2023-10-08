package components

type CarComponent struct {
}

func (c *CarComponent) getWheels() uint8 {
	return 4
}

func (c *CarComponent) Accept(visitor IVisitor) {
	visitor.visitCarComponent(c)
}
