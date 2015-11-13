package main

import (
	"fmt"
	"math"
	"os"
)

const arrsize int = 100000
const filename string = "D:\\James\\Downloads\\IntegerArray.txt"

// Main:
func main() {
	a := readArray()
	count := qscount(a)
	//printarray(a)
	fmt.Println(count)
}

// qscount: (a)
func qscount(a []int) int {
	count := 0
	// if not base case (array length 1 or 0)
	if len(a) > 1 {
		count = (len(a) - 1)
		p := chpivot(a)
		p = partition(a, p)
		count += qscount(a[0:p])
		count += qscount(a[p+1 : len(a)])
	}
	return count
}

//  chpivot: (a)
func chpivot(a []int) int {
	i1, i2, i3, t := 0, int(math.Ceil(float64(len(a))/2))-1, len(a)-1, 0

	if a[i1] > a[i2] {
		t = i1
		i1 = i2
		i2 = t
	}
	if a[i2] > a[i3] {
		t = i2
		i2 = i3
		i3 = t
	}
	if a[i1] > a[i2] {
		t = i1
		i1 = i2
		i2 = t
	}
	return i2
}

//  partition (a, p)
func partition(a []int, p int) int {
	i := 0
	swap(a, p, 0)
	//  for j in [1 to length[a]]
	for j, v := range a {
		//  	if [j] < [0]
		if v < a[0] {
			i++
			swap(a, j, i)
		}
	}
	swap(a, 0, i)
	return i
}

// swap (a, i1, i2)
func swap(a []int, i1 int, i2 int) {
	j := a[i1]
	a[i1] = a[i2]
	a[i2] = j
}

/*
//	printarray: (a)
func printarray(a []int) {
//  for i in range (a)
	c := 0
	for i, v := range a {
//		printline a[i]
		fmt.Println(v)
		c += i
	}
}
*/
func readArray() []int {

	var i int
	a := make([]int, arrsize)
	// open file and create reader
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// for not EOF and not array is full
	for t := range a {
		fmt.Fscanln(file, &i)
		a[t] = i
	}
	// fscanln
	return a
}
