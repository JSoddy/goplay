package main

import (
	"UJS/file"
	"UJS/heaptypes"
	"container/heap"
	"fmt"
)

const dataFile string = "D:\\James\\Google Drive\\coding\\go\\data files\\median.txt"

func main() {
	var n int
	runningTotal := 0
	hiHalf := new(heaptypes.IntHeaplow)
	loHalf := new(heaptypes.IntHeaphigh)

	lines := file.FileInts(dataFile)
	fmt.Println(len(lines))
	// initialize low and high heaps
	heap.Init(hiHalf)
	heap.Init(loHalf)
	// get our first data value
	currentMedian := lines[0]
	for _, v := range lines {
		// push the data onto a heap
		if v <= currentMedian {
			heap.Push(loHalf, v)
		} else {
			heap.Push(hiHalf, v)
		}
		// balance the arrays
		if len(*loHalf) > len(*hiHalf)+1 {
			n = heap.Pop(loHalf).(int)
			heap.Push(hiHalf, n)
		} else if len(*hiHalf) > len(*loHalf) {
			n = heap.Pop(hiHalf).(int)
			heap.Push(loHalf, n)
		}
		currentMedian = heap.Pop(loHalf).(int)
		runningTotal += currentMedian
		heap.Push(loHalf, currentMedian)
	}
	fmt.Println(heap.Pop(loHalf).(int))
	fmt.Println(heap.Pop(hiHalf).(int))
	fmt.Println(runningTotal)
	fmt.Println(runningTotal % 10000)
}
