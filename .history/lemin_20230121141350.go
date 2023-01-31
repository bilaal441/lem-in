package main

import (
	"fmt"
	"sort"
)

// func ReduceInt(a []int, f func(int, int) int) int {
// 	occ := a[0]
// 	for _, curr := range a[1:] {
// 		occ = f(occ, curr)

// 	}
// 	return occ
// }

// func Chunk(slice []int, size int) [][]int {
// 	if size == 0 {
// 		return [][]int{}

// 	}
// 	val := len(slice) / size
// // 	s := [][]int{}

// 	for i := 0; i < val; i++ {
// 		s = append(s, slice[i*size:size*(i+1)])
// 	}

// 	if len(slice)%size != 0 {
// 		s = append(s, slice[size*val:])
// 	}

// 	return s

// }

// func mostPoints(questions [][]int) int64 {

// 	if len(questions) == 1 {
// 		fmt.Println(questions)
// 		return int64(questions[0][0])
// 	}

// 	val := int64(0)

// 	for i, curr := range questions {
// 		point := curr[0]

// 		brainPower := curr[1] + i + 1
// 		if len(questions) > brainPower {
// 			fmt.Println("point", curr, "next", questions[brainPower])
// 			val = int64(math.Max(float64(val), float64(point)+float64(questions[brainPower][0])))
// 		}

// 	}

// 	return val

// }

func findNonOverlappingPaths(graph *Graph, start, end *Vertex, visited map[*Vertex]bool, path []*Vertex) [][]*Vertex {
	visited[start] = true
	path = append(path, start)
	if start == end {
		return [][]*Vertex{path}
	}
	var allPaths [][]*Vertex
	for _, v := range start.Adj {
		if !visited[v] {
			edge := graph.getEdge(start, v)
			if edge.Flow-edge.capacity > 0 {

				newVisited := make(map[*Vertex]bool)
				for key, value := range visited {
					newVisited[key] = value
				}

				allPaths = append(allPaths, findNonOverlappingPaths(graph, v, end, newVisited, path)...)
			}
		}
	}
	return allPaths
}

func lemIn(graph *Graph) {
	start := graph.vertices[graph.GetIndexbyRoom("start")]
	end := graph.vertices[graph.GetIndexbyRoom("end")]
	visited := make(map[*Vertex]bool)
	allPaths := findNonOverlappingPaths(graph, start, end, visited, []*Vertex{})
	
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	for _, curr := range allPaths[1] {
		fmt.Println(curr.Name, len(allPaths))
	}

}
