package components

type IComponent interface {
	Accept(visitor IVisitor)
	getWheels() uint8
}
