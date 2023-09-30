package main

import (
	"design-pattern/visitor"
	"fmt"
)

func main() {
	car := visitor.CarComponent{}
	concrete := visitor.ConcreteVisitor{}
	concreteOther := visitor.ConcreteOther{}
	car.Accept(concrete)
	fmt.Println("===============")
	car.Accept(concreteOther)
}
