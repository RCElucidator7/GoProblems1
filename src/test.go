package main

import "fmt"
import "sort"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 4, 5, 6, 7}

	merged := merge(s1, s2)

	fmt.Println(merged)
}

func merge(s1, s2 []int) []int {
	// Create a map that holds the values from each slice
	unique := make(map[int]struct{}) // zero byte payload

	for _, v := range s1 {
		unique[v] = struct{}{} // zero byte payload
	}

	for _, v := range s2 {
		unique[v] = struct{}{} // zero byte payload
	}

	final := make([]int, len(unique)) // allocate to right size
	i := 0
	for k := range unique {
		final[i] = k
		i++ // index not a feature of map ranges
	}
	sort.Ints(final) // consistent order
	return final
}