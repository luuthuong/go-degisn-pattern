package visitor

import "fmt"

type MotorbikeComponent struct {
}

func (m *MotorbikeComponent) Accept(visitor IVisitor) {
	visitor.visitMotorbikeComponent(m)
}

func (m *MotorbikeComponent) motorbikeExecute() {
	fmt.Println("Motorbike component executing....")
}
