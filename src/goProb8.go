//Ryan Conway - G00332826

package main

import (
	"fmt"
)

//Adapted from https://play.golang.org/p/Ma2GXvj3XP
func Merge(left, right []int) []int {

	ret := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right...)
		}
		if len(right) == 0 {
			return append(ret, left...)
		}
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {

	if len(s) <= 1 {
		return s
	}
	//halfs the length of the array
	half := len(s) / 2
	//calls mergeSort on the left side of the array
	left := MergeSort(s[:half])
	//calls mergeSort on the left side of the array
	right := MergeSort(s[half:])
	return Merge(left, right)
}

// Adapted from https://play.golang.org/p/g3QTWcy9D-
func mergeArray(s1, s2 []int) []int {

	unique := make(map[int]struct{})

	for _, v := range s1 {
		unique[v] = struct{}{}
	}

	for _, v := range s2 {
		unique[v] = struct{}{} 
	}

	final := make([]int, len(unique)) 
	i := 0

	for k := range unique {
		final[i] = k
		i++ 
	}
	return final
}

func main() {
	
	a := []int{
		54,35,57,97,
		23,13,31,74,
	}

	s := []int{
		59,37,63,27,
		16,98,7,1,
	}

	merged := mergeArray(a, s)

	fmt.Printf("Array No1: %v\nArray No2: %v\nMerged Array: %v\n", a, s, MergeSort(merged))
}