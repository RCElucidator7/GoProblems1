//Ryan Conway - G00332826

package main

import (
	"fmt"
	"math"
)

//Mathematical equation used to determine a square root
func z_next(z float64, x float64) float64{
	return z - (((z * z) - x) / (2 * z))
}

func main() {
	x := 20.0
	
	var z float64
	
	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %12.8f\n", z)
	}

	//Prints the square root based off the mathematical equation
	fmt.Printf("The square root of %f is %f \n", x, z)
	//Prints the exact root using the math import
	fmt.Printf("The math.Sqrt Calculation is %f", math.Sqrt(x));
}