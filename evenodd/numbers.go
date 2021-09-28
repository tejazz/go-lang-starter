package main

import "fmt"

type numbers []int

func newNumberSet() numbers {
	var numSet numbers
	for i := 0; i < 10; i++ {
		numSet = append(numSet, i)
	}

	return numSet
}

func (nums numbers) print() {
	for i, n := range nums {
		status := "even"
		if n%2 != 0 {
			status = "odd"
		}
		fmt.Println(i, status)
	}
}
