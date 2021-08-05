package main

import (
	"github.com/karamel29/go-courses/task005/shape/circle"
	"github.com/karamel29/go-courses/task005/shape/rectangle"
	"github.com/karamel29/go-courses/task005/shape"
)

func main() {
	c := circle.Circle{Radius: 6}
	r := rectangle.Rectangle{
		Height: 8,
		Width:  2,
	}
	err := shape.DescribeShape(c)
	if err != nil{
		panic(err)
	}
	err = shape.DescribeShape(r)
	if err != nil{
		panic(err)
	}
}
