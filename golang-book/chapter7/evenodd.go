package main

import "fmt"

func half(aNumber int) (int, bool) {
	isEven := aNumber%2 == 0

	aNumber /= 2

	return aNumber, isEven
}

func main() {
	fmt.Println(half(18))
	fmt.Println(half(99))
}
