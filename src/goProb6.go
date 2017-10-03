//Ryan Conway - G00332826

package main
import (
	"fmt"
)

func main() {
	var temp, small, large int
	x := []int{
		54,35,57,97,
		23,13,31,74,
		54,37,63,27,
		16,98,7,1,
	}

	//Loops through the array and prints each number
	fmt.Printf("This is the array of numbers: %v \n\n", x)

	//Loops through the array comparing each number til the largest is found
  	for _,v:=range x {
		if v>temp {
    		fmt.Println(v,">",temp)
    		temp = v
    		large = temp
    	} else {
      	fmt.Println(v,"<",temp)
    	}
  	}

	fmt.Println("The largest number is ", large)

	//Loops through the array comparing each number til the smallest is found
	for _,v:=range x {
		if v>temp {
			fmt.Println(v,">",temp)
		} else {
			fmt.Println(v,"<",temp)
			temp = v
			small = temp
		}
	}

	fmt.Println("The smallest number is ", small)
}