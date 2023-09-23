package factory

const (
	LuxuryType   = 1
	FamiliarType = 2
)

type CarFactory struct {
}

type LuxuryCar struct{}

func (l LuxuryCar) GetWheels() uint8 {
	return 4
}

func (l LuxuryCar) GetSeats() uint8 {
	return 8
}

func (l LuxuryCar) GetDoors() uint8 {
	return 4
}

type FamiliarCar struct{}

func (f FamiliarCar) GetWheels() uint8 {
	return 8
}

func (f FamiliarCar) GetSeats() uint8 {
	return 12
}

func (f FamiliarCar) GetDoors() uint8 {
	return 6
}

func (c *CarFactory) GetVehicle(v int) (IVehicle, error) {
	switch v {
	case LuxuryType:
		return new(LuxuryCar), nil
	case FamiliarType:
		return new(FamiliarCar), nil
	default:
		panic("invalid type")
	}
}
