package render

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/camberlucian/go-world-generator/filemanager"
)

func RenderSolidImage() {
	m := image.NewRGBA(image.Rect(0, 0, 500, 500))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{C: blue}, image.ZP, draw.Src)
	file, err := filemanager.GetOrCreateFile("firstpic.png", true)
	if err != nil {
		fmt.Println("Could not get or create file: " + err.Error())
	}
	err = png.Encode(file, m)
	if err != nil {
		fmt.Println("COULD NOT WRITE IMAGE: " + err.Error())
	}

}
