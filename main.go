package main

import (
	"fmt"

	"github.com/camberlucian/go-world-generator/geogen"
	"github.com/camberlucian/go-world-generator/render"
)

func main() {
	GenerateWorld()

}

func GenerateWorld() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicMap(200, 200, -8, 8, "fullmap")
	fmt.Println("NORMALIZING ELEVATION")
	world = geogen.NormalizeElevation(world, 3)
	fmt.Println("GENERATING COASTS")
	world = geogen.GenerateCoastalOffset(world, 0, 50, 50, 0)
	fmt.Println("FLOODING MAP")
	world = geogen.FloodMap(world)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 20)
	fmt.Println("REMOVING SMALL LAKES")
	world = geogen.RemoveSmallLakes(world, 2)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 1)
	fmt.Println("REMOVING SMALL LAKES")
	world = geogen.RemoveSmallLakes(world, 10)
	fmt.Println("GENERATING RIVERS")
	world = geogen.RiverGen(world, 1)
	fmt.Println("PRINTING MAP")
	//
	err := render.DrawWorldMap(world, 50)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}
}
