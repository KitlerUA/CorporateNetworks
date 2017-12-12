package graph

//DynAdd - find paths after adding vertex or or edges (can add only one vertex at once)
func (g *Graph) DynAdd(startPoint int, edges [][]int) ([]int, []int) {

	optDist, optParent := g.Dijkstra(startPoint)
	if len(edges) == 0 {
		return optDist, optParent
	}
	for _, e := range edges {
		if _, ok := g.mapOfMap[e[0]]; !ok {
			optParent = append(optParent, e[1])
			g.AddEdge(e[0], e[1], e[2])
			optDist = append(optDist, FindDistance(FindPath(startPoint, e[0], optParent), *g))
			break
		} else if _, ok := g.mapOfMap[e[1]]; !ok {
			optParent = append(optParent, e[0])
			g.AddEdge(e[0], e[1], e[2])
			optDist = append(optDist, FindDistance(FindPath(startPoint, e[1], optParent), *g))
			break
		}
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
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
