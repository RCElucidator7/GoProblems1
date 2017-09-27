package main

import (
	"fmt"
)

func z_next(z float64, x float64) float64{
	return z - (((z * z) - x) / (2 * z))
}

func main() {
	x := 20.0
	
	var z float64
	
	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Printf("Current guess: %12.8f\n", z)
	}

	fmt.Printf("The square root of %f is %f \n", x, z)
	
	fmt.Printf("The math.sqrd Calculation is %f", Math.sqrd);
}