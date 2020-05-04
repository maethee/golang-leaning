package main

import (
	"fmt"
	"sort"
)

func main() {
	SampleSort()
}

func SampleSort() {
	i := []int{2, 4, 38, 4, 70, 123, 5}
	s := []string{"I", "Love", "Dog"}
	sort.Ints(i)
	sort.Strings(s)
	fmt.Println("Sort ", i)
	fmt.Println("Sort ", s)
}
