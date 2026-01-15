package interfaces

// Полиморфизм это возможность работать с объектами через единый интерфейс

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.1f\n", s.Area())
	fmt.Printf("Perimeter: %.1f\n", s.Perimeter())
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * (c.radius * c.radius)
}

func printShapeInfo() {
	shapes := []Shape{
		Circle{radius: 5},
		Circle{radius: 6},
		Circle{radius: 8},
		Rectangle{width: 10, height: 5},
		Rectangle{width: 4, height: 2},
		Rectangle{width: 8, height: 5},
	}

	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
}
