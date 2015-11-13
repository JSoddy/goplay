package main 

//type edge is the edge of a graph
type edge struct {
	weight  	int
	vertexa 	int
	vertexb 	int
}

// edgeList is a slice of edge
// implements sort.Interface
type edgeList []edge

type vertex struct {
	parent 		int
	depth 		int
}

type graph struct {
	edges 		edgeList
	vertices 	[]vertex
	kvalue		int
}

func (g graph) countUnions() int {
	count := 0
	m := make([]bool, len(g.vertices))
	for i, _ := range g.vertices {
		m[g.FIND(i+1)] = true
	}
	for _, v := range m {
		if v {
			count++
		}
	}
	return count
}

//returns the minimum spacing between two components
func (g graph) spacing() int {
	minSpace := 0
	for i := 0; i < len(g.edges) && minSpace == 0; i++ {
		if g.FIND(g.edges[i].vertexa) != g.FIND(g.edges[i].vertexb) {
			minSpace = g.edges[i].weight
		}
	}
	return minSpace
}
//returns the leader of the vertex in g numbered vert
func (g *graph) FIND(vert int) int {
	if p := g.vertices[vert-1].parent; p == vert {
		return p
	} else {
		gp := g.FIND(p)
		g.vertices[vert-1].parent = gp
		return gp
	}
}

func (g *graph) MERGE(vert1 int, vert2 int) {
	parent1 := g.FIND(vert1)
	parent2 := g.FIND(vert2)
	if (g.FIND(vert1) == g.FIND(vert2)) {
		return
	} else {
		depth1 := g.vertices[vert1-1].depth
		depth2 := g.vertices[vert2-1].depth
		if depth1 > depth2 {
			g.vertices[parent2-1].parent = parent1
		} else if depth2 > depth1 {
			g.vertices[parent1-1].parent = parent2
		} else {
			g.vertices[parent1-1].depth += 1
			g.vertices[parent2-1].parent = parent1
		}
		g.kvalue -= 1
		return
	}
}

func (e edgeList) Len() int {
	return len(e)
}

func (e edgeList) Less(i, j int) bool {
	return e[i].weight < e[j].weight
}

func (e edgeList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}