package circle

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

//Print information about radius
func (c Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %0.2f", c.Radius)
}

//Calculation of perimeter for struct Circle
func (c Circle) Perimeter() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("negative value of the radius")
	} else if c.Radius == 0 {
		return 0, errors.New("the radius value is 0")
	}
	return 2 * math.Pi * c.Radius, nil
}

//Calculation of area for struct Circle
func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("negative value of the radius")
	} else if c.Radius == 0 {
		return 0, errors.New("the radius value is 0")
	}
	return math.Pow(c.Radius, 2) * math.Pi, nil
}
