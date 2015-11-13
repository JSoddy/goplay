package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

const fileName string = "D:\\James\\Google Drive\\Coding\\go\\Data Files\\algo1-programming_prob-2sum.txt"

var iterations int = 0

var pairVals = make([]bool, 20001)

func main() {
	runtime.GOMAXPROCS(8)
	// read integer array
	s := loadSlice(fileName)
	// load map
	m := loadMap(s)
	// search for pairs
	twoSumCount(s, m)
	// output
	count := pairValsSort()
	fmt.Println(count)
}

func twoSumCount(s []int, m map[int]int) {
	done := make(chan bool)
	openGos := 0
	for i := -10000; i <= 10000; i++ {
		go pairInRange(i, s, m, done)
		openGos++
		for openGos > 200 {
			<-done
			fmt.Println(iterations)
			openGos--
		}
		time.Sleep(5 * time.Millisecond)
	}
	for openGos > 0 {
		fmt.Println(iterations)
		<-done
		openGos--
	}
}

func pairInRange(toMatch int, s []int, m map[int]int, done chan bool) {
	for i := 0; i < len(s); i++ {
		iterations++
		_, ok := m[toMatch-s[i]]
		if ok {
			pairVals[toMatch+10000] = true
		}
	}
	done <- true
}

func pairValsSort() int {
	count := 0
	for _, v := range pairVals {
		if v {
			count++
		}
	}
	return count
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
