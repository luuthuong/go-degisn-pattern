package motobike

import (
	"design-pattern/design-patterns/abstract-factory/factory"
)

const (
	SportType  = 1
	CruiseType = 2
)

type Factory struct {
}

type SportMotorbike struct {
}

func (s SportMotorbike) GetWheels() uint8 {
	return 8
}

func (s SportMotorbike) GetSeats() uint8 {
	return 2
}

func (s SportMotorbike) GetType() string {
	return "Sport"
}

type CruiseMotorbike struct {
}

func (c CruiseMotorbike) GetWheels() uint8 {
	return 6
}

func (c CruiseMotorbike) GetSeats() uint8 {
	return 12
}

func (c CruiseMotorbike) GetType() string {
	return "Cruise"
}

func (m *Factory) GetVehicle(v int) (factory.IVehicle, error) {
	switch v {
	case SportType:
		return new(SportMotorbike), nil
	case CruiseType:
		return new(CruiseMotorbike), nil
	default:
		panic("Invalid type motor")
	}
}
