package solutions

import (
	"container/list"
	"sort"
	"strconv"
)

type Solution interface {
	Day() string
	Execute1(input string) string
	Execute2(input string) string
}

func parseInt(x string) int {
	val, _ := strconv.Atoi(x)
	return val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

type Point struct {
	X, Y int
}

func (p Point) Hash() uint64 {
	return uint64(p.X)<<32 | uint64(p.Y)
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) GetNeighbors8() []Point {
	return []Point{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
	}
}

type Interval struct {
	Start, End int
}

func (iv Interval) Inside(value int) bool {
	return iv.Start <= value && value <= iv.End
}
func (iv Interval) Overlaps(other Interval) bool {
	return iv.Start <= other.End && other.Start <= iv.End
}
func (iv Interval) Merge(other Interval) Interval {
	return Interval{
		Start: min(iv.Start, other.Start),
		End:   max(iv.End, other.End),
	}
}
func CollapseIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	ints := intervals
	result := []Interval{}

	for {
		mergesCount := 0
		cursor := ints[0]
		for i := 1; i < len(ints); i++ {
			if cursor.Overlaps(ints[i]) {
				cursor = cursor.Merge(ints[i])
			} else {
				result = append(result, cursor)
				cursor = ints[i]
			}
		}
		result = append(result, cursor)

		if mergesCount == 0 {
			break
		}

		ints = result
		result = []Interval{}
	}

	return result
}

type Queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{l: list.New()}
}

func (q *Queue[T]) Enqueue(v T) {
	q.l.PushBack(v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T

	e := q.l.Front()
	if e == nil {
		return zero, false
	}

	v, ok := e.Value.(T)
	if !ok {
		return zero, false
	}

	q.l.Remove(e)
	return v, true
}

func (q *Queue[T]) Len() int {
	return q.l.Len()
}

// UF
type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(a int) int {
	if uf.parent[a] != a {
		uf.parent[a] = uf.Find(uf.parent[a])
	}
	return uf.parent[a]
}

func (uf *UnionFind) Union(a, b int) {
	ra := uf.Find(a)
	rb := uf.Find(b)
	if ra == rb {
		return
	}
	if uf.size[ra] < uf.size[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	uf.size[ra] += uf.size[rb]
}
