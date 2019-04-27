package main

import (
	"fmt"
)

type Vertex struct {
	X float32
	Y int
}

type Point struct {
	px float32
	py float32
}

func (p *Point) setXY(px, py float32) {
	p.px = px
	p.py = py
}

func (p *Point) getXY() (float32, float32) {
	return p.px, p.py
}

func main() {
	/*	fmt.Println(Vertex{1.2, 2})
		var v = Vertex{1, 2}
		v.X = 4
		fmt.Println(v.X)
		fmt.Println(v)

		p := &v
		p.X = 1e9
		fmt.Println(v)*/

	//point := new(Point)

	var point = Point{}
	point.setXY(1.23, 4.56)
	px, py := point.getXY()
	fmt.Println(px, py)

	var point2 = Point{}
	//point2.setXY(1.23, 4.56)
	px2, py2 := point2.getXY()
	fmt.Println(px2, py2)
}
