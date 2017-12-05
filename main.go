package main

import (
	"fmt"

	"github.com/KitlerUA/CorporateNetworks/graph"
)

func main() {
	g := graph.Graph{}
	g.Build(6)

	g.AddEdge(0, 1, 7)
	g.AddEdge(0, 2, 9)
	g.AddEdge(0, 5, 14)
	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 5, 2)
	g.AddEdge(2, 3, 11)
	g.AddEdge(5, 4, 9)
	g.AddEdge(3, 4, 6)
	fmt.Println("------BFS------")
	dist, parent := g.BFS(0)
	fmt.Println("Path from 0 to 3 ", graph.FindPath(0, 3, parent), " distance = ", dist[3])
	fmt.Println("----Dijkstra----")
	dist, parent = g.Dijkstra(0)
	fmt.Println("Path from 0 to 3 ", graph.FindPath(0, 3, parent), " distance = ", dist[3])
	fmt.Println("--Bellmab-Ford--")
	dist, parent = g.Bellman(0)
	fmt.Println("Path from 0 to 3 ", graph.FindPath(0, 3, parent), " distance = ", dist[3])
	fmt.Println("------Yen------")
	paths := g.Yen(0, 3, 4)
	fmt.Println("4 shortest pathes from 0 to 3 ", paths)
}
