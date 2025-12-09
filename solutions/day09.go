package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day09 struct{}

func (d Day09) Day() string {
	return "09"
}

func parse09(input string) []Point {
	lines := strings.Split(input, "\n")

	result := []Point{}
	for _, line := range lines {
		nums := strings.Split(line, ",")
		result = append(result, Point{X: parseInt(nums[0]), Y: parseInt(nums[1])})
	}

	return result
}

func (d Day09) Execute1(input string) string {
	ps := parse09(input)

	max := 0
	for i := 0; i < len(ps)-1; i++ {
		for j := i + 1; j < len(ps); j++ {
			area := abs(ps[i].X-ps[j].X+1) * abs(ps[i].Y-ps[j].Y+1)
			if area > max {
				max = area
			}
		}
	}

	return strconv.Itoa(max)
}

func hasPointsInside(ps []Point, p1, p2 Point) bool {
	maxY := max(p1.Y, p2.Y)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	minX := min(p1.X, p2.X)

	for _, p := range ps {
		if p.X >= maxX || p.X <= minX || p.Y >= maxY || p.Y <= minY {
			continue
		}

		return true
	}

	return false
}

func findRect(ps []Point) int {
	max := 0
	for i := 0; i < len(ps)-1; i++ {
		for j := i + 1; j < len(ps); j++ {
			area := abs(ps[i].X-ps[j].X+1) * abs(ps[i].Y-ps[j].Y+1)
			if area > max && !hasPointsInside(ps, ps[i], ps[j]) {
				fmt.Println(ps[i], ps[j])
				max = area
			}
		}
	}

	return max
}

func (d Day09) Execute2(input string) string {
	topY := 50049
	bottomY := 48719

	ps := parse09(input)

	top := []Point{}
	bottom := []Point{}
	for _, p := range ps {
		if p.Y >= topY {
			top = append(top, p)
		} else if p.Y <= bottomY {
			bottom = append(bottom, p)
		}
	}

	maxTop := findRect(top)
	maxBottom := findRect(bottom)

	max := max(maxBottom, maxTop)

	return strconv.Itoa(max)
}
