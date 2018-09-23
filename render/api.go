package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/camberlucian/go-world-generator/geogen"

	"github.com/camberlucian/go-world-generator/filemanager"
	"github.com/camberlucian/go-world-generator/render/types"
)

func RenderSolidImage() {
	m := image.NewRGBA(image.Rect(0, 0, 500, 500))
	blue := color.RGBA{255, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{C: blue}, image.ZP, draw.Src)
	filemanager.ExportImage(m, "firstpic.png")
}

func DrawElevationMap(world *geogen.World, multiplier int) error {
	colors := types.Colors
	eColors := types.EColors
	canvas := image.NewRGBA(image.Rect(0, 0, (multiplier * world.Width), (multiplier * world.Height)))
	worldMap := world.Tiles
	for k := 0; k < len(worldMap); k++ {
		row := &worldMap[k]
		for l := 0; l < len(*row); l++ {
			rX := multiplier * l
			rY := multiplier * k
			r := image.Rect(rX, rY, (rX + multiplier), (rY + multiplier))
			col := world.GetTile(k, l)
			if col.GeoType == 1 {
				draw.Draw(canvas, r, &image.Uniform{colors[col.GeoType]}, image.ZP, draw.Src)
			} else {
				draw.Draw(canvas, r, &image.Uniform{eColors[col.Elevation]}, image.ZP, draw.Src)
			}

		}
	}
	err := filemanager.ExportImage(canvas, world.ElevationFileName)
	if err != nil {
		fmt.Println("FAILED TO ENCODE DRAWING")
	}
	return err
}

func DrawHumidityMap(world *geogen.World, multiplier int) error {
	colors := types.Colors
	hColors := types.HColors
	canvas := image.NewRGBA(image.Rect(0, 0, (multiplier * world.Width), (multiplier * world.Height)))
	worldMap := world.Tiles
	for k := 0; k < len(worldMap); k++ {
		row := &worldMap[k]
		for l := 0; l < len(*row); l++ {
			rX := multiplier * l
			rY := multiplier * k
			r := image.Rect(rX, rY, (rX + multiplier), (rY + multiplier))
			col := world.GetTile(k, l)
			if col.GeoType == 1 {
				draw.Draw(canvas, r, &image.Uniform{colors[col.GeoType]}, image.ZP, draw.Src)
			} else {
				draw.Draw(canvas, r, &image.Uniform{hColors[col.Humidity]}, image.ZP, draw.Src)
			}

		}
	}
	err := filemanager.ExportImage(canvas, world.HumidityFileName)
	if err != nil {
		fmt.Println("FAILED TO ENCODE DRAWING")
	}
	return err
}
