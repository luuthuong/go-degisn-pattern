package main

import (
	visitor2 "design-pattern/design-patterns/visistor/components"
	"fmt"
)

func main() {
	car := visitor2.CarComponent{}
	concrete := visitor2.ConcreteVisitor{}
	concreteOther := visitor2.ConcreteOther{}
	car.Accept(concrete)
	fmt.Println("===============")
	car.Accept(concreteOther)
}
