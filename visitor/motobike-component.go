package visitor

type MotorbikeComponent struct {
}

func (m *MotorbikeComponent) getWheels() uint8 {
	return 2
}

func (m *MotorbikeComponent) Accept(visitor IVisitor) {
	visitor.visitMotorbikeComponent(m)
}
