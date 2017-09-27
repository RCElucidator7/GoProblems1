package main
import(
    "fmt"
)

func main() {

	var i, j, k int = 0, 0, 0

	for i < 100{
		i++
		j++
		k++
		if j == 3 && k == 5 {	
			fmt.Print("FizzBuzz\n")
			j = 0
			k = 0
		} else if j == 3{
			fmt.Print("Fizz\n")
			j = 0
			if k == 5{
				fmt.Print("Buzz\n")
				k = 0
			} 
		} else if k == 5{
			fmt.Print("Buzz\n")
			k = 0
		} else{
			fmt.Printf("%d \n", i)
		}
	}
}