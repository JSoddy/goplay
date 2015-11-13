package main 

import (
	"fmt"
	"UJS/file"
	"container/heap"
	"strings"
	"strconv"
)

// Node is an element of a Graph
type Node struct {
	name     int  // The node name (numbers)
	lowCost	 int  // Lost Cost edge to Node from spanning tree
	explored bool // True if Node has been explored
	index    int  // index of the item in the heap
	edges    []Edge
}

// Edge is an element of a Graph
type Edge struct {
	cost  int   // The length of this edge
	vert1   *Node
	vert2	*Node
}

// Graph implements heap.Interface and holds Nodes
type Graph []*Node

func (g Graph) Len() int { return len(g) }

func (g Graph) Less(i, j int) bool {
	return g[i].lowCost < g[j].lowCost
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

func (g *Graph) update(n *Node, lowCost int) {
	if n.lowCost > lowCost {
		heap.Remove(g, n.index)
		n.lowCost = lowCost
		heap.Push(g, n)
	}
}

const fileName = "d:\\James\\Temp\\edges.txt"

func main() {
	mainGraph := loadGraph()

	treeCost := span(mainGraph)

	fmt.Println(treeCost)
}

func span(mainGraph *Graph) int {
	graphCopy := &Graph{}
	heap.Init(graphCopy)
	for _, v := range *mainGraph {
		heap.Push(graphCopy, v)
	}
	treeCost := 0
	for len(*mainGraph) > 0 {
		node := heap.Pop(mainGraph).(*Node)
		node.explored = true
		for _, v := range node.edges {
			if v.vert1.explored == false {
				mainGraph.update(v.vert1, v.cost)
			}
			if v.vert2.explored == false {
				mainGraph.update(v.vert2, v.cost)
			}
		}
		treeCost += node.lowCost
	}
	return treeCost - 100000
}


func loadGraph() *Graph {
	fileStrings := file.FileLines(fileName)

	countStrings := strings.Split(strings.TrimSuffix(fileStrings[0], "\n"), " ")
	nodeCount, _ := strconv.Atoi(countStrings[0])
	graphNodes := make(Graph, nodeCount)

	fileStrings = fileStrings[1:]
	initializeNodes(&graphNodes)

	for _, v := range fileStrings {
		addEdge(graphNodes, v)
	}

	return &graphNodes
}

func addEdge(graphNodes Graph, edgeString string) {
	edgeSlice := strings.Split(strings.TrimSuffix(edgeString, "\n"), " ")
	edgeCost, _ := strconv.Atoi(edgeSlice[2])
	firstVert, _ := strconv.Atoi(edgeSlice[0])
	secondVert, _ := strconv.Atoi(edgeSlice[1])
	newEdge := Edge{edgeCost, graphNodes[firstVert-1], graphNodes[secondVert-1]}
	graphNodes[firstVert-1].edges = append(graphNodes[firstVert-1].edges, newEdge)
	graphNodes[secondVert-1].edges = append(graphNodes[secondVert-1].edges, newEdge)
}

func initializeNodes(graphNodes *Graph) {
	for i, _ := range *graphNodes {
		 (*graphNodes)[i] = new(Node)
		 (*graphNodes)[i].name = i+1
		 (*graphNodes)[i].lowCost = 100000
	}
}