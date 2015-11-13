package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type vertex struct {
	pos      int
	explored bool
	leader   int
	edges    []*edge
}

type edge struct {
	tail *vertex
	head *vertex
}

type graph struct {
	vertices []*vertex
	edges    []*edge
}

type leaders struct {
	name  int
	count int
}

const filename string = "D:\\James\\Google Drive\\Coding\\go\\Data Files\\SCC.txt"

const arrLen int = 875714

var currentLeader int = 0

func main() {
	// get the graph we will be working on as a pointer to type graph
	G := loadgraph(filename)

	// run DFSLoop twice, first time with vertices reversed
	DFSLoop(G, true)
	clearExplored(G)
	DFSLoop(G, false)

	// get an array of the leaders of graph G
	l := countLeaders(G)
	fmt.Println(l)
}

// Perform a depth first searches on G until all vertices are explored
// reorder the nodes into the reverse order in which they are fully explored
// assign a 'leader' to each vertex
// if rev = true, treat tail of all edges as heads <> heads as tails
func DFSLoop(G *graph, rev bool) {
	H := new(graph)
	currentLeader = 0
	for i := len(G.vertices) - 1; i >= 0; i-- {
		v := G.vertices[i]
		if !v.explored {
			currentLeader = v.pos
			DFSearch(H, v, rev)
		}
	}
	G.vertices = H.vertices
	return
}

func clearExplored(G *graph) {
	for _, v := range G.vertices {
		v.explored = false
	}
}

func DFSearch(G *graph, v *vertex, rev bool) {
	v.explored = true
	v.leader = currentLeader
	for _, w := range v.edges {
		if rev {
			if !w.tail.explored {
				DFSearch(G, w.tail, rev)
			}
		} else {
			if !w.head.explored {
				DFSearch(G, w.head, rev)
			}
		}
	}
	G.vertices = append(G.vertices, v)
}

// Tallies the leaders of all vertices in the given graph, returns
// the information as a slice of leaders
func countLeaders(G *graph) []leaders {
	var aL = make([]leaders, len(G.vertices))
	for _, v := range G.vertices {
		aL[v.leader-1].count++
	}
	aL = collapseAL(aL)
	return aL
}

// Increments the leader with value n in array aL,
// creates a leader in the appropriate position if it does not exist
func collapseAL(aL []leaders) []leaders {
	naL := make([]leaders, 0)
	for i, v := range aL {
		if v.count > 200 {
			v.name = i + 1
			naL = append(naL, v)
		}
	}
	return naL
}

// open and read a file containing graph information in the form of
// two integers per line, the first being the tail of an edge and the
// second being the head of that edge.
// Load the requested information into a graph data structure and return
// a pointer to it
// !!!
func loadgraph(filename string) *graph {

	fmt.Println("Open")
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fmt.Println("Split")
	lines := strings.Split(string(b), "\n")
	fmt.Println("Load")
	V := make([]*vertex, arrLen)
	E := make([]*edge, 0)

	for _, v := range lines {
		head, tail := 0, 0
		fmt.Sscanf(v, "%d %d", &tail, &head)
		V, E = addEdge(V, E, head, tail)
	}
	fmt.Println("Make G")
	G := new(graph)
	G.vertices = V
	G.edges = E
	fmt.Println("Return")
	return G
}

// adds an edge with given head and tail to []*edge E
// adds pointers to it to vertex head and vertex tail
// if vertices do not exist they will be created
// !!!
func addEdge(V []*vertex, E []*edge, head int, tail int) ([]*vertex, []*edge) {
	V, headVert := findVert(V, head)
	V, tailVert := findVert(V, tail)
	e := &edge{tailVert, headVert}
	headVert.edges = append(headVert.edges, e)
	tailVert.edges = append(tailVert.edges, e)
	return V, append(E, e)
}

func findVert(V []*vertex, pos int) ([]*vertex, *vertex) {
	if V[pos-1] == nil {
		V[pos-1] = new(vertex)
	}
	V[pos-1].pos = pos
	return V, V[pos-1]
}

/*
func printgraph(g *graph) {
	fmt.Println("Edges:")
	for _, v := range g.edges {
		fmt.Print(" ", v.tail.pos, "-", v.head.pos, " |")
	}
	fmt.Print("\n")
	fmt.Println("Vertices:")
	for _, v := range g.vertices {
		fmt.Print(v.pos, ":")
		for _, w := range v.edges {
			fmt.Print(" ", w.tail.pos, " ", w.head.pos, " |")
		}
		fmt.Print("\n")
	}
}
*/
