package main

import "fmt"

func main () {
	t1 := triangle{height: 20, base: 10}
	s1 := square{sideLength: 4}
	printArea(t1)
	printArea(s1)
}

type shape interface {
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}
type square struct{
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape)  {
	fmt.Println(s.getArea())
}
