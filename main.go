package main

import (
	f "design-pattern/factory"
	c "design-pattern/factory/car"
	m "design-pattern/factory/motobike"
	"fmt"
)

func main() {
	carFactory, _ := f.GetFactory(f.CarFactoryType)
	car, _ := carFactory.GetVehicle(c.FamiliarType)
	familiarCar := car.(*c.FamiliarCar)
	fmt.Printf("Familiar car\n")
	fmt.Printf("Doors: %d \n", familiarCar.GetDoors())
	PrintVehicle(car)

	motorbikeFactory, _ := f.GetFactory(f.MotorbikeFactoryType)
	motorbike, _ := motorbikeFactory.GetVehicle(m.SportType)
	fmt.Printf("Sport motobike\n")
	sportMotorbike := motorbike.(*m.SportMotorbike)
	fmt.Printf("Type: %s \n", sportMotorbike.GetType())
	PrintVehicle(motorbike)
}

func PrintVehicle(v f.IVehicle) {
	fmt.Printf("Seats: %d \n", v.GetSeats())
	fmt.Printf("Wheels: %d \n", v.GetWheels())
}
