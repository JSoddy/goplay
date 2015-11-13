package main 

import (
	"fmt"
	"UJS/file"
	"strings"
	"strconv"
	"math"
)

type item struct {
	value int
	weight int
}

const fileToken string = "d:\\James\\Temp\\Knapsack_big.txt"

func main() {
	// load dataset
	itemList, knapSize := loadArray()
	// run algorithm
	finalValue := knapsackLoops(itemList, knapSize)
	// print results
	fmt.Println(finalValue)
}

func knapsackLoops(itemList []item, knapSize int) int {
	solver 		:= newSolver(itemList, knapSize)
	indexToggle := 0
	for i := 1; i < len(itemList)+1; i++ {
		indexToggle		= i % 2
		for j := 1; j < len(solver[indexToggle]); j++ {
			if j < itemList[i-1].weight {
				solver[indexToggle][j]	= solver[(i-1)%2][j]
			} else {
				solver[indexToggle][j]	= int(math.Max(float64(solver[(i-1)%2][j]), 
									float64(solver[(i-1)%2][j-itemList[i-1].weight] + itemList[i-1].value)))
			}
		}
	}
	return solver[len(solver)-1][len(solver[len(solver)-1])-1]
}

func newSolver(itemList []item, knapSize int) [][]int {
	solver 		:= make([][]int, 2)
	for i, _ 	:= range solver {
		solver[i] 		= make([]int, knapSize+1)
		solver[i][0]	= 0
	}
	for i, _ 	:= range solver[0] {
		solver[0][i]	= 0
	}
	return solver
}

func loadArray() ([]item, int) {
	fileLines 			:= file.FileLines(fileToken)
	paramLine	 		:= strings.Split(fileLines[0], "\n")
	fileLines 			= fileLines[1:]
	paramLine 			= strings.Split(paramLine[0], " ")
	knapSize, _			:= strconv.Atoi(paramLine[0])

	itemList 			:= readKnap(fileLines)
	return itemList, knapSize
}

func readKnap(fileLines []string) []item {
	if len(fileLines) == 0 {
		return make([]item, 0)
	} else {
		line 			:= strings.Split(fileLines[0], "\n")
		line 	 		= strings.Split(line[0], " ")
		value, _		:= strconv.Atoi(line[0])
		weight, _		:= strconv.Atoi(line[1])
		fileLines 		= append(fileLines[:0], fileLines[1:]...)
		return append(readKnap(fileLines), item{value:value, weight:weight})
	}
}