package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var theGuess, guessCount uint
	theRand := uint(pickANumber())

	fmt.Println("Guess a number between 1 and 100.")
	guessCount++
	theGuess = getUserNumber()

	for theGuess != theRand {
		guessCount++
		if theGuess < theRand {
			fmt.Println("Guess again. That's too low.")
			theGuess = getUserNumber()
		} else if theGuess > theRand {
			fmt.Println("Guess again. That's too high.")
			theGuess = getUserNumber()
		}
	}

	fmt.Println("That's right! The number was", theRand)
	fmt.Println("You got it in", guessCount, "tries.")
}

func pickANumber() int {
	seedTime := time.Now()
	rand.Seed(int64(seedTime.Nanosecond()))
	return rand.Intn(99) + 1
}

func getUserNumber() uint {
	var discard string
	input := uint(0)
	fmt.Scanf("%d%s", &input, &discard)
	return input
}
