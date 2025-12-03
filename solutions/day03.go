package solutions

import (
	"strconv"
	"strings"
)

type Day03 struct{}

func (d Day03) Day() string {
	return "03"
}

func parse03(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		row := make([]int, 0, len(line))
		for _, ch := range line {
			digit, _ := strconv.Atoi(string(ch))
			row = append(row, digit)
		}
		grid = append(grid, row)
	}
	return grid
}

var PowersOf10 = map[int]int{
	1:  10,
	2:  100,
	3:  1000,
	4:  10000,
	5:  100000,
	6:  1000000,
	7:  10000000,
	8:  100000000,
	9:  1000000000,
	10: 10000000000,
	11: 100000000000,
	12: 1000000000000,
}

func findMax(digits []int, limit int) int {
	len := len(digits)
	max := 0
	var backtrack func(pos int, current int, depth int)
	backtrack = func(pos int, current int, depth int) {
		if depth == limit {
			if current > max {
				max = current
			}
			return
		}
		for i := pos; i < len; i++ {
			nextDepth := depth + 1
			next := current*10 + digits[i]
			nextPos := i + 1

			level := limit - nextDepth
			if level > 0 {
				maxPart := max / PowersOf10[level]
				if next < maxPart {
					continue
				}
			}

			backtrack(nextPos, next, nextDepth)
		}
	}
	backtrack(0, 0, 0)
	return max
}

func (d Day03) Execute1(input string) string {
	grid := parse03(input)

	sum := 0
	for _, row := range grid {

		max := findMax(row, 2)

		sum += max
	}

	return strconv.Itoa(sum)
}

func (d Day03) Execute2(input string) string {
	grid := parse03(input)

	sum := 0
	for _, row := range grid {
		max := findMax(row, 12)

		sum += max
	}

	return strconv.Itoa(sum)
}
