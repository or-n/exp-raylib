package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Point struct {
	X, Y float64
}

type Edge struct {
	U, V int
}

func orientation(a, b, c Point) int {
	val := (b.Y-a.Y)*(c.X-b.X) - (b.X-a.X)*(c.Y-b.Y)
	if val > 0 {
		return 1
	}
	if val < 0 {
		return -1
	}
	return 0
}

func onSegment(a, b, c Point) bool {
	return math.Min(a.X, b.X) <= c.X && c.X <= math.Max(a.X, b.X) &&
		math.Min(a.Y, b.Y) <= c.Y && c.Y <= math.Max(a.Y, b.Y)
}

const (
	radius = 0.1
	size   = 1000
	r      = radius * 72 / size
)

func distPointSegment(p, a, b Point) float64 {
	px := b.X - a.X
	py := b.Y - a.Y
	norm := px*px + py*py
	u := ((p.X-a.X)*px + (p.Y-a.Y)*py) / norm
	if u < 0 {
		u = 0
	} else if u > 1 {
		u = 1
	}
	dx := a.X + u*px - p.X
	dy := a.Y + u*py - p.Y
	return math.Hypot(dx, dy)
}

func segmentsIntersect(p1, q1, p2, q2 Point) bool {
	if p1 == p2 || p1 == q2 || q1 == p2 || q1 == q2 {
		return false
	}
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)
	if o1 != o2 && o3 != o4 {
		return true
	}
	if o1 == 0 && onSegment(p1, q1, p2) {
		return true
	}
	if o2 == 0 && onSegment(p1, q1, q2) {
		return true
	}
	if o3 == 0 && onSegment(p2, q2, p1) {
		return true
	}
	if o4 == 0 && onSegment(p2, q2, q1) {
		return true
	}
	return false
}

type Graph struct {
	points []Point
	edges  []Edge
}

func (g Graph) graphviz() {
	fmt.Println("graph G {")
	fmt.Printf(`  graph [layout=neato, size="1,1!", dpi=%d];`, size)
	fmt.Println("")
	fmt.Println(`  invis1 [pos="0,0!", shape=point, width=0, height=0];`)
	fmt.Printf(`  invis2 [pos="%d,%d!", shape=point, width=0, height=0];`, size, size)
	fmt.Println("")
	fmt.Printf(`  node  [shape=circle, width=%f, height=%f, fixedsize=true, label=""];`, 2*radius, 2*radius)
	fmt.Println("")
	for i, p := range g.points {
		fmt.Printf("  %d [pos=\"%f,%f!\"];\n", i, p.X*size, p.Y*size)
	}
	for _, e := range g.edges {
		fmt.Printf("  %d -- %d;\n", e.U, e.V)
	}
	fmt.Println("}")
}

func new(n int, rng *rand.Rand) Graph {
	points := make([]Point, 0, n)
	for len(points) < n {
		p := Point{rng.Float64(), rng.Float64()}
		ok := true
		for _, q := range points {
			dx := p.X - q.X
			dy := p.Y - q.Y
			if math.Hypot(dx, dy) < 4*r {
				ok = false
				break
			}
		}
		if ok {
			points = append(points, p)
		}
	}
	edges := []Edge{}
	type candidate struct {
		u, v   int
		length float64
	}
	var candidates []candidate
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y
			candidates = append(candidates, candidate{i, j, math.Hypot(dx, dy)})
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].length < candidates[j].length
	})
	for _, c := range candidates {
		ok := true
		for _, e := range edges {
			if segmentsIntersect(points[c.u], points[c.v], points[e.U], points[e.V]) {
				ok = false
				break
			}
		}
		for i, p := range points {
			if i != c.u && i != c.v {
				d := distPointSegment(p, points[c.u], points[c.v])
				if d < 3*r {
					ok = false
					break
				}
			}
		}
		if ok {
			edges = append(edges, Edge{c.u, c.v})
		}
	}
	return Graph{points, edges}
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := new(128, rng)
	g.graphviz()
}
