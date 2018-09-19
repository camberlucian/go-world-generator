package geogen

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/camberlucian/go-world-generator/filemanager"
	"github.com/camberlucian/go-world-generator/geogen/types"
	"github.com/camberlucian/go-world-generator/utils"
)

type World struct {
	PrintedFileName string
	CodedFileName   string
	Tiles           [][]types.Tile
}

func (w *World) GetTile(y int, x int) *types.Tile {
	return &w.Tiles[y][x]
}

func (w *World) GetSurroundingTiles(y int, x int) []*types.Tile {
	type pair struct {
		Y int
		X int
	}
	width := len(w.Tiles[0]) - 1
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
		if p.X < 0 || p.Y < 0 || p.X > width || p.Y > width {
			continue
		}
		tiles = append(tiles, w.GetTile(p.Y, p.X))
	}
	return tiles
}

func GenerateBasicMap(xVal int, yVal int, name string) *World {
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
	world := World{
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           worldMap,
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
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           worldMap,
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
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           worldMap,
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

func PrintMap(world *World) error {
	worldMap := world.Tiles
	stringArray := []string{}
	for _, row := range worldMap {
		rowString := ""
		for _, col := range row {
			rowString += types.Symbols[types.GeoTypes[col.GeoType]]
		}
		rowString += "\n"
		stringArray = append(stringArray, rowString)
	}
	err := filemanager.WriteStringsToFile(&stringArray, world.PrintedFileName)
	if err != nil {
		return errors.New("PrintMap Error: " + err.Error())
	}
	return nil
}

func PrintElevationMap(world *World) error {
	worldMap := world.Tiles
	stringArray := []string{}
	for _, row := range worldMap {
		rowString := ""
		for _, col := range row {
			if col.Elevation > 0 {
				rowString += (" " + strconv.Itoa(col.Elevation) + " ")
			} else {
				rowString += (strconv.Itoa(col.Elevation) + " ")
			}

		}
		rowString += "\n"
		stringArray = append(stringArray, rowString)
	}
	err := filemanager.WriteStringsToFile(&stringArray, world.PrintedFileName)
	if err != nil {
		return errors.New("PrintMap Error: " + err.Error())
	}
	return nil
}
