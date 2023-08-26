package main

import (
	"design-pattern/singleton"
	"fmt"
)

func main() {
	instance := singleton.GetInstanceByMutex()
	instance.Name = "Config Name"
	fmt.Printf("%s", instance.Name)
}
