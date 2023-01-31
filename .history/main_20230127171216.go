package main

import (
	"fmt"
	"math"
	"os"
)

type Graph struct {
	vertices []*Vertex
}
type Edge struct {
	from, to       *Vertex
	capacity, Flow int
}

type Vertex struct {
	Name, Room string
	Adj        []*Vertex
	NumAnts    int
	edges      []*Edge
}

func Contains(s []*Vertex, name string) bool {
	for _, v := range s {
		if v.Name == name {
			return true

		}

	}
	return false
}

func (g *Graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex  %v :", v.Name)
		for _, v := range v.Adj {
			fmt.Printf("%v ", v.Name)
		}
	}
}

func (graph *Graph) getEdge(current, v *Vertex) *Edge {
	for _, edge := range current.edges {
		if edge.to == v {
			return edge
		}
	}
	return nil
}
func (graph *Graph) FindAugmentingPath() ([]*Vertex, bool) {
	start := graph.vertices[graph.GetIndexbyRoom("start")]
	end := graph.vertices[graph.GetIndexbyRoom("end")]
	queue := []*Vertex{start}
	parents := make(map[*Vertex]*Vertex)
	parents[start] = nil
	visited := make(map[*Vertex]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			path := []*Vertex{}

			for current != nil {
				// fmt.Println(current.Name)
				path = append(path, current)
				current = parents[current]
			}

			for i := 0; i < len(path)/2; i++ {
				path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
			}

			return path, true
		}

		for _, v := range current.Adj {
			edge := graph.getEdge(current, v)

			if edge.Flow-edge.capacity < 0 && !visited[v] {
				visited[v] = true
				parents[v] = current
				queue = append(queue, v)

			}
		}
	}

	return []*Vertex{}, false
}

func incrementFlow(path []*Vertex) int {
	bottleneck := math.MaxInt32
	for i := 0; i < len(path)-1; i++ {
		for _, edge := range path[i].edges {
			if edge.to == path[i+1] || edge.from == path[i+1] {
				bottleneck = int(math.Min(float64(bottleneck), float64(edge.capacity-edge.Flow)))
				break
			}
		}
	}

	for i := 0; i < len(path)-1; i++ {
		for _, edge := range path[i].edges {
			if edge.to == path[i+1] || edge.from == path[i+1] {
				edge.Flow += bottleneck

				// Add flow to reverse edge
				for _, revEdge := range edge.to.edges {
					if revEdge.to == path[i] || revEdge.from == path[i] {

						revEdge.Flow -= bottleneck
						break
					}
				}
				break
			}
		}
	}

	return bottleneck
}

func lem_in(graph *Graph) int {
	maxflow := 0

	for {
		path, found := graph.FindAugmentingPath()
		// for _, curr := range path {
		// 	fmt.Println(curr.Name)
		// }
		if !found {
			break
		}

		maxflow += incrementFlow(path)
	}
	return maxflow
}
func assignPaths(numAnts int, paths [][]string) map[int]int {
	pathWithAnts := make(map[int]int)
	currentPath := 0
	for i := 1; i <= numAnts; i++ {

		pathWithAnts[currentPath]++
		if currentPath+1 < len(paths) &&
			pathWithAnts[currentPath]+len(paths[currentPath]) >
				pathWithAnts[currentPath+1]+len(paths[currentPath+1]) {

			currentPath++

		} else {

			currentPath = 0
		}
	}

	// fmt.Println(pathWithAnts)
	return pathWithAnts
}

func sendAnts(paths [][]string, ants int, pathAnt map[int]int) {
	// create a map to store the current position of each ant
	antPos := make(map[int]int)
	// set the initial position of each ant to -1 (not yet started)
	for i := 1; i <= ants; i++ {
		antPos[i] = -1
	}

	// keep track of the current time
	time := 0

	// keep track of the number of ants that have reached the end
	finishedAnts := 0

	for finishedAnts < ants {
		// increment the time
		time++

		// check if there are any ants that can move to the next room
		for i := 1; i <= ants; i++ {
			// if the ant is not yet started or is not at the end of the path
			if antPos[i] != -1 && antPos[i] != len(paths[pathAnt[i]])-1 {
				// move the ant to the next room
				antPos[i]++

				// if the ant has reached the end, increment the finishedAnts count
				if antPos[i] == len(paths[pathAnt[i]])-1 {
					finishedAnts++
				}
			}
		}

		// print the current status of each ant
		for i := 1; i <= ants; i++ {
			if antPos[i] == -1 {
				fmt.Printf("L%d- ", i)
			} else {
				if antPos[i] == len(paths[pathAnt[i]])-1 {
					fmt.Printf("L%d-end ", i)
				} else {
					fmt.Printf("L%d-%s ", i, paths[pathAnt[i]][antPos[i]])
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	fileName := os.Args[1]
	numAnts, rooms, links := GetData(fileName)
	data := Graph{}
	data.AddVertexes(rooms, links)
	data.AddEdges(links, 1)
	fmt.Println(lem_in(&data))
	paths := data.getPaths()
	fmt.Println(numAnts, paths)

	sendAnts(paths, numAnts, assignPaths(numAnts, paths))

}
