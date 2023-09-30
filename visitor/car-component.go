package visitor

import "fmt"

type CarComponent struct {
}

func (c *CarComponent) Accept(visitor IVisitor) {
	visitor.visitCarComponent(c)
}

func (c *CarComponent) carExecute() {
	fmt.Println("Car component executing...")
}
