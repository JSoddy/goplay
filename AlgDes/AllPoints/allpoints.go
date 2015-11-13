package main 

import (
	"fmt"
	"UJS/packages/graph"
	"UJS/file"
	"strings"
	"strconv"
	//"math"
)

const fileToken string = "d:\\James\\Temp\\g3.txt"

func main() {
	//load file
	graph := loadGraph()
	//perform graph search
	results := johnson(graph)
	//sort through results
	output := smallest(results)
	//output results
	fmt.Println(output)
}

func johnson(g graph.Graph) [][]int {
	output 			:= make([][]int, g.Len())
	//run Bellman-Ford
	fmt.Println("Running Bellman-Ford...")
	lengths, err	:= g.BellmanFord(1)
	if err != 0 {
		fmt.Println("Shooooot, that graph has cycles.")
	} else {


		//reweight
		fmt.Println("Reweighting...")
		g = reweight(g, lengths)


		fmt.Println("Running Dijkstra's...")
		for i, _ := range g.Vertices {
			// restore Vertex ordering
			tempVertices := make([]*graph.Node, g.Len())
			for _, v := range g.Vertices {
				tempVertices[v.Name - 1] = v
			}
			g.Vertices = tempVertices
			output[i] = g.Dijkstra(i+1)
		}

		//unreweight
		fmt.Println("Removing reweighting...")
		output = unreweight(output, lengths)
		//return values
		//placeholder 
	}
	return output
}

func smallest(results [][]int) int {
	currentSmall := 0x3f3f3f3f
	for i, v := range results {
		for j, w := range v {
			if i != j {
				if w < currentSmall {
					currentSmall = w
				}
			}
		}
	}
	return currentSmall
}

func reweight(g graph.Graph, weights []int) graph.Graph {
	for j, w := range g.Vertices {
		for i, v := range w.EdgesOut {
			g.Vertices[j].EdgesOut[i].Length += weights[v.Tail.Name-1] - weights[v.Head.Name-1]
		}
	}
	return g
}

func unreweight(paths [][]int, weights []int) [][]int {
	for i, v := range paths {
		for j, _ := range v {
			paths[i][j] = paths[i][j] + weights[j] - weights[i]
		}
	}
	return paths
}

// () -> Graph
// Outputs a graph containing the contents of file "fileToken"
func loadGraph() graph.Graph {
	fmt.Println("Opening file...")
	rawLines	:= file.FileLines(fileToken)
	fmt.Println("Creating structures...")
	values		:= strings.Split(rawLines[0], "\n")
	values		=  strings.Split(values[0], " ")
	vCount, _	:=  strconv.Atoi(values[0])
	eCount, _	:=  strconv.Atoi(values[1])
	rawLines	=  rawLines[1:]
	vertices 	:= make([]*graph.Node, vCount)
	edges		:= make([]graph.Edge, eCount)
	for i := 0; i < len(vertices); i++ {
		vertices[i] = &graph.Node{Name: i+1}
	}
	fmt.Println("Entering data...")
	for i, v := range rawLines {
		edges[i]	= addEdge(vertices, v)
	}
	return graph.Graph{Vertices: vertices, Edges: edges}
}

func addEdge(vertices []*graph.Node, line string) graph.Edge {
	values			:= strings.Split(line, "\n")
	values			= strings.Split(values[0], " ")
	tailName, _		:= strconv.Atoi(values[0])
	headName, _		:= strconv.Atoi(values[1])
	length, _		:= strconv.Atoi(values[2])
	tail 			:= vertices[tailName - 1]
	head 			:= vertices[headName - 1]
	edge 			:= graph.Edge{Tail: tail, Head: head, Length: length}
	tail.EdgesOut 	= append(tail.EdgesOut, edge)
	head.EdgesIn 	= append(head.EdgesIn, edge)
	return edge
}