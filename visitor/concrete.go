package visitor

import "fmt"

type ConcreteVisitor struct {
}

func (c ConcreteVisitor) visitCarComponent(component IComponent) {
	car := component.(*CarComponent)
	fmt.Printf("Visit from concrete.\n")
	car.carExecute()
}

func (c ConcreteVisitor) visitMotorbikeComponent(component IComponent) {
	motorbike := component.(*MotorbikeComponent)
	fmt.Printf("Visit from concrete.\n")
	motorbike.motorbikeExecute()
}

type ConcreteOther struct {
}

func (c ConcreteOther) visitCarComponent(component IComponent) {
	car := component.(*CarComponent)
	fmt.Printf("Visit from concrete other.\n")
	car.carExecute()
}

func (c ConcreteOther) visitMotorbikeComponent(component IComponent) {
	car := component.(*CarComponent)
	fmt.Printf("Visit from concrete other.\n")
	car.carExecute()
}
