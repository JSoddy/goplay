package main

import "fmt"

func swapem(first *string, second *string) {
	temp := *first

	*first = *second
	*second = temp
}

func main() {
	var x, y string

	fmt.Print("Input a value for M: ")
	fmt.Scanf("%s\n", &x)
	fmt.Println()
	fmt.Println("M =", x)
	fmt.Println()
	fmt.Print("Input a value for B: ")
	fmt.Scanf("%s\n", &y)
	fmt.Println()
	fmt.Println("B =", y)
	fmt.Println()
	fmt.Println("Swapping...")
	fmt.Println()
	swapem(&x, &y)
	fmt.Println("M =", x, "B =", y)
}
