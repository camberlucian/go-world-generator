package main

import (
	"fmt"

	"github.com/camberlucian/go-world-generator/geogen"
)

func main() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicMap(50, 50, "coasts")
	fmt.Println("NORMALIZING ELEVATION")
	world = geogen.NormalizeElevation(world, 4)
	fmt.Println("GENERATING COASTS")
	world = geogen.GenerateCoastalOffset(world, 0, 6, 6, 0)
	world = geogen.FloodMap(world)
	fmt.Println("PRINTING MAP")
	err := geogen.PrintMap(world)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}

}
