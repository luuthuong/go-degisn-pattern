package main

import "design-pattern/visitor"

func main() {
	car := visitor.CarComponent{}
	concrete := visitor.ConcreteVisitor{}
	concreteOther := visitor.ConcreteOther{}
	car.Accept(concrete)
	car.Accept(concreteOther)
}
