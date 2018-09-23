package types

type Tile struct {
	X         int
	Y         int
	Elevation int
	Humidity  int
	GeoType   int
	Resources []int
	Locations []int
}

var GeoTypes = map[int]string{
	1: "Ocean",
	2: "Land",
}

var Symbols = map[string]string{
	"Ocean": " _ ",
	"Land":  "[ ]",
}
