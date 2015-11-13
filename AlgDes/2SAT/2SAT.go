package main

import (
	"fmt"
	"UJSPackage/file"
	"sort"
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

type Vertices []*vertex

func (this Vertices) Len() int {
	return len(this)
}

func (this Vertices) Less(i, j int) bool {
	return this[i].pos < this[j].pos
}

func (this Vertices) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

const fileToken string = "D:\\James\\Temp\\2sat6.txt"

var currentLeader int = 0

func main() {
	// get the graph we will be working on as a pointer to type graph
	G := loadgraph()

	// run DFSLoop twice, first time with vertices reversed
	DFSLoop(G, true)
	clearExplored(G)
	DFSLoop(G, false)
	sort.Sort(Vertices(G.vertices))

	// get an array of the leaders of graph G
	l := findConflicts(G)
	//l := countLeaders(G)
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

func findConflicts(G *graph) bool {
	conflict := false
	spread		:= len(G.vertices) / 2
	for i := 0; i < spread && conflict == false ; i++ {
		conflict = G.vertices[i].leader == G.vertices[i+spread].leader
	}
	return !conflict
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
		if v.count > 0 {
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
func loadgraph() *graph {

	fileInts	:= file.IntLines(fileToken)
	size 		:= fileInts[0][0]
	fileInts	 = fileInts[1:]
	V 			:= fillVertices(size * 2)
	E 			:= fillEdges(size * 2)
	for i, v := range fileInts {
		if pos(v[0]) {
			if pos(v[1]) {
				E[i].tail = V[v[0]+size-1]
				E[i].head = V[v[1]-1]
				V[v[0]+size-1].edges = append(V[v[0]+size-1].edges, E[i])
				V[v[1]-1].edges = append(V[v[1]-1].edges, E[i])
				E[i+size].tail = V[v[1]+size-1]
				E[i+size].head = V[v[0]-1]
				V[v[1]+size-1].edges = append(V[v[1]+size-1].edges, E[i+size])
				V[v[0]-1].edges = append(V[v[0]-1].edges, E[i+size])
			} else {
				E[i].tail = V[v[0]+size-1]
				E[i].head = V[abs(v[1])+size-1]
				V[v[0]+size-1].edges = append(V[v[0]+size-1].edges, E[i])
				V[abs(v[1])+size-1].edges = append(V[abs(v[1])+size-1].edges, E[i])
				E[i+size].tail = V[abs(v[1])-1]
				E[i+size].head = V[v[0]-1]
				V[abs(v[1])-1].edges = append(V[abs(v[1])-1].edges, E[i+size])
				V[v[0]-1].edges = append(V[v[0]-1].edges, E[i+size])
			}
		} else {
			if pos(v[1]) {
				E[i].tail = V[abs(v[0])-1]
				E[i].head = V[v[1]-1]
				V[abs(v[0])-1].edges = append(V[abs(v[0])-1].edges, E[i])
				V[v[1]-1].edges = append(V[v[1]-1].edges, E[i])
				E[i+size].tail = V[v[1]+size-1]
				E[i+size].head = V[abs(v[0])+size-1]
				V[v[1]+size-1].edges = append(V[v[1]+size-1].edges, E[i+size])
				V[abs(v[0])+size-1].edges = append(V[abs(v[0])+size-1].edges, E[i+size])
			} else {
				E[i].tail = V[abs(v[0])-1]
				E[i].head = V[abs(v[1])+size-1]
				V[abs(v[0])-1].edges = append(V[abs(v[0])-1].edges, E[i])
				V[abs(v[1])+size-1].edges = append(V[abs(v[1])+size-1].edges, E[i])
				E[i+size].tail = V[abs(v[1])-1]
				E[i+size].head = V[abs(v[0])+size-1]
				V[abs(v[1])-1].edges = append(V[abs(v[1])-1].edges, E[i+size])
				V[abs(v[0])+size-1].edges = append(V[abs(v[0])+size-1].edges, E[i+size])
			}
		}
	}
	G 		:= new(graph)
	G.vertices 		= V
	G.edges 		= E
	return G
}

func fillVertices(size int) []*vertex {
	V 		:= make([]*vertex, size)
	for i, _ := range V {
		V[i] = new(vertex)
		V[i].pos 	= i + 1
		V[i].leader = i + 1
	}
	return V
}

func fillEdges(size int) []*edge {
	E 		:= make([]*edge, size)
	for i, _ := range E {
		E[i] = new(edge)
	}
	return E
}

func pos(this int) bool {
	return this > 0
}

func abs(this int) int {
	if this > 0 {
		return this
	} else {
		return -this
	}
}


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

