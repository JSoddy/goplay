package main

import "fmt"

func makeFibGenerator() func() uint64 {
	fib1 := uint64(0)
	fib2 := uint64(1)

	return func() (next uint64) {
		next = fib2
		fib2 += fib1
		fib1 = next
		return
	}
}

func main() {
	makeFib := makeFibGenerator()
	count := 0

	fmt.Print("How many Fibonacci numbers do you want? ")

	fmt.Scanf("%d", &count)

	fmt.Println()

	for i := 0; i < count; i++ {
		fmt.Println(makeFib())
	}
}
