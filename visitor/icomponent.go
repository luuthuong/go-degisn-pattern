package visitor

type IComponent interface {
	Accept(visitor IVisitor)
}
