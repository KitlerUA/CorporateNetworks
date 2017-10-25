# CorporateNetworks

Simple package to build network and find pathes between vertexes

### Network

Implemented like graph (which implemented like list of edges)

### How to use

Simply create Graph with number of vertex. After add edges to Graph. Use BFS, Dijkstra methods to find pathes. Be happy!

### Example

```go
package main

import (
	"github.com/KitlerUA/CorporateNetworks/graph"
	"fmt"
)

func main() {
	g := graph.Graph{}
	g.Build(6)
	//from Wikipedia
	g.AddEdge(0,1,7)
	g.AddEdge(0,2,9)
	g.AddEdge(0,5,14)
	g.AddEdge(1,2,10)
	g.AddEdge(1,3,15)
	g.AddEdge(2,5,2)
	g.AddEdge(2,3,11)
	g.AddEdge(5,4,9)
	g.AddEdge(3,4,6)
  
	fmt.Println("Using BFS algorithm")
	g.BFS(0)
	fmt.Println("Using Dijkstra algorithm")
	g.Dijkstra(0)
}
````
