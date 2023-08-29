package factory

import "errors"

type SportType uint8

const (
	ShoeFactory  SportType = 0
	ShirtFactory SportType = 1
)

func GetFactory(factoryType SportType) (IOutfit, error) {
	switch factoryType {
	case ShoeFactory:
		return &Shoe{}, nil
	case ShirtFactory:
		return &Shirt{}, nil
	default:
		panic(errors.New("factory not found"))
	}
}
