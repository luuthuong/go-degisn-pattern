package main

import (
	i "design-pattern/design-patterns/chain-of-responsibility/interfaces"
	"fmt"
)

type Login struct {
	nextHandler i.IHandler
}

func (l *Login) Next(handler i.IHandler) i.IHandler {
	l.nextHandler = handler
	return handler
}

func (l *Login) Handle() {
	fmt.Println("Login handler")
	if l.nextHandler == nil {
		return
	}
	l.nextHandler.Handle()
}

type Authentication struct {
	nextHandler i.IHandler
}

func (a *Authentication) Next(handler i.IHandler) i.IHandler {
	a.nextHandler = handler
	return handler
}

func (a *Authentication) Handle() {
	fmt.Println("Authentication handler")
	if a.nextHandler == nil {
		return
	}
	a.nextHandler.Handle()
}

type Authorization struct {
	nextHandler i.IHandler
}

func (a *Authorization) Next(handler i.IHandler) i.IHandler {
	a.nextHandler = handler
	return handler
}

func (a *Authorization) Handle() {
	fmt.Println("Authorization handler")
	if a.nextHandler == nil {
		return
	}
	a.nextHandler.Handle()
}

func main() {
	login := &Login{}
	authen := &Authentication{}
	author := &Authorization{}
	login.Next(authen).Next(author)
	login.Handle()
}
