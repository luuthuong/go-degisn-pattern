package components

import "fmt"

type ConcreteVisitor struct {
}

func (c ConcreteVisitor) visitCarComponent(component IComponent) {
	car := component.(*CarComponent)
	fmt.Printf("Visit from concrete.\n")
	fmt.Printf("Num of motorbike wheels: %d\n", car.getWheels())
}

func (c ConcreteVisitor) visitMotorbikeComponent(component IComponent) {
	motorbike := component.(*MotorbikeComponent)
	fmt.Printf("Visit from concrete.\n")
	fmt.Printf("Num of motorbike wheels: %d\n", motorbike.getWheels())
}

type ConcreteOther struct {
}

func (c ConcreteOther) visitCarComponent(component IComponent) {
	car := component.(*CarComponent)
	fmt.Printf("Visit from concrete other.\n")
	fmt.Printf("Num of motorbike wheels: %d\n", car.getWheels())
}

func (c ConcreteOther) visitMotorbikeComponent(component IComponent) {
	motorbike := component.(*MotorbikeComponent)
	fmt.Printf("Visit from concrete other.\n")
	fmt.Printf("Num of motorbike wheels: %d\n", motorbike.getWheels())
}
