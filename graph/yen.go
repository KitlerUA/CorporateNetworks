package graph

import (
	"math"
	"reflect"
)

//Yen - find first 'number' paths with Yen algorithm
func (g *Graph) Yen(startPoint, endPoint, number int) [][]int {
	result := make([][]int, 0, number)
	_, parents := g.Dijkstra(startPoint)
	path := FindPath(startPoint, endPoint, parents)
	result = append(result, path)
	potential := make([][]int, 0)
	for k := 1; k < number; k++ {
		for i := 0; i < len(result[k-1])-1; i++ {
			edgeToRestore := make([][]int, 0)
			spurNode := result[k-1][i]
			rootPath := result[k-1][:i]
			for _, p := range result {
				if reflect.DeepEqual(rootPath, p[:i]) {

					if _, ok := g.mapOfMap[result[k-1][i]][result[k-1][i+1]]; ok {
						edgeToRestore = append(edgeToRestore, []int{result[k-1][i], result[k-1][i+1], g.mapOfMap[result[k-1][i]][result[k-1][i+1]]})
						g.RemoveEdge(result[k-1][i], result[k-1][i+1])
					}
				}
			}
			_, parents := g.Dijkstra(spurNode)
			spurPath := FindPath(spurNode, endPoint, parents)
			var pot []int
			pot = append(pot, rootPath...)
			pot = append(pot, spurPath...)
			potential = append(potential, pot)
			for _, r := range edgeToRestore {
				g.AddEdge(r[0], r[1], r[2])
			}

		}

		if len(potential) == 0 {
			break
		}
		index := 0
		for {
			index = findIndexOfMinPath(potential, *g)
			if sliceContainsSlice(result, potential[index]) {
				potential = append(potential[:index], potential[index+1:]...)
				continue
			}
			break
		}
		result = append(result, potential[index])
		potential = append(potential[:index], potential[index+1:]...)
	}
	return result
}

func findIndexOfMinPath(paths [][]int, graph Graph) int {
	if len(paths) == 0 {
		return -1
	}
	min := math.MaxInt32
	index := 0
	for i := range paths {
		temp := FindDistance(paths[i], graph)
		if temp < min {
			min = temp
			index = i
		}
	}
	return index
}

func sliceContainsSlice(matrix [][]int, vector []int) bool {
	for i := range matrix {
		if reflect.DeepEqual(matrix[i], vector) {
			return true
		}
	}
	return false
}
