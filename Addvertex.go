package main

import (
	"fmt"
	"strings"
)

func (g *Graph) AddVertex(name, room string) {
	if Contains(g.vertices, name) {
		fmt.Printf("vertex %v because it is already exists", name)
	} else {
		g.vertices = append(g.vertices, &Vertex{Name: name, Room: room})

	}
}

func (g *Graph) AddVertexes(r, l []string) {
	nmofn := r[1:]
	for i, curr := range nmofn {

		switch {
		case i-1 >= 0 && strings.ReplaceAll(nmofn[i-1], "#", "") == "start":

			g.AddVertex(strings.Split(curr, " ")[0], "start")
		case i-1 >= 0 && strings.ReplaceAll(nmofn[i-1], "#", "") == "end":
			g.AddVertex(strings.Split(curr, " ")[0], "end")
		case curr == "##start" || curr == "##end":
			continue
		default:
			g.AddVertex(strings.Split(curr, " ")[0], "normal")
		}
	}
}
