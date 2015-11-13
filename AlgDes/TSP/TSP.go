package main 

import (
	"fmt"
	"strconv"
	"strings"
	"UJSPackage/file"
	"math"
)

type Edge struct {
	head int
	distance float64
}

type Vertex struct {
	xPos float64
	yPos float64
	edges []Edge 
}

type Graph []Vertex

const Inf float64 = float64(0x3f3f3f3f)
const fileToken string = "d:\\James\\Temp\\tsp.txt"

func main() {
	g := loadGraph()
	val := TSP(g)
	fmt.Println(math.Floor(val))
}

//!!!
func TSP(graph Graph) float64 {
	lookups 	:= makeTable(graph)
	ansSlice	:= primeGraph(lookups)
	for i := 2; i < len(graph); i++ {
		fmt.Println(i, "...")
		for index := makeSet(i); index < int(math.Pow(2.0, float64(len(graph) - 1))); index = gospers(index) {
			ansSlice[index - 1] = computeSet(index, ansSlice, lookups)
		}
	}
	return bestDistances(1, int(math.Pow(2.0, float64(len(graph)-1)))-1, ansSlice, lookups)
}

func computeSet(setInt int, ansSlice [][]float64, lookups [][]float64) []float64 {
	distances	:= make([]float64, 0)
	fullSet		:= IntToS(setInt)
	for i, _ := range fullSet {
		this, rest 	:= subSet(fullSet, i)
		bestDistance := bestDistances(this, rest, ansSlice, lookups)
		distances	= append(distances, bestDistance)
	}
	return distances
}

func bestDistances(this int, rest int, ansSlice [][]float64, lookups [][]float64) float64 {
	subSet		:= IntToS(rest)
	curMin		:= Inf
	for i, v 	:= range ansSlice[rest-1] {
		thisDist := v + lookups[subSet[i]-1][this-1]
		if thisDist < curMin {
			curMin = thisDist
		}
	}
	return curMin
}

func subSet(fullSet []int, index int) (int, int) {
	tempSet := make([]int, len(fullSet))
	for i, _ := range tempSet {
		tempSet[i] = fullSet[i]
	}
	currNode	:= tempSet[index]
	paredSet	:= SToInt(append(tempSet[:index],tempSet[index+1:]...))
	return currNode, paredSet
}

func primeGraph(table [][]float64) [][]float64 {
	graphSlice	:= make([][]float64, int(math.Pow(2.0, float64(len(table) - 1))) -1)
	index		:= makeSet(1)
	for i := 1; i < len(table); i++ {
		graphSlice[index - 1]	= []float64{table[0][i]}
		index	= gospers(index)
	}
	return graphSlice
}

func makeTable(graph Graph) [][]float64 {
	table := make([][]float64, len(graph))
	for i, _ := range table {
		table[i] = make([]float64, len(graph))
	}
	for i, v := range graph {
		for _, w := range v.edges {
			table[i][w.head] = w.distance
		}
	}
	return table
}

func IntToS(setInt int) []int {
	set := *new([]int)

	for j := 0; j < 25; j++ {
		if v := int(math.Pow(float64(2), float64(j))); setInt & v != 0 {
			set = append(set, j + 2)
		}
	}
	return set
}


func SToInt(set []int) int {

	setInt := 0

	//set our goal int & set correct number of bits in current
	for _,v := range set {
	
		setInt = setInt | int(math.Pow(float64(2), float64(v-2)))
		
	}
	return setInt
}

func gospers(from int) int {
	c := (from & -from)
 	r := from + c
	return (((r ^ from) >> 2) / c) | r
}

func nChoosek(n, k int) int {
	if k == 0 || k == n {
		return 1
	} else {
		return nChoosek(n-1, k-1) + nChoosek(n-1, k)
	}
}

func nChooseAdder(n int) int {
	var total int
	for i := 1; i <= n; i++ {
		total += nChoosek(n, i)
	}
	return total
}

func timeTest(dataSize int) (int, [][]int) {
	a := 0
	arraySets := make([][]int, int(math.Pow(float64(2), float64(dataSize))) - 1)
	for i, _ := range arraySets {
		arraySets[i] = IntToS(i+1)
		a = SToInt(arraySets[i])
	}
	return a, arraySets
}

// returns an integer with the first "size" bits set to 1
func makeSet(size int) int {
	set := 0
	for i := 0; i < size; i++ {
		set = set | int(math.Pow(float64(2), float64(i)))
	}
	return set
}

func loadGraph() Graph {
	fileLines 		:= file.FileLines(fileToken)
	dataLine		:= strings.Split(fileLines[0], "\n")
	dataLine		=  strings.Split(dataLine[0], "\r")
	nodeCount, _ 	:= strconv.Atoi(dataLine[0])
	fileLines 		=  fileLines[1:]

	graph := make(Graph, nodeCount)

	for i, v := range fileLines {
		graph[i] = readNode(v)
	}

	graph = calculateEdges(graph)

	return graph
}

func readNode(line string) Vertex {
	lines 	:= strings.Split(line, "\n")
	lines 	= strings.Split(lines[0], "\r")
	lines 	= strings.Split(lines[0], " ")
	xPos, _	:= strconv.ParseFloat(lines[0], 64)
	yPos, _ := strconv.ParseFloat(lines[1], 64)
	return Vertex{xPos: xPos, yPos: yPos, edges: make([]Edge, 0)}
}

func calculateEdges(graph Graph) Graph {
	for i, v := range graph {
		for j, w := range graph {
			if j != i {
				edgeLength := math.Sqrt(math.Pow(v.yPos - w.yPos, 2) + 
										math.Pow(v.xPos - w.xPos, 2))
				graph[i].edges = append(graph[i].edges, Edge{j, edgeLength})
			}
		}
	}
	return graph
}