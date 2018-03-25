package main

import (
	"fmt"
	"math"
)

//Shape define an interface
type Shape interface {
	area() float64
}

//Circle define a circle
type Circle struct {
	x, y, radius float64
}

//Rectangle define a rectangle */
type Rectangle struct {
	width, height float64
}

// define a method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for rectangle (implementation of Shape.area())
func (rectangle Rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 3, height: 4}
	fmt.Printf("Circle Area : %f \n", circle.area())
	fmt.Printf("Rectangle Area : %f \n", rectangle.area())
	fmt.Printf("Rectangle Area from Shape.getArea : %f \n", getArea(rectangle))
}
