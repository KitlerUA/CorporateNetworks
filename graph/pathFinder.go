package graph

//FindPath - restore path from slice of parents
func FindPath(startPoint, endPoint int, parent []int) []int {
	var result []int
	if startPoint == endPoint || endPoint == -1 {
		result = append(result, startPoint)
		return result
	}

	result = append(result, FindPath(startPoint, parent[endPoint], parent)...)
	result = append(result, endPoint)
	return result
}

//FindDistance - find distance with given path and Graph
func FindDistance(path []int, graph Graph) int {
	dist := 0
	for i := 0; i < len(path)-1; i++ {
		dist += graph.mapOfMap[path[i]][path[i+1]]
	}
	return dist
}

func isInPath(endPoint, pathPoint int, parents []int) bool {
	for endPoint != -1 {
		if parents[endPoint] == pathPoint {
			return true
		}
		endPoint = parents[endPoint]
	}
	return false
}

func isLeaf(point int, parents []int) bool {
	for i := range parents {
		if parents[i] == point {
			return false
		}
	}
	return true
}
