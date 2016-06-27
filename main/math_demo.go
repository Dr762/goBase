package main

import (

"github.com/goBase/geometry"
	"fmt"
	"image/color"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4,6}

	fmt.Println(p.Distance(q))

	var distanceFromP = p.Distance  //method value
	fmt.Println(distanceFromP(q))
	perim :=geometry.Path{
		{1,1},
		{4,1},
		{6,4},
		{1,1},
	}

	fmt.Println(perim.Distance())

	(&p).ScaleBy(6)
	fmt.Println(p)

	blue:= color.RGBA{255,0,0,255}
	cp := geometry.ColoredPoint{p,blue}
	fmt.Println(cp.Point.X)

	p1 := p.Add(q)
	fmt.Println(p1)

	p2 := p.Sub(q)
	fmt.Println(p2)

	perim.TranslateBy(p,true)
	fmt.Println(perim)
}
