package main

import (
	"fmt"
	"math"
)

type geometry interface {
	calculateArea() float64
	calculatePerimeter() float64
}

func printArea(g geometry) {
	fmt.Println(g.calculateArea())
}
func printPerimeter(g geometry) {
	fmt.Println(g.calculatePerimeter())
}

type circle struct {
	radio float64
}

type square struct {
	side float64
}

func (circle *circle) calculateArea() float64 {
	return 3.14159 * math.Pow(circle.radio, 2)
}
func (circle *circle) calculatePerimeter() float64 {
	return 3.14159 * (circle.radio * 2)
}

func (s *square) calculateArea() float64 {
	return math.Pow(s.side, 2)
}

func (s *square) calculatePerimeter() float64 {
	return s.side * 4
}

func main() {
	square1 := square{side: 3}
	circle1 := circle{radio: 4}

	printArea(&square1)
	printArea(&circle1)
	printPerimeter(&square1)
	printPerimeter(&circle1)

	//geometry1 := make([]geometry, 0)
	//geometry1 = append(geometry1, &square{side: 6})
	//geometry1 = append(geometry1, &circle{radio: 4})
	//
	//for _, g := range geometry1 {
	//	printArea(g)
	//	printPerimeter(g)
	//}
}