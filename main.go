package main

import (
	f "design-pattern/factory"
	"fmt"
)

func main() {
	carFactory, _ := f.GetFactory(f.CarFactoryType)
	car, _ := carFactory.GetVehicle(f.FamiliarType)
	familiarCar := car.(*f.FamiliarCar)
	fmt.Printf("Familiar car\n")
	fmt.Printf("Doors: %d \n", familiarCar.GetDoors())
	PrintVehicle(car)

	motorbikeFactory, _ := f.GetFactory(f.MotorbikeFactoryType)
	motorbike, _ := motorbikeFactory.GetVehicle(f.SportType)
	fmt.Printf("Sport motobike\n")
	sportMotorbike := motorbike.(*f.SportMotorbike)
	fmt.Printf("Type: %s \n", sportMotorbike.GetType())
	PrintVehicle(motorbike)
}

func PrintVehicle(v f.IVehicle) {
	fmt.Printf("Seats: %d \n", v.GetSeats())
	fmt.Printf("Wheels: %d \n", v.GetWheels())
}
