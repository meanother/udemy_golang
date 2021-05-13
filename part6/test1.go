package main

import "fmt"

type shape interface {
	getArea() float64
}

type trianlge struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := trianlge{base: 10.5, height: 2.5}
	sq := square{sideLength: 5}

	printArea(tr)
	printArea(sq)
}

func (t trianlge) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
