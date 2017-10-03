//Ryan Conway - G00332826

package main
import(
    "fmt"
)

func main() {

	var i int = 0

	for i < 100{
		i++
		//i goes from 1 to 100, everytime it hits 3 and 5 it will print Fizzbuzz
		if i % 3 == 0 && i % 5 == 0{	
			fmt.Print("FizzBuzz\n")
		} else if i % 3 == 0{//every time i hits 3 it prints Fizz
			fmt.Print("Fizz\n")
		} else if i % 5 == 0{//every time i hits 5 it prints Buzz
			fmt.Print("Buzz\n")
		} else{//prints the number
			fmt.Printf("%d \n", i)
		}
	}
}