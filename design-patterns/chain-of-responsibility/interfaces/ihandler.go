package interfaces

type IHandler interface {
	Next(handler IHandler) IHandler
	Handle()
}
