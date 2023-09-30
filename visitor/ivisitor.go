package visitor

type IVisitor interface {
	visitCarComponent(component IComponent)
	visitMotorbikeComponent(component IComponent)
}
