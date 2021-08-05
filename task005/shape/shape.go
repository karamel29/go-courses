package shape

import (
	"fmt"
)

type Shape interface {
	String() string
	Area() (float64, error)
	Perimeter() (float64, error)
}

func DescribeShape(s Shape) error {
	a, err := s.Area()
	if err != nil {
		return err
	}
	p, err := s.Perimeter()
	if err != nil {
		return err
	}
	fmt.Println(s)
	fmt.Printf("Perimeter: %.2f\n", p)
	fmt.Printf("Area: %.2f\n", a)
	return nil
}