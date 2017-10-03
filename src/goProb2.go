//Ryan Conway - G00332826

package main

import (
	"fmt"
	"time"
)

func main() {
	//Assigning the time.Now() to a variable to call hour/minutes/seconds
	curt := time.Now()
	fmt.Printf("The time is %d hours, %d minutes, and %d seconds", curt.Hour(), curt.Minute(), curt.Second())
}