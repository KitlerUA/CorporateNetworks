package graph

func (g *Graph) DynAdd(startPoint int, edges [][]int) {
	optDist, optParent := g.Dijkstra(startPoint)
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[3])
	}

}
