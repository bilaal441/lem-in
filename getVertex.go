package main

func (g *Graph) GetVertex(name string) *Vertex {
	for i, v := range g.vertices {
		if v.Name == name {
			return g.vertices[i]

		}

	}
	return nil
}
