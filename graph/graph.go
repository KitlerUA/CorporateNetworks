package graph

import "fmt"

const inf = 66666

//Graph - represent graph
type Graph struct {
	v        int
	mapOfMap map[int]map[int]int
}

//Build - set number of vertex and create empty graph
func (g *Graph) Build(v int) {
	g.v = v
	g.mapOfMap = make(map[int]map[int]int)
	for i := 0; i < v; i++ {
		g.mapOfMap[i] = make(map[int]int)
	}
}

//AddEdge - adds edge (u,v) and (v,u) with weight w
func (g *Graph) AddEdge(u, v, w int) {
	if _, ok := g.mapOfMap[u]; !ok {
		g.mapOfMap[u] = make(map[int]int)
	}
	g.mapOfMap[u][v] = w
	if _, ok := g.mapOfMap[v]; !ok {
		g.mapOfMap[v] = make(map[int]int)
	}
	g.mapOfMap[v][u] = w
	fmt.Println("Edge ", u, "-", v, " with weight ", w, " added")
}

//RemoveEdge - remove edge (u,v) and (v,u)
func (g *Graph) RemoveEdge(u, v int) {
	if _, ok := g.mapOfMap[u]; ok {
		if _, ok := g.mapOfMap[u][v]; ok {
			delete(g.mapOfMap[u], v)
		}
	}
	if _, ok := g.mapOfMap[v]; ok {
		if _, ok := g.mapOfMap[v][u]; ok {
			delete(g.mapOfMap[v], u)
		}
	}
	fmt.Println("Edge ", u, "-", v, " removed")
}

//NumberOfVertexes - return number of vertexes
func (g *Graph) NumberOfVertexes() int {
	return g.v
}
