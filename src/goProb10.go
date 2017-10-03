//Ryan Conway - G00332826

package main 
import "fmt"

func main() { 
	
        input := "knalB"
        n := 0
        fmt.Println("Please input a word...")
        //If the user inputs a sentence it will only reverse the first word
        fmt.Scanf("%s", &input)
        // Rune turns the string letters into unicode 
        // Getting the unicode code points.
        rune := make([]rune, len(input))
        for _, r := range input { 
                rune[n] = r
                n++
        } 
        rune = rune[0:n]
        // Reverses the numbers to read backwards
        for i := 0; i < n/2; i++ { 
                rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
        } 
        // Convert back to interger values.
        output := string(rune)
        fmt.Println(output)
}