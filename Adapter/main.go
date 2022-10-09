package main

import "fmt"

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Line []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(image RasterImage) {
	maxX, maxY := 0, 0
	points := image.GetPoints()
	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	maxX += 1
	maxY += 1
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			found := false
			for _, p := range points {
				if p.X == x && p.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

type vectorToRasterAdapter struct {
	points []Point
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func VectorToRaster(v *VectorImage) RasterImage {
	points := []Point{}
	for _, line := range v.Line {
		points = append(points, Point{line.X1, line.Y1})
		points = append(points, Point{line.X2, line.Y2})
	}
	return vectorToRasterAdapter{points}
}

func main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	DrawPoints(a)
}
