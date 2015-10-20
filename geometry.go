package picfit

import (
	"fmt"
	"regexp"
	"strconv"
)

// Geometry struct.
type Geometry struct {
	X int
	Y int
}

// GeometryRegex is the regex to extract geometry integers
var GeometryRegex = regexp.MustCompile(`^(?P<x>\d+)?(?:x(?P<y>\d+))?`)

// ParseGeometry parses a geometry string syntax and returns a width/height integers.
func ParseGeometry(geometry string) (*Geometry, error) {
	x, y := 0, 0
	names := GeometryRegex.SubexpNames()
	matches := GeometryRegex.FindStringSubmatch(geometry)
	matchesMap := map[string]string{}

	for i, value := range matches {
		matchesMap[names[i]] = value
	}

	v, ok := matchesMap["x"]
	if ok {
		x, _ = strconv.Atoi(v)
	}

	v, ok = matchesMap["y"]
	if ok {
		y, _ = strconv.Atoi(v)
	}

	if x == 0 && y == 0 {
		return nil, fmt.Errorf("Geometry does not have the correct syntax: %s", geometry)
	}

	return &Geometry{X: x, Y: y}, nil
}

// ParseGeometryWithRatio parses a geometry string syntax with a given ratio and returns a width/height integers.
func ParseGeometryWithRatio(geometry string, ratio float64) (*Geometry, error) {
	g, err := ParseGeometry(geometry)
	if err != nil {
		return nil, err
	}

	x := g.X
	y := g.Y

	if g.X == 0 {
		x = int(float64(g.Y) * ratio)
	}

	if g.Y == 0 {
		y = int(float64(g.X) / ratio)
	}

	return &Geometry{X: x, Y: y}, nil
}
