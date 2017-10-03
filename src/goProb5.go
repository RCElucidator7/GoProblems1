package main
import(
    "fmt"
    "math/rand"
    "time"
)

//this generates random number between given range
func rando(min, max int) int {
    //Uses time to get a random number as it guarentees a random number every time
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := rando(1, 100)
    guessTaken := 0
    var guess int
    var lastGuess int = 0
	
    fmt.Printf("I am thinking of a number between 1 and 100.\n")
    
    //loops through 1-100 so the user gets 100 guesses
    for guessTaken < 100 {
        fmt.Println("Take a guess...")
        fmt.Scanf("%d", &guess)
        //Flushes the buffer
        fmt.Scanf("%d")
        guessTaken++
        //Checks if the user inputs the same number as they last entered
        if lastGuess == guess {
            fmt.Println("You cannot guess the same number!")
            guessTaken--
        }
        if guess < myrand {
            fmt.Println("Your guess is too low.")
        }
        if guess > myrand {
            fmt.Println("Your guess is too high.")
        }
        if guess == myrand {
            break
        }
        lastGuess = guess
    }
    if guess == myrand {
        fmt.Printf("You guessed my number in %d tries\n", guessTaken)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}