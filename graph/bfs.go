package graph

import (
	"fmt"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
)

//BFS - find paths with BFS algorithm
func (g *Graph) BFS(startPoint int) ([]int, []int) {
	dist := make([]int, len(g.mapOfMap))
	parent := make([]int, len(g.mapOfMap))

	for i := range dist {
		dist[i] = inf
	}
	for i := range parent {
		parent[i] = -1
	}

	dist[startPoint] = 0

	queue := singlylinkedlist.List{}
	queue.Add(startPoint)

	for !queue.Empty() {
		uI, ok := queue.Get(queue.Size() - 1)
		if !ok {
			fmt.Print("Cannot convert interface to int")
			return []int{}, []int{}
		}
		u := uI.(int)
		queue.Remove(queue.Size() - 1)

		for k, v := range g.mapOfMap[u] {
			if dist[k] == inf {
				queue.Insert(0,k)
				dist[k] = dist[u] + v
				parent[k] = u
			}
		}
	}

	/*for k := range g.mapOfMap {
		fmt.Printf("Distance from Vertex %d to %d = %d \n", startPoint, k, dist[k])
		if dist[k] == inf {
			fmt.Print("\tPass not found")
			fmt.Println()
			continue
		}
		fmt.Print("\tWith path ")
		findPath(startPoint, k, parent)
		fmt.Println()
	}*/
	return dist, parent
}


