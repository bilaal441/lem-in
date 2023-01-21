package main

import (
	"strings"
)

func (g *Graph) addEdge(fromv, tov string, capacity int) {
	from := g.GetVertex(fromv)
	to := g.GetVertex(tov)
	edge := &Edge{from: from, to: to, capacity: capacity, Flow: 0}
	reverseEdge := &Edge{from: to, to: from, capacity: capacity, Flow: 0}
	from.edges = append(from.edges, edge)
	to.edges = append(to.edges, reverseEdge)
	from.Adj = append(from.Adj, to)
	to.Adj = append(to.Adj, from)
}

func (g *Graph) AddEdges(links []string, capacity int) {

	for _, curr := range links {
		link := strings.Split(curr, "-")
		g.addEdge(string(link[0]), string(link[1]), capacity)
	}

}
