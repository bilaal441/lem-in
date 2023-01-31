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
	Name, Room   string
	Adj          []*Vertex
	NumAnts      int
	edges        []*Edge
	RoomCapacity int
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
func checkPathOverlap(paths [][]*Vertex) bool {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			for k := len(paths[i][1 : len(paths[i])-1]); k > 0; k-- {
				for l := len(paths[j][1 : len(paths[j])-1]); l > 0; l-- {
					vertex1 := paths[i][k]
					vertex2 := paths[j][l]

					if vertex1 == vertex2 {
						fmt.Println(vertex1.Name, vertex2.Name)
						for _, curr := range paths[i][1 : len(paths[i])-1] {
							fmt.Println("f", curr.Name)
						}
						for _, curr := range paths[j][1 : len(paths[j])-1] {
							fmt.Println("l", curr.Name)
						}

						return true
					}
				}
			}
		}
	}
	return false
}

// func (g *Graph) InitializeResidualNetwork(links []string) *Graph {
// 	residualNetwork := &Graph{}
// 	// Add all the vertices to the residual network
// 	for _, vertex := range g.vertices {
// 		residualNetwork.AddVertex(vertex.Name, vertex.Room)
// 	}
// 	// Add all the edges to the residual network, with a residual capacity of 0

// 	residualNetwork.AddEdges(links, 1)

// 	return residualNetwork
// }

// func (g *Graph) FindAugmentingPath(source, sink string) ([]*Vertex, int) {

// 	queue := []*Vertex{g.vertices[g.GetIndexbyRoom(source)]}
// 	predecessors := make(map[string]*Vertex)

// 	for len(queue) > 0 {
// 		u := queue[0]
// 		queue = queue[1:]

// 		if u.Name == sink {
// 			fmt.Println(predecessors)

// 		}

// 		for _, v := range u.Adj {

// 			if predecessors[v.Name] == v && v.Flow < 1 && v.Room != "start" {

// 				queue = append(queue, v)
// 				predecessors[v.Name] = v

// 			}

// 		}

// 	}

// 	// If we couldn't find an augmenting path, return an empty path and 0 capacity
// 	return []*Vertex{}, 0
// }

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

			if edge.Flow-edge.capacity < 0 && !visited[v] && v.RoomCapacity < 1 {
				visited[v] = true
				parents[v] = current
				queue = append(queue, v)

			}
		}
	}

	return []*Vertex{}, false
}

func incrementFlow(path []*Vertex) {
	bottleneck := math.MaxInt32
	// d:= math.MinInt32
	// Finding bottleneck capacity
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
}

func lem_in(graph *Graph) [][]*Vertex {
	paths := [][]*Vertex{}
	for {
		path, found := graph.FindAugmentingPath()
		// fmt.Println("path")
		// for _, curr := range path {
		// 	fmt.Println(curr.Name)
		// }

		// fmt.Println(found)
		if !found {
			break
		}
		paths = append(paths, path)
		incrementFlow(path)
	}
	return paths
}

func (data *Graph) findWrightPaths() {
	d := lem_in(data)
	// b := lem_in(data)
	if checkPathOverlap(d) {

		fmt.Println(d)
	}

	// check paths overlap

}

func main() {
	fileName := os.Args[1]

	_, rooms, links := GetData(fileName)
	data := Graph{}
	data.AddVertexes(rooms, links)

	data.AddEdges(links, 1)

	data.findWrightPaths()
	// lemIn(&data)

	// fmt.Println(findNonOverlappingPaths(data.getPaths()))

}
