package main

import (
	"fmt"
	"math"
)

type Shape interface {
	String() string
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

//Print information about radius
func (c Circle) String() string {
	return fmt.Sprintf("Cirlce radius: %0.2f", c.Radius)
}

//Calculation of perimeter for struct Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//Calculation of area for struct Circle
func (c Circle) Area() float64 {
	r := math.Pow(c.Radius, 2)
	return r * math.Pi
}

type Rectangle struct {
	Height float64
	Width  float64
}

//Print information about height and width
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle height %0.2f and width %0.2f", r.Height, r.Width)
}

//Calculation of perimeter for struct Rectangle
func (r Rectangle) Perimeter() float64 {
	a := r.Width + r.Height
	return 2 * a
}

//Calculation of area for struct Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Value of perimeter is: %.2f\n", s.Perimeter())
	fmt.Printf("Value of area is: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 6}
	r := Rectangle{
		Height: 8,
		Width:  2,
	}
	DescribeShape(c)
	DescribeShape(r)
}
