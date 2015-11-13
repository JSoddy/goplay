package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	n := 50 // The number of different values we will run
			// fibonacci sequence on for each method
	repetitions := 100 // Number of times we will repeat each fibonacci
	                   // calculation


	// Just create a 2D slice to hold time values
	var results = make([][]time.Duration, 3)

	for i := range results {
		results[i] = make([]time.Duration, n)
	}

	for i := range results[0] {
		// Run first method and average it's run times for values (i+1)*5

		if i < 7 { // We can't run the recursive function for large values
			
			// First run the function once to initialize everything
			_ = recursiveFib((i+1) * 5)

			// Then start the timer and run it repeatedly
			t0 := time.Now()

				for j := 0; j < repetitions; j++ {
					_ = recursiveFib((i+1) * 5)
				}

			t1 := time.Now()

		results[0][i] = t1.Sub(t0) // Store total run time in slice
		}

		// Run second method and average it's run times for values (i+1)*5
		
		// First run the function once to initialize everything
		_ = inductiveFib((i+1) * 5)

		// Then start the timer and run it repeatedly
		t0 := time.Now()

			for j := 0; j < repetitions; j++ {
				_ = inductiveFib((i+1) * 5)
			}

		t1 := time.Now()

		results[1][i] = t1.Sub(t0) // Store total run time in slice

		// Run third method and average it's run times for values (i+1)*5

		// First run the function once to initialize everything
		_ = formulaFib((i+1) * 5)

		// Then start the timer and run it repeatedly
		t0 = time.Now()

			for j := 0; j < repetitions; j++ {
				_ = formulaFib((i+1) * 5)
			}

		t1 = time.Now()

		results[2][i] = t1.Sub(t0) // Store total run time in slice
	}

	// Print out our slice of times
	fmt.Println(results)
}


// Function to compute the nth fibonacci number inductively
// by adding up to the given total
func inductiveFib(n int) int {
	fibonacciN2 := 1 // Will always hold n-2
	fibonacciN1 := 1 // Will always hold n-1
	fibonacciCurrent := 0 // Will always hold nth fibonacci sequence

	// First clear out base cases
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	// Otherwise calculate from there
	} else {
		// We'll start at 
		for i := 2; i <= n; i++ {
			// calculate the current fib number
			fibonacciCurrent = fibonacciN1 + fibonacciN2
			// Then update n-1 and n-2
			fibonacciN1 = fibonacciN2
			fibonacciN2 = fibonacciCurrent
		}
	}
	return fibonacciCurrent
}

// Function to compute the nth fibonacci number recursively
func recursiveFib(n int) int {
	if n == 0 {
		return 0; // Base case 1
	} else if n == 1 {
		return 1; // Base case 2
	} else {
			// Recurse case
		return recursiveFib(n - 1) + recursiveFib(n - 2)
	}
}

// Function to compute the nth fibonacci number by using the formula
func formulaFib(n int) int {
	sqrt5 := math.Sqrt(float64(5)) // No reason to repeatedly compute SQRT of 5
	
	// Just computes nth fibonacci and returns it
	return int((math.Pow((1 + sqrt5), float64(n)) - 
		math.Pow((1 - sqrt5), float64(n))) / 
		(math.Pow(float64(2), float64(n)) * sqrt5))
}