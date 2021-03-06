package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
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
	 if utf8.RuneCountInString(word) > utf8.RuneCountInString(maxWord) {
	  maxWord = word
	 }
	}
	return maxWord
   }

// function that returns the copy of the original slice in reverse order
func reverse(numbers []int64) []int64 {
	reversed := make([]int64, len(numbers))
	for i := 0; i < len(numbers); i++ {
		reversed[i] = numbers[len(numbers) - 1 - i]
	}
	return reversed
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

	words1 := []string{"从", "two", "three"}
	w := maxString(words1)
	fmt.Println("The longest word:", w)

	words2 := []string{"one", "two"}
	maxWord1 := maxString(words2)
	fmt.Println("The longest word of two equals:", maxWord1)

	r := []int64{1, 2, 5, 15}
	fmt.Println("Original slice", r)
	n := reverse(r)
	fmt.Println("The copy of the original slice in reverse order:", n)
	fmt.Println("Original slice after reverse", r)

	key1 := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(key1)

	key2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(key2)
}
