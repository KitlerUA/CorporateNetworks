package graph

import "fmt"

const inf = 66666

type Graph struct {
	v        int
	mapOfMap map[int]map[int]int
}

func (g *Graph) Build(v int) {
	g.v = v
	g.mapOfMap = make(map[int]map[int]int)
	for i:= 0; i<v;i++{
		g.mapOfMap[i] = make(map[int]int)
	}
}

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
