package main

import (
	"fmt"
	//	"io/ioutil"
	//	"log"
	"os"
)

const filename = "D:\\James\\Downloads\\IntegerArray.txt"
const arrsize = 100000

func main() {

	a := loadFromFile()
	i := mSort(a)
	//printIntArray(a)
	fmt.Println(i)
}

func mSort(a []int) (i int) {

	if len(a) < 2 {
		return 0
	}
	// call mSort on left
	i += mSort(a[0 : len(a)/2])
	// call mSort on right
	i += mSort(a[len(a)/2 : len(a)])
	// copy results into 2 new arrays
	l := make([]int, len(a)/2)
	r := make([]int, (len(a) - (len(a) / 2)))
	copy(l, a[0:len(a)/2])
	copy(r, a[len(a)/2:len(a)])
	i += mergeInts(a, l, r)
	return i
}

// merge l and r arrays into a: return the number of inversions
// between left and right
func mergeInts(a []int, l []int, r []int) (i int) {
	// create three zeroed indices
	i1, i2, i3 := 0, 0, 0
	// for incex 1 is in range l and index 2 is in range r
	for i1 < len(l) && i2 < len(r) {
		// if l[index1] is > r[index2]
		if l[i1] < r[i2] {
			// set a[index3] to l[index1]
			a[i3] = l[i1]
			// indrement index 1 and 3
			i1++
			i3++
		} else {
			// set a [index3] to r[index2]
			a[i3] = r[i2]
			// increment index 2 and 3
			i2++
			i3++
			// i += len(l) - index1
			i += (len(l) - i1)
		}
	}
	// for index1 < len(l)
	for i1 < len(l) {
		// set a[index3] to l[index1]
		a[i3] = l[i1]
		// increment index 1 and 3
		i1++
		i3++
	}
	// for index2 < len(r)
	for i2 < len(r) {
		// set a[index3] to r[index2]
		a[i3] = r[i2]
		// increment index 2 and 3
		i2++
		i3++
	}
	return i
}

func loadFromFile() []int {

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

/*
func printIntArray(a []int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
*/
