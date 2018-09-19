package main

import (
	"fmt"

	"github.com/camberlucian/go-world-generator/geogen"
	"github.com/camberlucian/go-world-generator/render"
)

func main() {
	render.RenderSolidImage()

}

func GenerateWorld() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicMap(50, 50, -16, 18, "coasts")
	fmt.Println("NORMALIZING ELEVATION")
	world = geogen.NormalizeElevation(world, 7)
	fmt.Println("GENERATING COASTS")
	world = geogen.GenerateCoastalOffset(world, 0, 6, 6, 0)
	fmt.Println("FLOODING MAP")
	world = geogen.FloodMap(world)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 5)
	fmt.Println("REMOVING SMALL LAKES")
	world = geogen.RemoveSmallLakes(world, 2)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 1)
	fmt.Println("REMOVING SMALL LAKES")
	world = geogen.RemoveSmallLakes(world, 1)
	fmt.Println("PRINTING MAP")
	err := geogen.PrintMap(world)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}
}
