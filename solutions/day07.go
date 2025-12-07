package solutions

import (
	"aoc2025/ext"
	"strconv"
	"strings"
)

type Day07 struct{}

func (d Day07) Day() string {
	return "07"
}

func parse07(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func getStart(grid [][]rune) Point {
	for i, xs := range grid[0] {
		if xs == 'S' {
			return Point{X: i, Y: 0}
		}
	}

	panic("Not found S")
}

func (d Day07) Execute1(input string) string {
	g := parse07(input)
	start := getStart(g)
	lastRowIdx := len(g) - 1

	q := NewQueue[Point]()
	q.Enqueue(start)

	vs := ext.New[Point]()

	splitCounter := 0

	for {
		if v, ok := q.Dequeue(); ok && v.Y < lastRowIdx {
			if vs.Has(v) {
				continue
			}

			if g[v.Y][v.X] == '^' {
				splitCounter++
				q.Enqueue(Point{X: v.X - 1, Y: v.Y + 1})
				q.Enqueue(Point{X: v.X + 1, Y: v.Y + 1})
			} else {
				q.Enqueue(Point{X: v.X, Y: v.Y + 1})
			}

			vs.Add(v)
		} else {
			break
		}
	}

	return strconv.Itoa(splitCounter)
}

func calculatePaths(g [][]rune, start Point, cache map[Point]int) int {
	lastRowIdx := len(g) - 1

	q := NewQueue[Point]()
	q.Enqueue(start)

	counter := 0

	for {
		if v, ok := q.Dequeue(); ok {
			if v.Y >= lastRowIdx {
				counter++

				continue
			}

			nextPoint := Point{X: v.X, Y: v.Y + 1}

			if g[nextPoint.Y][nextPoint.X] == '^' {
				weight, ok := cache[nextPoint]
				if ok {
					counter += weight
					continue
				}

				q.Enqueue(Point{X: nextPoint.X - 1, Y: nextPoint.Y})
				q.Enqueue(Point{X: nextPoint.X + 1, Y: nextPoint.Y})
			} else {
				q.Enqueue(nextPoint)
			}
		} else {
			break
		}
	}

	return counter
}

func (d Day07) Execute2(input string) string {
	g := parse07(input)
	lastRowIdx := len(g) - 1

	cache := make(map[Point]int)
	lastWeight := 0

	for i := lastRowIdx; i >= 0; i-- {
		for j, cell := range g[i] {
			if cell == '^' {
				p := Point{X: j, Y: i}
				weight := calculatePaths(g, Point{X: j, Y: i - 1}, cache)
				cache[p] = weight
				lastWeight = weight
			}
		}
	}

	return strconv.Itoa(lastWeight)
}
