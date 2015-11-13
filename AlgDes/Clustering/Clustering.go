package main 

import (
	"fmt"
	"sort"
	"UJS/file"
	"strconv"
	"strings"
)

const fileToken = "D:\\James\\Temp\\clustering1.txt"

func main() {
	// import graph
	g := loadGraph()
	// sort by lowest distance
	sort.Sort(g.edges)
	// clustering
	g = cluster(g)
	// find min spacing
	spacing := g.spacing()
	// output
	fmt.Println(spacing)
}

func cluster(g graph) graph {
	for i := 0; i < len(g.edges) && g.kvalue > 4; i++ {
		//fmt.Println(i, g.kvalue)
		//fmt.Println(g.edges[i].vertexa)
		//fmt.Println(g.edges[i].vertexb)
		g.MERGE(g.edges[i].vertexa, g.edges[i].vertexb)
	}
	return g
}

func loadGraph() graph {
	lines := file.FileLines(fileToken)
	graphSize, _ := strconv.Atoi((strings.Split(lines[0], "\n")[0]))
	lines = lines[1:]
	vertices := make([]vertex, graphSize)
	for i, _ := range vertices {
		vertices[i] = vertex{i+1, 0}
	}
	edges := make(edgeList, len(lines))
	for i, v := range lines {
		edges[i] = newEdge(v)
	}
	return graph{edges:edges, vertices:vertices, kvalue:len(vertices)}
}

func newEdge(line string) edge {
	lessNewLine := strings.Split(line, "\n")
	lineSlice := strings.Split(lessNewLine[0], " ")
	edge1, _ := strconv.Atoi(lineSlice[0])
	edge2, _ := strconv.Atoi(lineSlice[1])
	weight, _ := strconv.Atoi(lineSlice[2])
	return edge{weight, edge1, edge2}
}