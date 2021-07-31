package rectangle

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Height float64
	Width  float64
}

//Print information about height and width
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %0.2f and width %0.2f", r.Height, r.Width)
}

//Calculation of perimeter for struct Rectangle
func (r Rectangle) Perimeter() (float64, error) {
	if r.Height < 0 {
		return 0, errors.New("negative value of the height")
	} else if r.Height == 0 {
		return 0, errors.New("the height value is 0")
	}
	if r.Width < 0 {
		return 0, errors.New("negative value of the width")
	} else if r.Width == 0 {
		return 0, errors.New("the width value is 0")
	}
	a := r.Width + r.Height
	return 2 * a, nil
}

//Calculation of area for struct Rectangle
func (r Rectangle) Area() (float64, error) {
	if r.Height <= 0 {
		return 0, errors.New("invalid value radius")
	}
	if r.Width <= 0 {
		return 0, errors.New("invalid value radius")
	}
	return r.Width * r.Height, nil
}