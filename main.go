package main

import (
	"fmt"

	"github.com/camberlucian/go-world-generator/geogen"
)

func main() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicIsland(50, 50, 6, "default")
	fmt.Println("PRINTING MAP")
	err := geogen.PrintMap(world)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}

}
