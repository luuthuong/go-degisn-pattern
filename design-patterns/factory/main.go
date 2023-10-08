package main

import (
	"design-pattern/design-patterns/factory/components"
	"fmt"
)

func main() {
	shirt, _ := components.GetFactory(components.ShirtFactory)
	shoe, _ := components.GetFactory(components.ShoeFactory)

	shirt.SetName("T-Shirt")
	shoe.SetName("Sneaker")

	fmt.Printf("Shirt created: %s\nShoe created: %s", shirt.GetName(), shoe.GetName())
}
