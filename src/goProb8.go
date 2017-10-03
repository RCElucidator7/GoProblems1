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

func main() {
	
	s := []int{
		54,35,57,97,
		23,13,31,74,
		54,37,63,27,
		16,98,7,1,
	}
	fmt.Printf("%v\n%v\n", s, MergeSort(s))
}
