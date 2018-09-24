package geogen

import (
	"fmt"
	"strconv"

	"github.com/camberlucian/go-world-generator/geogen/types"
	"github.com/camberlucian/go-world-generator/utils"
)

type World struct {
	PrintedFileName   string
	CodedFileName     string
	ElevationFileName string
	HumidityFileName  string
	Height            int
	Width             int
	Tiles             [][]types.Tile
}

func (w *World) GetTile(y int, x int) *types.Tile {
	// fmt.Println(strconv.Itoa(y) + " " + strconv.Itoa(x))
	return &w.Tiles[y][x]
}

func (w *World) GetSurroundingTiles(y int, x int) []*types.Tile {
	type pair struct {
		Y int
		X int
	}
	width := w.Width
	potentials := []pair{}
	tiles := []*types.Tile{}
	potentials = append(potentials, pair{y + 1, x - 1})
	potentials = append(potentials, pair{y + 1, x})
	potentials = append(potentials, pair{y + 1, x + 1})
	potentials = append(potentials, pair{y, x - 1})
	potentials = append(potentials, pair{y, x + 1})
	potentials = append(potentials, pair{y - 1, x - 1})
	potentials = append(potentials, pair{y - 1, x})
	potentials = append(potentials, pair{y - 1, x + 1})
	for _, p := range potentials {
		if p.X < 0 || p.Y < 0 || p.X > width-1 || p.Y > w.Height-1 {
			continue
		}
		tiles = append(tiles, w.GetTile(p.Y, p.X))
	}
	return tiles
}

func GenerateBasicMap(xVal int, yVal int, minElev int, maxElev int, name string) *World {
	worldMap := [][]types.Tile{}
	fmt.Println(types.GeoTypes[2])
	fmt.Println(types.Symbols[types.GeoTypes[2]])
	for i := 0; i < yVal; i++ {
		row := []types.Tile{}
		for j := 0; j < xVal; j++ {
			tile := types.Tile{
				X:         j,
				Y:         i,
				GeoType:   2,
				Elevation: utils.Random(minElev, maxElev),
			}
			row = append(row, tile)
		}
		worldMap = append(worldMap, row)
	}
	world := World{
		PrintedFileName:   name + "-Printed.txt",
		CodedFileName:     name + "-Coded.csv",
		ElevationFileName: name + "-Elevation.png",
		HumidityFileName:  name + "-Humidity.png",
		Tiles:             worldMap,
		Height:            yVal,
		Width:             xVal,
	}
	return &world
}

func GenerateThenFloodBasicMap(xVal int, yVal int, name string) *World {
	worldMap := [][]types.Tile{}
	fmt.Println(types.GeoTypes[2])
	fmt.Println(types.Symbols[types.GeoTypes[2]])
	for i := 0; i < yVal; i++ {
		row := []types.Tile{}
		for j := 0; j < xVal; j++ {
			tile := types.Tile{
				X:         j,
				Y:         i,
				GeoType:   2,
				Elevation: utils.Random(-8, 8),
			}
			row = append(row, tile)
		}
		worldMap = append(worldMap, row)
	}
	for k := 0; k < len(worldMap); k++ {
		row := &worldMap[k]
		for l := 0; l < len(*row); l++ {
			col := &worldMap[k][l]
			if col.Elevation < 0 {
				col.GeoType = 1
			}
		}
	}
	world := World{
		PrintedFileName:   name + "-Printed.txt",
		CodedFileName:     name + "-Coded.csv",
		ElevationFileName: name + "-Elevation.png",
		HumidityFileName:  name + "-Humidity.png",
		Tiles:             worldMap,
		Height:            yVal,
		Width:             xVal,
	}
	return &world
}

func GenerateBasicIsland(xVal int, yVal int, offset int, name string) *World {
	worldMap := [][]types.Tile{}
	fmt.Println(types.GeoTypes[2])
	fmt.Println(types.Symbols[types.GeoTypes[2]])
	for i := 0; i < yVal; i++ {
		row := []types.Tile{}
		for j := 0; j < xVal; j++ {
			typeInt := 2
			if i < offset || j < offset || i >= yVal-offset || j >= xVal-offset {
				typeInt = 1
			}
			tile := types.Tile{
				X:         j,
				Y:         i,
				GeoType:   typeInt,
				Elevation: utils.Random(-8, 8),
			}
			row = append(row, tile)
		}
		worldMap = append(worldMap, row)
	}
	world := World{
		PrintedFileName:   name + "-Printed.txt",
		CodedFileName:     name + "-Coded.csv",
		ElevationFileName: name + "-Elevation.png",
		HumidityFileName:  name + "-Humidity.png",
		Tiles:             worldMap,
		Height:            yVal,
		Width:             xVal,
	}
	return &world
}

func FloodMap(world *World) *World {
	worldMap := world.Tiles
	for k := 0; k < len(worldMap); k++ {
		row := &worldMap[k]
		for l := 0; l < len(*row); l++ {
			col := world.GetTile(k, l)
			if col.Elevation < 0 {
				col.GeoType = 1
			}
		}
	}
	return world
}

func NormalizeElevation(world *World, passes int) *World {
	worldMap := world.Tiles
	for n := 0; n < passes; n++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				tiles := world.GetSurroundingTiles(k, l)
				AverageElev := 0
				for _, t := range tiles {
					AverageElev += t.Elevation
				}
				AverageElev = AverageElev / len(tiles)
				if tile.Elevation > AverageElev {
					tile.Elevation--
				} else if tile.Elevation < AverageElev {
					tile.Elevation++
				}
			}
		}

	}
	return world

}

func GenerateCoastalOffset(world *World, tOffset int, rOffset int, bOffset int, lOffset int, maxDepth int) *World {
	worldMap := world.Tiles
	if rOffset > 0 {
		for r := 0; r <= rOffset; r++ {
			// slope := float64(maxDepth) / float64(rOffset)
			// delevation := maxDepth - math.Round(r*slope)
			// fmt.Println("DELEVATING BY")
			// fmt.Println(delevation)
			// fmt.Println(slope)
			for y := 0; y < len(worldMap); y++ {
				row := &worldMap[y]
				for x := 0; x < len(*row); x++ {
					if x == world.Width-(rOffset-r) {
						col := world.GetTile(y, x)
						col.Elevation = col.Elevation - r
						if col.Elevation < 1 {
							col.GeoType = 1
						}
					}
				}
			}
		}
	}

	if lOffset > 0 {
		for l := 0; l <= lOffset; l++ {
			for y := 0; y < len(worldMap); y++ {
				row := &worldMap[y]
				for x := 0; x < len(*row); x++ {
					if x == 0+(lOffset-l) {
						col := world.GetTile(y, x)
						col.Elevation = col.Elevation - l
						if col.Elevation < 1 {
							col.GeoType = 1
						}
					}
				}
			}
		}
	}

	if tOffset > 0 {
		for t := 0; t <= tOffset; t++ {
			for y := 0; y < len(worldMap); y++ {
				row := &worldMap[y]
				for x := 0; x < len(*row); x++ {
					if y == 0+(tOffset-t) {
						col := world.GetTile(y, x)
						col.Elevation = col.Elevation - t
						if col.Elevation < 1 {
							col.GeoType = 1
						}
					}
				}
			}

		}
	}

	if bOffset > 0 {
		for b := 0; b <= bOffset; b++ {
			for y := 0; y < len(worldMap); y++ {
				row := &worldMap[y]
				for x := 0; x < len(*row); x++ {
					if y == world.Height-(bOffset-b) {
						col := world.GetTile(y, x)
						col.Elevation = col.Elevation - b
						if col.Elevation < 1 {
							col.GeoType = 1
						}
					}
				}
			}

		}
	}
	return world
}

func RemoveOutliers(world *World, passes int) *World {
	worldMap := world.Tiles
	for i := 0; i < passes; i++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				tiles := world.GetSurroundingTiles(k, l)
				seaCount := 0
				landCount := 0
				for _, t := range tiles {
					if t.GeoType == 1 {
						seaCount++
					}
					if t.GeoType == 2 {
						landCount++
					}
				}
				if seaCount < 3 {
					tile.GeoType = 2
					if tile.Elevation < 1 {
						tile.Elevation++
					}
				}
				if landCount < 3 {
					tile.GeoType = 1
					if tile.Elevation > 0 {
						tile.Elevation = 0
					}
				}
			}
		}
	}
	return world
}

func RiverGen(world *World, passes int) *World {
	worldMap := world.Tiles
	for n := 0; n < passes; n++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				if tile.GeoType == 1 {
					tiles := world.GetSurroundingTiles(k, l)
					for _, t := range tiles {
						if t.Elevation < tile.Elevation {
							t.GeoType = 1
						}
					}
				}
			}
		}

	}
	return world
}

func RemoveSmallLakes(world *World, passes int) *World {
	worldMap := world.Tiles
	for i := 0; i < passes; i++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				if tile.GeoType == 1 {
					tiles := world.GetSurroundingTiles(k, l)
					seaCount := 0
					for _, t := range tiles {
						if t.GeoType == 1 {
							seaCount++
						}
					}
					if seaCount < 4 {
						tile.GeoType = 2
					}
				}

			}
		}
	}
	return world
}

func RemoveCoastalPeaks(world *World, passes int) *World {
	worldMap := world.Tiles
	for n := 0; n < passes; n++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				if tile.Elevation > 3 {
					coast := false
					tiles := world.GetSurroundingTiles(k, l)
					for _, t := range tiles {
						if t.GeoType == 1 {
							coast = true
						}
					}
					if coast {
						tile.Elevation = 0
						tile.GeoType = 1
					}
				}
			}
		}

	}
	return world

}

func FindPeaks(world *World, peaks int) []*types.Tile {
	finalPeaks := []*types.Tile{}
	worldMap := world.Tiles
	for k := 0; k < len(worldMap); k++ {
		row := &worldMap[k]
		for l := 0; l < len(*row); l++ {
			tile := world.GetTile(k, l)
			if tile.Elevation > 4 && len(finalPeaks) < peaks && tile.GeoType == 2 {
				finalPeaks = append(finalPeaks, tile)
			}
		}
	}
	minimumSurroundingPeaks := 8
	for len(finalPeaks) < peaks {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				if tile.Elevation == 4 && len(finalPeaks) < peaks && tile.GeoType == 2 {
					surroundingPeaks := 0
					tiles := world.GetSurroundingTiles(k, l)
					for _, t := range tiles {
						if t.Elevation == 4 {
							surroundingPeaks += 1
						}
					}
					if surroundingPeaks < 1 {
						tile.Elevation--
					}
					if surroundingPeaks >= minimumSurroundingPeaks {
						adj := PeakOrAdjacent(finalPeaks, tiles, tile)
						if adj {
							fmt.Println("Adjascent")
						} else {
							fmt.Println("Not ADJA")
							finalPeaks = append(finalPeaks, tile)
						}
					}
				}
			}
		}
		minimumSurroundingPeaks--
	}
	for i := 0; i < len(finalPeaks); i++ {
		fmt.Println("PEAK " + strconv.Itoa(i) + ": " + strconv.Itoa(finalPeaks[i].Y) + ", " + strconv.Itoa(finalPeaks[i].X))
	}
	return finalPeaks
}

func PeakOrAdjacent(peaks []*types.Tile, surroundingTiles []*types.Tile, p *types.Tile) bool {
	adj := false
	for _, peak := range peaks {
		if p == peak {
			adj = true
		}
	}
	for _, tile := range surroundingTiles {
		for _, peak := range peaks {
			if tile == peak {
				adj = true
			}
		}
	}
	return adj
}

func MakeMountain(world *World, tile *types.Tile) *World {
	// fmt.Println("MAKING MOUNTAIN")
	tiles := world.GetSurroundingTiles(tile.Y, tile.X)
	for _, t := range tiles {
		if t.Elevation < tile.Elevation && t.GeoType == 2 {
			// fmt.Println("TILE ELEVATION:" + strconv.Itoa(t.Elevation))
			t.Elevation += 1
			if t.Elevation == 4 {
				MakeMountain(world, t)
			}
		}
	}
	return world
}

func GradePeaks(world *World) *World {
	worldMap := world.Tiles
	passes := 7
	for passes >= 3 {
		for i := 0; i < passes; i++ {
			for k := 0; k < len(worldMap); k++ {
				row := &worldMap[k]
				for l := 0; l < len(*row); l++ {
					tile := world.GetTile(k, l)
					tiles := world.GetSurroundingTiles(k, l)
					for _, t := range tiles {
						if t.Elevation == passes && tile.Elevation < t.Elevation {
							tile.Elevation = passes - 1
						}
					}
				}
			}
		}
		passes--
	}
	return world

}

func MakeMountainsFromPeaks(world *World, tiles []*types.Tile) *World {
	for _, t := range tiles {
		if t.Elevation < 7 {
			fmt.Println("RAISING ELEVATION")
			t.Elevation += 1
			fmt.Println(t.Elevation)
		}
		MakeMountain(world, t)
	}
	world = GradePeaks(world)
	return world
}

func RaiseLand(world *World, passes int) *World {
	worldMap := world.Tiles
	for n := 0; n < passes; n++ {
		for k := 0; k < len(worldMap); k++ {
			row := &worldMap[k]
			for l := 0; l < len(*row); l++ {
				tile := world.GetTile(k, l)
				if tile.GeoType == 1 {
					tiles := world.GetSurroundingTiles(k, l)
					AverageElev := 0
					for _, t := range tiles {
						AverageElev += t.Elevation
					}
					AverageElev = AverageElev / len(tiles)
					if tile.Elevation > AverageElev {
						for _, t := range tiles {
							t.Elevation++
						}
					}
				}
			}
		}
	}
	return world

}

func DisperseHumidity(world *World, t bool, r bool, b bool, l bool) *World {
	if r {
		fmt.Println("DISPERSING RIGHT HUMIDITY")
		h := 0
		hDrop := 0
		for y := 0; y < world.Height; y++ {
			for x := world.Width - 1; x >= 0; x-- {
				col := world.GetTile(y, x)
				if col.GeoType == 1 {
					h += 3
					col.Humidity = 100
					continue
				}
				if x > 0 {
					nextTile := world.GetTile(y, x-1)
					if nextTile.Elevation > col.Elevation {
						hDrop = 3 * (nextTile.Elevation - col.Elevation)
					} else if h > 10 {
						hDrop = 2
					} else {
						hDrop = 1
					}
					if h > 0 {
						h -= hDrop
						col.Humidity += hDrop
					}
				} else {
					if h > 0 {
						h -= 1
						col.Humidity += 1
					}

				}

			}
		}
	}

	if l {
		fmt.Println("DISPERSING LEFT HUMIDITY")
		h := 0
		hDrop := 0
		for y := 0; y < world.Height; y++ {
			for x := 0; x < world.Width; x++ {
				col := world.GetTile(y, x)
				if col.GeoType == 1 {
					h += 3
					col.Humidity = 100
					continue
				}
				if x < world.Width-1 {
					nextTile := world.GetTile(y, x+1)
					if nextTile.Elevation > col.Elevation {
						hDrop = 3 * (nextTile.Elevation - col.Elevation)
					} else if h > 10 {
						hDrop = 2
					} else {
						hDrop = 1
					}
					if h > 0 {
						h -= hDrop
						col.Humidity += hDrop
					}
				} else {
					if h > 0 {
						h -= 1
						col.Humidity += 1
					}
				}

			}
		}
	}

	if t {
		fmt.Println("DISPERSING TOP HUMIDITY")
		h := 0
		hDrop := 0
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				col := world.GetTile(y, x)
				if col.GeoType == 1 {
					h += 3
					col.Humidity = 100
					continue
				}
				if y < world.Height-1 {
					nextTile := world.GetTile(y+1, x)
					if nextTile.Elevation > col.Elevation {
						hDrop = 3 * (nextTile.Elevation - col.Elevation)
					} else if h > 10 {
						hDrop = 2
					} else {
						hDrop = 1
					}
					if h > 0 {
						h -= hDrop
						col.Humidity += hDrop
					}
				} else {
					if h > 0 {
						h -= 1
						col.Humidity += 1
					}
				}
			}
		}
	}

	if b {
		fmt.Println("DISPERSING BOTTOM HUMIDITY")
		h := 0
		hDrop := 0
		for x := 0; x < world.Width; x++ {
			for y := world.Height - 1; y >= 0; y-- {
				col := world.GetTile(y, x)
				if col.GeoType == 1 {
					h += 3
					col.Humidity = 100
					continue
				}
				if y > 0 {
					nextTile := world.GetTile(y-1, x)
					if nextTile.Elevation > col.Elevation {
						hDrop = 3 * (nextTile.Elevation - col.Elevation)
					} else if h > 10 {
						hDrop = 2
					} else {
						hDrop = 1
					}
					if h > 0 {
						h -= hDrop
						col.Humidity += hDrop
					}
				} else {
					if h > 0 {
						h -= 1
						col.Humidity += 1
					}
				}
			}
		}
	}
	return world
}
