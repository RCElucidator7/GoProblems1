//Ryan Conway - G00332826

package main

import (
 "fmt"
 "strings"
)

func main() {

	var input string

	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &input)

	input = strings.ToLower(input)
	fmt.Println(pall(input))
}

//Palindrom function that check if the string entered is/is not a palindroms

func pall(s string) string {

	//middle takes the length of the string and halfs it to find the middle
	//Last finds the last letter of the string 
	middle := len(s) / 2
	last := len(s) - 1

	//loop compares the last letter with the first letter 
	for i := 0; i < middle; i++ {

		if s[i] != s[last-i] {
			return "The string you've entered is not a Palimdrome."
		}
	}

	return "The string you've entered is a Palimdrome!."
}