package main

import "fmt"

func swapem(first *int, second *int) {
	temp := *first

	*first = *second
	*second = temp
}

func main() {
	var x, y int

	fmt.Print("Input a value for X: ")
	fmt.Scanf("%d\n", &x)
	fmt.Println()
	fmt.Println("X =", x)
	fmt.Println()
	fmt.Print("Input a value for Y: ")
	fmt.Scanf("%d\n", &y)
	fmt.Println()
	fmt.Println("Y =", y)
	fmt.Println()
	fmt.Println("Swapping...")
	fmt.Println()
	swapem(&x, &y)
	fmt.Println("X =", x, "Y =", y)
}
