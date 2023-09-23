package main

import (
	f "design-pattern/factory"
	"fmt"
)

func main() {

	carFactory, _ := f.GetFactory(f.CarFactoryType)
	familyCar, _ := carFactory.GetVehicle(f.FamiliarType)

	fmt.Printf("Familiar car\n")
	PrintVehicle(familyCar)

	motorbikeFactory, _ := f.GetFactory(f.MotorbikeFactoryType)
	sportMotorbike, _ := motorbikeFactory.GetVehicle(f.SportType)
	fmt.Printf("Sport motobike\n")
	PrintVehicle(sportMotorbike)
}

func PrintVehicle(v f.IVehicle) {
	fmt.Printf("Seats: %d \n", v.GetSeats())
	fmt.Printf("Wheels: %d \n", v.GetWheels())
}
