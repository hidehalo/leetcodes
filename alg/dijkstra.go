package main

import (
	"fmt"
)

type vertexID uint

type weight uint

type vertex struct {
	ID vertexID
}

func (v *vertex) String() string {
	return fmt.Sprintf("vertex %d\n", v.ID)
}

type edge struct {
	w   weight
	src *vertex
	tgt *vertex
}

func (e *edge) String() string {
	return fmt.Sprintf("%d->%d|weight:%d\n", e.src.ID, e.tgt.ID, e.w)
}

type Graph struct {
	e map[vertexID]map[vertexID]*edge
	v map[vertexID]*vertex
}

func (g *Graph) String() string {
	contents := ""
	for _, v := range g.v {
		contents += v.String()
		for _, e := range g.e[v.ID] {
			contents += e.String()
		}
	}

	return contents
}

func (g *Graph) AppendVertex(v vertex) {
	g.v[v.ID] = &v
}

func (g *Graph) AppendEdge(src *vertex, tgt *vertex, w weight) {
	if _, ok := g.e[src.ID]; !ok {
		g.e[src.ID] = make(map[vertexID]*edge)
	}
	g.e[src.ID][tgt.ID] = &edge{
		w:   w,
		src: src,
		tgt: tgt,
	}
}

func (g *Graph) GetWeight(e *edge) weight {
	if e, ok := g.e[e.src.ID][e.tgt.ID]; ok {
		return e.w
	}

	return 0
}

func (g *Graph) GetVertex(ID vertexID) *vertex {
	return g.v[ID]
}

func Constructor(v []vertexID, w map[vertexID]map[vertexID]weight) *Graph {
	g := &Graph{
		e: make(map[vertexID]map[vertexID]*edge),
		v: make(map[vertexID]*vertex),
	}
	for _, ID := range v {
		g.AppendVertex(vertex{ID})
	}
	for srcID, edges := range w {
		for tgtID, weight := range edges {
			g.AppendEdge(g.GetVertex(srcID), g.GetVertex(tgtID), weight)
		}
	}

	return g
}

// func relax(u, v, w) {
// 	if d[v] > d[u]+w[u][v] {
// 		d[v] = d[u] + w[u][v]
// 		s[v] = u
// 	}
// }

func Dijkstra(g *Graph, w [][]weight, s *vertex) uint {
	// init(g, s)
	// S := make([]weight, 0, len(g.v))
	// Q := g.v
	// for len(Q) > 0 {
	// u := extractMin(Q)
	// S = append(S, u)
	// for _, v := range g.e[u] {
	// relax(u, v, w)
	// }
	// }

	return 0
}

func main() {
	v := []vertexID{1, 2, 3, 4, 5, 6, 8, 12, 15}
	w := map[vertexID]map[vertexID]weight{
		1: {
			2: 5,
			3: 2,
		},
		2: {
			4: 1,
			6: 2,
		},
		3: {
			8:  4,
			15: 5,
		},
		8: {
			12: 3,
			15: 3,
		},
	}
	g := Constructor(v, w)
	fmt.Println(g)
}
