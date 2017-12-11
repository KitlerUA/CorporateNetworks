package graph

//PairMove - find new paths after changing edge (e1,e2) weight to 'val'
func (g *Graph) PairMove(startPoint, e1, e2, val int) ([]int, []int) {
	optDist, optParent := g.Dijkstra(startPoint)
	g.AddEdge(e1, e2, val)
	optDist[e1] = FindDistance(FindPath(startPoint, e1, optParent), *g)
	optDist[e2] = FindDistance(FindPath(startPoint, e2, optParent), *g)
	queue := make([]int, 0)
	queue = append(queue, e1)
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
	queue = append(queue, e2)
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
	return optDist, optParent
}
