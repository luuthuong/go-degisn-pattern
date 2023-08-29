package factory

type IOutfit interface {
	GetName() string
	SetName(name string)
	SetSize(size string)
	GetSize() string
}
