package main

import (
	"fmt"

	"github.com/karamel29/go-courses/task002/fibonacci"
)

func main() {
	fmt.Println("Hi")
	defer fmt.Println("Mission accomplished")
	fibonacci.Printer(5)
}
