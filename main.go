package main

import (

	"github.com/KitlerUA/CorporateNetworks/graph"
	"fmt"
)

func main() {
	g := graph.Graph{}
	g.Build(6)
	/*s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 20; i++ {
		g.AddEdge(r.Intn(g.NumberOfVertexes()), r.Intn(g.NumberOfVertexes()), r.Intn(7)+1)
	}
	*/
	g.AddEdge(0,1,7)
	g.AddEdge(0,2,9)
	g.AddEdge(0,5,14)
	g.AddEdge(1,2,10)
	g.AddEdge(1,3,15)
	g.AddEdge(2,5,2)
	g.AddEdge(2,3,11)
	g.AddEdge(5,4,9)
	g.AddEdge(3,4,6)
	g.BFS(0)
	fmt.Println("----------------------")
	g.Dijkstra(0)
}
