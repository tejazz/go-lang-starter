package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 10.0, height: 5.0}
	sq := square{sideLength: 4.0}

	printArea(t)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
