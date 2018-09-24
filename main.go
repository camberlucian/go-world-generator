package main

import (
	"fmt"
	"strconv"

	"github.com/camberlucian/go-world-generator/geogen"
	"github.com/camberlucian/go-world-generator/render"
)

func main() {
	GenerateWorld()

}

func GenerateWorld() {
	fmt.Println("GENERATING WORLD")
	world := geogen.GenerateBasicMap(50, 50, -8, 8, "fullmap")
	fmt.Println("NORMALIZING ELEVATION")
	world = geogen.NormalizeElevation(world, 3)
	fmt.Println("GENERATING COASTS")
	world = geogen.GenerateCoastalOffset(world, 0, 8, 8, 0, 7)
	fmt.Println("FLOODING MAP")
	world = geogen.FloodMap(world)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 20)
	fmt.Println("REMOVING SMALL LAKES")
	world = geogen.RemoveSmallLakes(world, 2)
	fmt.Println("REMOVING OUTLIERS")
	world = geogen.RemoveOutliers(world, 1)
	world = geogen.RemoveSmallLakes(world, 2)
	world = geogen.RemoveOutliers(world, 1)
	world = geogen.RaiseLand(world, 1)
	world = geogen.RemoveCoastalPeaks(world, 1)
	peaks := geogen.FindPeaks(world, 4)
	fmt.Println("PEAK COUNT: " + strconv.Itoa(len(peaks)))
	world = geogen.MakeMountainsFromPeaks(world, peaks)
	world = geogen.DisperseHumidity(world, false, true, true, false)
	fmt.Println("PRINTING MAP")
	//
	err := render.DrawHumidityMap(world, 50)
	err = render.DrawElevationMap(world, 50)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("DONE")
	}
}
