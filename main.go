package main

import (
	"design-pattern/factory"
	"fmt"
)

func main() {
	shirt, _ := factory.GetFactory(factory.ShirtFactory)
	shoe, _ := factory.GetFactory(factory.ShoeFactory)

	shirt.SetName("T-Shirt")
	shoe.SetName("Sneaker")

	fmt.Printf("Shirt created: %s\nShoe created: %s", shirt.GetName(), shoe.GetName())
}
