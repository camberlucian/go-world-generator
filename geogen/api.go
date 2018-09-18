package geogen

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/camberlucian/go-world-generator/filemanager"
	"github.com/camberlucian/go-world-generator/geogen/types"
	"github.com/camberlucian/go-world-generator/utils"
)

func GenerateBasicMap(xVal int, yVal int, name string) *types.World {
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
	world := types.World{
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           &worldMap,
	}
	return &world
}

func GenerateThenFloodBasicMap(xVal int, yVal int, name string) *types.World {
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
	world := types.World{
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           &worldMap,
	}
	return &world
}

func GenerateBasicIsland(xVal int, yVal int, offset int, name string) *types.World {
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
	world := types.World{
		PrintedFileName: name + "-Printed.txt",
		CodedFileName:   name + "-Coded.csv",
		Tiles:           &worldMap,
	}
	return &world
}

func PrintMap(world *types.World) error {
	worldMap := world.Tiles
	stringArray := []string{}
	for _, row := range *worldMap {
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

func PrintElevationMap(world *types.World) error {
	worldMap := world.Tiles
	stringArray := []string{}
	for _, row := range *worldMap {
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
