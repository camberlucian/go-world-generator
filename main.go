package main

import (
	"fmt"

	"github.com/camberlucian/go-world-generator/geogen"
)

func main() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicMap(50, 50, "normalized")
	world = geogen.NormalizeElevation(world, 4)
	world = geogen.FloodMap(world)
	fmt.Println("PRINTING MAP")
	err := geogen.PrintMap(world)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}

}
