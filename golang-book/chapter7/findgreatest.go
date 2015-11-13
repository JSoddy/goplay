package main

import "fmt"

func findGreatest(list ...int) int {
	greatest := 0

	for _, n := range list {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

func main() {

	fmt.Println(findGreatest(8, 9, 99, 107, 2, 6, 4, 308, 0, 5, 1))

}
