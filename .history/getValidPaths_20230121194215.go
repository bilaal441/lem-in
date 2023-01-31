package main

import (
	"fmt"
)

func pop(s *[][]string) []string {
	toBeRemove := (*s)[0]
	*s = (*s)[1:]
	// fmt.Print(s)
	return toBeRemove
}

func (g *Graph) GetIndexbyRoom(t string) int {
	for i, v := range g.vertices {
		if v.Room == t {
			return i

		}

	}
	return -1
}

func Clone(path []string) (res []string) {
	res = append(res, path...)

	return
}

func isNotVisited(path []string, node string) bool {
	for _, v := range path {

		if v == node {
			// fmt.Print(nodeAddj, node)
			return false
		}
	}
	return true

}

func ContainsPath(visited, path []string) bool {
	// fmt.Print(path1)

	for _, currv := range visited {
		for _, p := range path[1 : len(path)-1] {

			if currv == p {
				fmt.Println(currv, p)
				return false
			}
		}

	}
	return true
}

func (g *Graph) getPaths() [][]string {
	result := [][]string{}
	queue := [][]string{}
	// add start node name into the queue
	queue = append(queue, []string{g.vertices[g.GetIndexbyRoom("start")].Name})

	// declear goal node  set end  in the graph
	goalNode := g.vertices[g.GetIndexbyRoom("end")].Name
	// fmt.Println(g.vertices[g.GetIndexbyRoom("end")].Room)

	// run loop while queue is not empty
	for len(queue) > 0 {

		// fmt.Print(queue)
		// pop first element from the queue and remove it from the queue
		path := pop(&queue)
		// fmt.Println(queue)
		// fmt.Println(queue)

		// get last index from paths
		lastNode := len(path) - 1
		// check if last node of the path is the goal node
		if Clone(path)[lastNode] == goalNode {
			// append that path to

			result = append(result, Clone(path))

		} else {
			// range over adjacent nodes of the last node in the path

			for _, nbr := range g.GetVertex(Clone(path)[lastNode]).Adj {
				// check if the adjacent node has not been visited and if the last node in the path has space for more ants

				if isNotVisited(Clone(path), nbr.Name) {
					// add the adjacent node to the path
					newPath := append(Clone(path), nbr.Name)
					// add the path to the queue
					queue = append(queue, newPath)

					// remove the last element from the path to backtrack
					path = Clone(newPath[:len(newPath)-1])
					// g.GetVertex(Clone(path)[lastNode]).NumAnts--

				}

			}
		}
	}

	return result
}
