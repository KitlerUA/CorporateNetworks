package graph

import "math"

//DynRemove - find paths after removing edge or vertex (mean remove all edges contains this vertex)
func (g *Graph) DynRemove(startPoint int, edges [][]int) ([]int, []int) {
	optDist, optParent := g.Dijkstra(startPoint)
	if len(edges) == 0 {
		return optDist, optParent
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1], math.MaxInt32)
	}
	for i := range optDist {
		optDist[i] = FindDistance(FindPath(startPoint, i, optParent), *g)
	}
	queue := make([]int, 0)
	for _, e := range edges {
		queue = append(queue, e[0])
		for len(queue) > 0 {
			for j := range g.mapOfMap[queue[0]] {
				if j != optParent[queue[0]] && !isInPath(j, queue[0], optParent) {
					d := optDist[j] + g.mapOfMap[j][queue[0]]
					if d < optDist[queue[0]] {
						optDist[queue[0]] = d
						optParent[queue[0]] = j
					}
				} else if isInPath(j, queue[0], optParent) {
					queue = append(queue, j)
				}
			}
			queue = queue[1:]
		}
		queue = append(queue, e[1])
		for len(queue) > 0 {
			for j := range g.mapOfMap[queue[0]] {
				if j != optParent[queue[0]] && !isInPath(j, queue[0], optParent) {
					d := optDist[j] + g.mapOfMap[j][queue[0]]
					if d < optDist[queue[0]] {
						optDist[queue[0]] = d
						optParent[queue[0]] = j
					}
				} else if isInPath(j, queue[0], optParent) {
					queue = append(queue, j)
				}
			}
			queue = queue[1:]
		}
	}
	return optDist, optParent
}
