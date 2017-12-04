package graph

import (
	"math"
	"fmt"
)

//Bellman - find paths with Bellman-Ford algorithm
func (g *Graph) Bellman(startPoint int) {
	result := make([]int, len(g.mapOfMap))
	for i := range result {
		result[i] = math.MaxInt32
	}
	result[startPoint] = 0
	parents := make([]int, len(g.mapOfMap))
	for i := range parents {
		parents[i] = -1
	}
	for range g.mapOfMap {
		for i:= range g.mapOfMap {
			for j := range g.mapOfMap[i] {
				if result[i] < math.MaxInt32 {
					if result[j] > result[i]+g.mapOfMap[i][j] {
						result[j] = result[i] + g.mapOfMap[i][j]
						parents[j] = i
					}
				}
			}
		}
	}
	fmt.Println(result)

	for k := range g.mapOfMap {
		fmt.Printf("Distance from Vertex %d to %d = %d \n", startPoint, k, result[k])
		if result[k] == inf {
			fmt.Print("\tPass not found")
			fmt.Println()
			continue
		}
		fmt.Print("\tWith path ")
		findPath(startPoint, k, parents)
		fmt.Println()
	}
}