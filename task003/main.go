package main

import (
	"fmt"
	"sort"
)

// function that returns an average value of array
func avgArrays(s [6]int) float64 {
	var sum float64
	for _, number := range s {
		x := float64(number)
		sum += x
	}
	n := sum / float64(len(s))
	return n
}

// function that returns the longest word from the slice of strings
func maxString(words []string) string {
	var maxWord string
	for _, word := range words {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	return maxWord
}

// function that returns the copy of the original slice in reverse order
func reverse(numbers []int64) []int64 {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// function that prints map values sorted in order of increasing keys
func printSorted(m map[int]string) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var values []string
	for _, key := range keys {
		values = append(values, m[key])
	}
	fmt.Println("Map values sorted in order of increasing keys:", values)
}

func main() {
	z := [6]int{1, 2, 3, 4, 5, 6}
	y := avgArrays(z)
	fmt.Println("An average value of array:", y)

	words1 := []string{"one", "two", "three"}
	w := maxString(words1)
	fmt.Println("The longest word:", w)

	words2 := []string{"one", "two"}
	maxWord1 := maxString(words2)
	fmt.Println("The longest word of two equals:", maxWord1)

	r := []int64{1, 2, 5, 15}
	n := reverse(r)
	fmt.Println("The copy of the original slice in reverse order:", n)

	key1 := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(key1)

	key2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(key2)
}
