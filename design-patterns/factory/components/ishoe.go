package components

type Shoe struct {
	name, size string
}

func (s *Shoe) GetName() string {
	return s.name
}

func (s *Shoe) SetName(name string) {
	s.name = name
}

func (s *Shoe) SetSize(size string) {
	s.size = size
}

func (s *Shoe) GetSize() string {
	return s.size
}
