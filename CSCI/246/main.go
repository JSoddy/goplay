package main

/* Simple program to run Euclid's algorithm on 10 
 * sets of randomly chosen integers between 1 and 500000
 */

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intA, intB uint

	for i := 0; i < 10; i++ {
		intA = uint(pickANumber())
		intB = uint(pickANumber())

		fmt.Printf("The GCD of %6d and %6d is %d.\n", intA, intB, euclidIteration(intA, intB))
	}
}

// Gives a random number
func pickANumber() int {
	seedTime := time.Now()
	rand.Seed(int64(seedTime.Nanosecond()))
	return rand.Intn(50) + 1
}

// Performs Euclid's algorithm recursively
func euclidIteration(A, B uint) uint {
	if A < B {
		return euclidIteration(B, A)
	} else if B == 0 {
		return A
	} else {
		return euclidIteration(B, A % B)
	}
}
