package geometry
import (
	"math"

	"image/color"
)

type Point struct {
	X, Y float64
}

type Path []Point //named slice type

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 { //method; p-reciever
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i - 1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point)ScaleBy(factor float64) { //pointer reciever method(rule : all methods now
	// must be with pointer rec )
	p.X *= factor
	p.Y *= factor

}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

//move all points on offsetffff
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i].Add(offset)
		path[i] = op(path[i], offset)
	}
}

