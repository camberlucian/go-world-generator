package geogen

import (
	"errors"
	"fmt"

	"github.com/camberlucian/go-world-generator/filemanager"
	"github.com/camberlucian/go-world-generator/geogen/types"
)

func GenerateBasicMap(xVal int, yVal int, name string) *types.World {
	worldMap := [][]types.Tile{}
	fmt.Println(types.GeoTypes[2])
	fmt.Println(types.Symbols[types.GeoTypes[2]])
	for i := 0; i < yVal; i++ {
		row := []types.Tile{}
		for j := 0; j < xVal; j++ {
			tile := types.Tile{
				X:       j,
				Y:       i,
				GeoType: 2,
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
