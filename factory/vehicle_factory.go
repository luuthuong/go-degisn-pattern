package factory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type IVehicleFactory interface {
	GetVehicle(v int) (IVehicle, error)
}

func GetFactory(f int) (IVehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %s invalid.", f))
	}
}
