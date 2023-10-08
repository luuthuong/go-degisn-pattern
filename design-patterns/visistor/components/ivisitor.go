package components

type IVisitor interface {
	visitCarComponent(component IComponent)
	visitMotorbikeComponent(component IComponent)
}
