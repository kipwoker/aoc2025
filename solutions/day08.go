package solutions

import (
	"sort"
	"strconv"
	"strings"
)

type Day08 struct{}

func (d Day08) Day() string {
	return "08"
}

type Point3 struct {
	X, Y, Z int
}

type Edge struct {
	I, J  int
	Dist2 int64
}

func parse08(input string) []Point3 {
	lines := strings.Split(input, "\n")

	result := []Point3{}
	for _, l := range lines {
		coords := strings.Split(l, ",")
		result = append(result, Point3{
			X: parseInt(coords[0]),
			Y: parseInt(coords[1]),
			Z: parseInt(coords[2]),
		})
	}

	return result
}

type Output8 struct {
	Ds     []int
	LastX1 int
	LastX2 int
}

func groupSizes(points []Point3, limit int) Output8 {
	n := len(points)

	edges := []Edge{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := int64(points[i].X - points[j].X)
			dy := int64(points[i].Y - points[j].Y)
			dz := int64(points[i].Z - points[j].Z)
			d2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{I: i, J: j, Dist2: d2})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].Dist2 < edges[b].Dist2
	})

	uf := NewUnionFind(n)

	lastX1 := 0
	lastX2 := 0
	for i, e := range edges {
		if limit != -1 && i == limit {
			break
		}

		ri := uf.Find(e.I)
		rj := uf.Find(e.J)

		if ri == rj {
			continue
		}

		uf.Union(e.I, e.J)

		lastX1 = points[e.I].X
		lastX2 = points[e.J].X
	}

	compMap := make(map[int]int)
	for i := range n {
		r := uf.Find(i)
		compMap[r]++
	}

	result := []int{}
	for _, sz := range compMap {
		result = append(result, sz)
	}
	return Output8{
		Ds:     result,
		LastX1: lastX1,
		LastX2: lastX2,
	}
}

func (d Day08) Execute1(input string) string {
	points := parse08(input)
	limit := 0
	if len(points) == 20 {
		limit = 10
	} else {
		limit = 1000
	}

	output := groupSizes(points, limit)
	sizes := output.Ds

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	n := len(sizes)

	mul := sizes[n-1] * sizes[n-2] * sizes[n-3]

	return strconv.Itoa(mul)
}

func (d Day08) Execute2(input string) string {
	points := parse08(input)
	output := groupSizes(points, -1)

	result := output.LastX1 * output.LastX2

	return strconv.Itoa(result)
}
