package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName string = "D:\\James\\Google Drive\\Coding\\go\\Data Files\\algo1-programming_prob-2sum.txt"

var iterations int = 0

var parVals = make([]bool, 20001)

func main() {
	// read integer array
	s := loadSlice(fileName)
	// load map
	m := loadMap(s)
	// search for pairs
	count := twoSumCount(s, m)
	// output
	fmt.Println(count)
}

func twoSumCount(s []int, m map[int]int) int {
	count := 0
	for i := -10000; i <= 10000; i++ {
		fmt.Println(iterations, " ", count)
		if pairInRange(i, s, m) {
			count++
		}
	}
	return count
}

func pairInRange(toMatch int, s []int, m map[int]int) bool {
	found := false
	for i := 0; i < len(s) && found == false; i++ {
		iterations++
		_, ok := m[toMatch-s[i]]
		found = ok
	}
	return found
}

func loadMap(s []int) map[int]int {
	m := make(map[int]int)
	for _, val := range s {
		m[val] = val
	}
	return m
}

func loadSlice(fileToken string) []int {
	var num int
	s := make([]int, 1000000)
	file, err := os.Open(fileToken)
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Attach file reader to buffered reader
	r := bufio.NewReader(file)
	// Read the first line into a string
	for i := 0; i < 1000000; i++ {
		fmt.Fscanln(r, &num)
		s[i] = num
	}
	return s
}
