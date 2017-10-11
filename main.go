package main

import (
	"math/rand"
	"time"

	"github.com/KitlerUA/CorporateNetworks/graph"
)

func main() {
	g := graph.Graph{}
	g.Build(15)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 20; i++ {
		g.AddEdge(r.Intn(g.NumberOfVertexes()), r.Intn(g.NumberOfVertexes()), r.Intn(7)+1)
	}

	g.BFS(13)
}
