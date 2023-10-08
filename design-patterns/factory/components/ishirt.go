package components

type Shirt struct {
	size, name string
}

func (s *Shirt) GetName() string {
	return s.name
}

func (s *Shirt) SetName(name string) {
	s.name = name
}

func (s *Shirt) SetSize(size string) {
	s.size = size
}

func (s *Shirt) GetSize() string {
	return s.size
}
