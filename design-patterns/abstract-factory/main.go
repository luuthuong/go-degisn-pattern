package main

import (
	c "design-pattern/design-patterns/abstract-factory/factory/car"
	m "design-pattern/design-patterns/abstract-factory/factory/motobike"
	"fmt"
)

func main() {
	carFactory, _ := GetFactory(CarFactoryType)
	car, _ := carFactory.GetVehicle(c.FamiliarType)
	familiarCar := car.(*c.FamiliarCar)
	fmt.Printf("Familiar car\n")
	fmt.Printf("Doors: %d \n", familiarCar.GetDoors())
	PrintVehicle(car)

	motorbikeFactory, _ := GetFactory(MotorbikeFactoryType)
	motorbike, _ := motorbikeFactory.GetVehicle(m.SportType)
	fmt.Printf("Sport motobike\n")
	sportMotorbike := motorbike.(*m.SportMotorbike)
	fmt.Printf("Type: %s \n", sportMotorbike.GetType())
	PrintVehicle(motorbike)
}
