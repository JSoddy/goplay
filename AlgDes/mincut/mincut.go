package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// type vertex is struct {pos int, edges *[]edge}
type vertex struct {
	pos   int
	edges []*edge
}

// type edge is struct {first *vertex, second *vertex}
type edge struct {
	first  *vertex
	second *vertex
}

// type graph is struct {vertices []vertex, edges []edge}
type graph struct {
	vertices []*vertex
	edges    []*edge
}

const filename string = "D:\\James\\Downloads\\kargermincut.txt"

func main() {
	aa := loadgraph(filename)
	g := atog(aa)
	reps := int(math.Pow(float64(len(g.vertices)), 2) + math.Log2(float64(len(g.vertices))))
	minCut := reduce(g)
	for i := 0; i < reps; i++ {
		g = atog(aa)
		currentCut := reduce(g)
		if currentCut < minCut {
			minCut = currentCut
		}
	}
	fmt.Println("Min cut = ", minCut)
}

// Runs Karger contraction algorithm on the provided graph
// returns the number of edges left as mincut
func reduce(g *graph) int {
	// seed random
	seedTime := time.Now()
	rand.Seed(int64(seedTime.Nanosecond()))
	// choose an edge to contract
	for len(g.vertices) > 2 {
		n := rand.Intn(len(g.edges))
		// identify the vertices at the ends of edge
		vert1, vert2 := vertsFromEdge(g, n)
		// delete all edges connecting these two vertices
		deleteEdges(g, vert1, vert2)
		// copy all edges from vert2 to vert1
		copyEdges(g, vert1, vert2)
		// change all edges pointing to second of vertex to point to first instead
		redirect(g, vert1, vert2)
		// remove second from the list of vertices
		rmVertex(g, vert2)
	}
	return len(g.edges)
}

// remove vertex indicated by vert2 from graph g
func rmVertex(g *graph, vert2 int) {
	l := 0
	for g.vertices[l].pos != vert2 {
		l++
	}
	tempSlice := make([]*vertex, len(g.vertices)-1)
	copy(tempSlice, g.vertices[:l])
	if l < len(g.vertices)-1 {
		copy(tempSlice[l:], g.vertices[l+1:])
	}
	g.vertices = tempSlice
}

func redirect(g *graph, vert1, vert2 int) {
	k, l := 0, 0
	for g.vertices[k].pos != vert1 {
		k++
	}
	for g.vertices[l].pos != vert2 {
		l++
	}
	for i := 0; i < len(g.vertices[l].edges); i++ {
		if g.vertices[l].edges[i].first == g.vertices[l] {
			g.vertices[l].edges[i].first = g.vertices[k]
		}
		if g.vertices[l].edges[i].second == g.vertices[l] {
			g.vertices[l].edges[i].second = g.vertices[k]
		}
	}
}

// append the list of edges from vertex matching vert2 to
//  vertex matching vert1 in provided graph
func copyEdges(g *graph, vert1, vert2 int) {
	k, l := 0, 0
	for g.vertices[k].pos != vert1 {
		k++
	}
	for g.vertices[l].pos != vert2 {
		l++
	}
	if len(g.vertices[l].edges) != 0 {
		tempSlice := make([]*edge, len(g.vertices[k].edges)+len(g.vertices[l].edges))
		copy(tempSlice, g.vertices[k].edges)
		copy(tempSlice[len(g.vertices[k].edges):], g.vertices[l].edges)
		g.vertices[k].edges = tempSlice
	}
}

// Return the pos of the vertices at each end of edge
// in provided graph represented by provided int
func vertsFromEdge(g *graph, n int) (vert1, vert2 int) {
	return g.edges[n].first.pos, g.edges[n].second.pos
}

// delete all edges from provided array which point to both
// provided vertices
func deleteEdges(g *graph, vert1, vert2 int) {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i].first.pos == vert1 || g.edges[i].first.pos == vert2 {
			if g.edges[i].second.pos == vert1 || g.edges[i].second.pos == vert2 {
				tempSlice := make([]*edge, len(g.edges)-1)
				copy(tempSlice, g.edges[:i])
				if i < len(g.edges)-1 {
					copy(tempSlice[i:], g.edges[i+1:])
				}
				g.edges = tempSlice
				i--
			}
		}
	}
	deleteVertEdges(g, vert1, vert2)
	deleteVertEdges(g, vert2, vert1)
	return
}

func deleteVertEdges(g *graph, vert1, vert2 int) {
	// find graph element with pos == vert1
	k := 0
	for g.vertices[k].pos != vert1 {
		k++
	}
	// search its edges for ones pointing to vert1 or vert2
	for i := 0; i < len(g.vertices[k].edges); i++ {
		if g.vertices[k].edges[i].first.pos == vert1 || g.vertices[k].edges[i].first.pos == vert2 {
			if g.vertices[k].edges[i].second.pos == vert1 || g.vertices[k].edges[i].second.pos == vert2 {
				// delete them
				tempSlice := make([]*edge, len(g.vertices[k].edges)-1)
				copy(tempSlice, g.vertices[k].edges[:i])
				copy(tempSlice[i:], g.vertices[k].edges[i+1:])
				g.vertices[k].edges = tempSlice
				i--
			}
		}
	}
}

// opens file "filename" and loads its contents to an array of vertices
//    a "graph"
func loadgraph(fileToken string) [][]int {
	b, err := ioutil.ReadFile(fileToken)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	lines := strings.Split(string(b), "\n")

	arr := make([][]int, 0, len(lines))

	for _, v := range lines {
		textints := strings.Split(v, "\t")
		subarr := make([]int, 0, len(textints))
		for _, w := range textints {
			n, err := strconv.Atoi(w)
			if err != nil {
				continue
			}
			subarr = append(subarr, n)
		}
		arr = append(arr, subarr)
	}
	return arr
}

func atog(aa [][]int) *graph {
	var g *graph = new(graph)
	for _, v := range aa {
		g.vertices = append(g.vertices, &vertex{pos: v[0]})
	}
	for _, v := range aa {
		vp1 := vertexByPos(g, v[0])
		for i := 1; i < len(v); i++ {
			if v[i] > v[0] {
				vp2 := vertexByPos(g, v[i])
				ne := &edge{vp1, vp2}
				g.edges = append(g.edges, ne)
				vp1.edges = append(vp1.edges, ne)
				vp2.edges = append(vp2.edges, ne)
			}
		}
	}
	return g
}

func vertexByPos(g *graph, n int) *vertex {
	for _, v := range g.vertices {
		if v.pos == n {
			return v
		}
	}
	return nil
}

func printgraph(g *graph) {
	fmt.Println("Edges:")
	for _, v := range g.edges {
		fmt.Print(" ", v.first.pos, "-", v.second.pos, " |")
	}
	fmt.Print("\n")
	fmt.Println("Vertices:")
	for _, v := range g.vertices {
		fmt.Print(v.pos, ":")
		for _, w := range v.edges {
			fmt.Print(" ", w.first.pos, " ", w.second.pos, " |")
		}
		fmt.Print("\n")
	}
}
