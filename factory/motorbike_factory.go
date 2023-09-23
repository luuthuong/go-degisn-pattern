package factory

const (
	SportType  = 1
	CruiseType = 2
)

type MotorbikeFactory struct {
}

type SportMotorbike struct {
}

func (s SportMotorbike) GetWheels() uint8 {
	return 8
}

func (s SportMotorbike) GetSeats() uint8 {
	return 2
}

type CruiseMotorbike struct {
}

func (c CruiseMotorbike) GetWheels() uint8 {
	return 6
}

func (c CruiseMotorbike) GetSeats() uint8 {
	return 12
}

func (m *MotorbikeFactory) GetVehicle(v int) (IVehicle, error) {
	switch v {
	case SportType:
		return new(SportMotorbike), nil
	case CruiseType:
		return new(CruiseMotorbike), nil
	default:
		panic("Invalid type motor")
	}
}
