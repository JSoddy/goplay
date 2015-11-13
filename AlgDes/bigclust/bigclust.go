package main 

import (
	"fmt"
	"UJS/file"
	"strconv"
	"strings"
)

const fileToken string = "d:\\James\\Temp\\clustering_big.txt"

func main() {
	// read the nodes in
	graph := readNodes()
	// count up clusters
	unionSlice := unions(graph)
	// output k-count
	fmt.Println(len(unionSlice))
}

func unions(graph [][]bool) [][][]bool {
	clustered := make([][][]bool, 0)
	for i := 0; len(graph) > 0; i++ {
		fmt.Println(i)
		graph, clustered = reUnion(graph, clustered)
	}
	return clustered
}

func reUnion(graph [][]bool, clustered[][][]bool) ([][]bool, [][][]bool) {
	newUnion := make([][]bool, 1)
	newUnion[0] = graph[0]
	graph = graph[1:]
	for j := 0; j < len(newUnion); j++ {
		for i := 0; i < len(graph); {
			if distance(newUnion[j], graph[i]) < 3 {
				newUnion = append(newUnion, graph[i])
				graph = append(graph[:i], graph[i+1:]...)
			} else {
				i++
			}
		}
	}
	clustered = append(clustered, newUnion)
	return graph, clustered
}

func distance(first []bool, second []bool) int {
	curDist := 0
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			curDist++
		}
	}
	return curDist
}

func readNodes() [][]bool {
	lines := file.FileLines(fileToken)
	params := strings.Split(strings.Split(lines[0], "\n")[0], " ")
	lines = lines[1:]
	count, _ := strconv.Atoi(params[0])
	size, _  := strconv.Atoi(params[1])
	graph := make([][]bool, count)
	for i := 0; i < count; i++ {
		graph[i] = newNode(lines[i], size)
	}
	return graph
}

func newNode(line string, size int) []bool {
	boolStrings := strings.Split(line, " ")
	node := make([]bool, size)
	for i := 0; i < len(node); i++ {
		node[i], _ = strconv.ParseBool(boolStrings[i])
	}
	return node
}