package main

import "fmt"

func makeOddGenerator() func() uint {
	odds := uint(1)

	return func() (ret uint) {
		ret = odds
		odds += 2
		return
	}
}

func main() {
	makeOdd := makeOddGenerator()
	count := 0

	fmt.Println("How many odd numbers do you want?")

	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		fmt.Println(makeOdd())
	}
}
