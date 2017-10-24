package graph

import (
	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"fmt"
)

func (g *Graph) Dijkstra(startPoint int){

	parent := make([]int, len(g.mapOfMap))
	dist := make([]int, len(g.mapOfMap))

	q := singlylinkedlist.List{}

	for i := range dist{
		parent[i] = -1
		dist[i] = inf
		q.Add(i)
	}

	dist[startPoint] = 0

	for !q.Empty(){
		u := minInSlice(dist, q)
		fmt.Println("Index of min ",u,dist[u], q.Size())
		q.Remove(u)

		for k, v := range g.mapOfMap[u]{
			if q.Contains(k){
				alt := dist[u] + v
				if alt < dist[k] {
					dist[k] = alt
					parent[k] = u
				}
			}
		}

	}

	for k := range g.mapOfMap {
		fmt.Printf("Distance from Vertex %d to %d = %d \n", startPoint, k, dist[k])
		if dist[k] == inf {
			fmt.Print("\tPass not found")
			fmt.Println()
			continue
		}
		fmt.Print("\tWith path ")
		findPath(startPoint, k, parent)
		fmt.Println()
	}
}

//IndexOfMin - return index of minimal element or -1
func IndexOfMin(list singlylinkedlist.List) int{
	it := list.Iterator()
	min := inf
	minIndex := -1
	for it.Next() {
		if it.Value().(int) < min{
			minIndex = it.Index()
			min = it.Value().(int)
		}
	}
	return minIndex
}

func minInSlice(list []int, q singlylinkedlist.List) int{
	m :=inf
	ind := -1
	for i, v := range list{
		if q.Contains(i) &&  v < m {
			m = v
			ind = i
		}
	}
	return ind
}