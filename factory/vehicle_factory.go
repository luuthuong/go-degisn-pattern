package factory

import (
	c "design-pattern/factory/car"
	m "design-pattern/factory/motobike"
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
		return new(c.Factory), nil
	case MotorbikeFactoryType:
		return new(m.Factory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d invalid.", f))
	}
}
