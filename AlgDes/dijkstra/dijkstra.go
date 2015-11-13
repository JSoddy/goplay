package main

import (
	"UJS/file"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	//"math"
)

// Node is an element of a Graph
type Node struct {
	name     int  // The node name (numbers)
	distance int  // Shortest path distance to Node
	explored bool // True if Node has been explored
	index    int  // index of the item in the heap
	edges    []Edge
}

// Edge is an element of a Graph
type Edge struct {
	length int   // The length of this edge
	node   *Node // The node at this Edge's head
}

// Graph implements heap.Interface and holds Nodes
type Graph []*Node

func (g Graph) Len() int { return len(g) }

func (g Graph) Less(i, j int) bool {
	return g[i].distance < g[j].distance
}

func (g Graph) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
	g[i].index = i
	g[j].index = j
}

func (g *Graph) Push(x interface{}) {
	n := len(*g)
	node := x.(*Node)
	node.index = n
	*g = append(*g, node)
}

func (g *Graph) Pop() interface{} {
	old := *g
	n := len(old)
	node := old[n-1]
	node.index = -1
	*g = old[0 : n-1]
	return node
}

func (g *Graph) update(n *Node, distance int) {
	if n.distance > distance {
		heap.Remove(g, n.index)
		n.distance = distance
		heap.Push(g, n)
	}
}

const fileToken = "D:\\James\\Google Drive\\Coding\\go\\Data Files\\dijkstraData.txt"

var returnDistances = [...]int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}

func main() {
	// load data into array - all distances are 10000 - all nodes 'unseen'
	g := loadGraph()
	// pass array and element #1 to djikstra loop function
	outputThis	:= g.dijkstra(1)
	// print out the min distances to nodes in returnDistances
	//printDistances(g)
	fmt.Println(outputThis)
}

// Graph Int -> ()
// computes the shortest path to all elements in Graph g from node n
func (g Graph) dijkstra(n int) []int{
	maxIndex	:= len(g)
	if maxIndex < 1 {
		return make([]int, 0)
	}//set all Node distance to maxInt
	for i := 0; i < len(g); i++ {
		g[i].distance	= 0x3f3f3f3f
		g[i].index		= i
	}
	//initialize heap
	heap.Init(&g)
	//set starting Node distance to 0
	g.update(g[n-1], 0)
	heap.Init(&g)
	for len(g) > 1 {
		node := heap.Pop(&g).(*Node)
		node.explored = true
		for _, edge := range node.edges {
			if !edge.node.explored {
				g.update(edge.node, (edge.length + node.distance))
			}
		}
	}
	g 	= g[:maxIndex]
	values		:= make([]int, len(g))
	for _, v := range g{
		values[v.name-1]	= v.distance
	}
	return values
}

/*
// Graph Int -> ()
// computes the shortest path to all elements in Graph g from node n
func (g Graph) dijkstra(n int) {
	if len(g) < 1 {
		return
	}
	//set element #1 distance to 0
	g[n-1].distance = 0
	heap.Init(&g)
	for len(g) > 1 {
		node := heap.Pop(&g).(*Node)
		node.explored = true
		for _, edge := range node.edges {
			if !edge.node.explored {
				g.update(edge.node, (edge.length + node.distance))
			}
		}
	}
	g 	= g[:cap(g)]
	return
}

// Graph Int -> ()
// computes the shortest path to all elements in Graph g from node n
func djikstra(g Graph, n int) {
	//set element #1 distance to 0
	if len(g) < 1 {
		return
	}
	g[n-1].distance = 0
	h := &Graph{}
	heap.Init(h)
	for _, v := range g {
		heap.Push(h, v)
	}
	heap.Init(h)
	for len(*h) > 1 {
		node := heap.Pop(h).(*Node)
		node.explored = true
		for _, edge := range node.edges {
			if !edge.node.explored {
				h.update(edge.node, (edge.length + node.distance))
			}
		}
	}
	return
}
*/

// Graph -> ()
// Outputs the shortest path to elements in Graph which are members of
// const returnDistances
func printDistances(g Graph) {
	l := len(g)
	for _, n := range returnDistances {
		if n <= l {
			fmt.Println(g[n-1].name, " ", g[n-1].distance)
		} else {
			fmt.Println("index: ", n, " out of range!")
		}
	}
}

// () -> Graph
// Outputs a graph containing the contents of file "fileToken"
func loadGraph() Graph {

	a := file.FileLines(fileToken)
	g := make(Graph, 0)
	for i := 0; i < len(a); i++ {
		g = append(g, new(Node))
	}
	for _, l := range a {
		addNode(g, l)
	}
	return g
}

func addNode(g Graph, line string) {
	parts := strings.Split(line, "\t")
	name, _ := strconv.Atoi(parts[0])
	edges := make([]Edge, 0)
	for i := 1; i < len(parts)-1; i++ {
		edgepart := strings.Split(parts[i], ",")
		node, _ := strconv.Atoi(edgepart[0])
		length, _ := strconv.Atoi(edgepart[1])
		e := Edge{length, g[node-1]}
		edges = append(edges, e)
	}
	v := Node{name, 1000000, false, 0, edges}
	*g[name-1] = v
}
