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
		minInd := IndexOfMin(q, dist)
		uI, ok := q.Get(minInd)
		u := uI.(int)
		if !ok{
			fmt.Println("Shit")
		}
		q.Remove(minInd)

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
func IndexOfMin(list singlylinkedlist.List, dist []int) int{
	it := list.Iterator()
	min := inf
	minIndex := -1
	for it.Next() {
		if dist[it.Value().(int)] < min {
			min = dist[it.Value().(int)]
			minIndex = it.Index()
		}
	}
	return minIndex
}

