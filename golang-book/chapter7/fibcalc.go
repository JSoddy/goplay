package main

import "fmt"

func fibCalc(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	return fibCalc(x-1) + fibCalc(x-2)

}

func main() {
	sequence := uint(0)

	fmt.Print("Which Fibonacci number do you want? ")

	fmt.Scanf("%d", &sequence)

	fmt.Println()
	fmt.Println(fibCalc(sequence))
}
