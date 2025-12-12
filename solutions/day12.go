package solutions

import (
	"aoc2025/ext"
	"fmt"
	"strconv"
	"strings"
)

type Day12 struct{}

func (d Day12) Day() string {
	return "12"
}

type Field struct {
	N, M   int
	Counts []int
}

type FigureInput struct {
	Figures          []ext.Set[Point]
	MinSize, MaxSize int
	Fields           []Field
}

func parse12(input string) FigureInput {
	lines := strings.Split(input, "\n")

	figures := []ext.Set[Point]{}

	minSize := 10
	maxSize := 0

	for i := 0; i <= 5; i++ {
		start := i*5 + 1
		fig := ext.New[Point]()
		for j := start; j < start+3; j++ {
			for k := 0; k < 3; k++ {
				if lines[j][k] == '#' {
					fig.Add(Point{X: k, Y: j - start})
				}
			}
		}

		minSize = min(minSize, fig.Size())
		maxSize = max(maxSize, fig.Size())

		figures = append(figures, *fig)
	}

	fields := []Field{}

	for i := 30; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, ":")
		dims := strings.Split(parts[0], "x")
		cs := strings.Fields(parts[1])
		counts := []int{}
		for _, c := range cs {
			counts = append(counts, parseInt(c))
		}

		field := Field{N: parseInt(dims[0]), M: parseInt(dims[1]), Counts: counts}
		fields = append(fields, field)
	}

	return FigureInput{
		Figures: figures,
		Fields:  fields,
		MinSize: minSize,
		MaxSize: maxSize,
	}
}

func printSlice(slice []int) {
	fmt.Print("[")
	first := true
	for v := range slice {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(v)
		first = false
	}
	fmt.Println("]")
}

func sum(cs []int) int {
	sum := 0

	for _, c := range cs {
		sum += c
	}

	return sum
}

// if field big enough, no need to place figures, they are cool
func getMaxFieldIndexes(in FigureInput) []int {
	indexes := []int{}
	for i, f := range in.Fields {
		if f.M*f.N >= in.MaxSize*sum(f.Counts) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// if field too small, no need to place figures, they are will not fit
func getMinFieldIndexes(in FigureInput) []int {
	indexes := []int{}
	for i, f := range in.Fields {
		if f.M*f.N < in.MinSize*sum(f.Counts) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func (d Day12) Execute1(input string) string {
	in := parse12(input)
	println("MinSize", in.MinSize, "MaxSize", in.MaxSize)

	maxIdxs := getMaxFieldIndexes(in)
	println("Max", len(maxIdxs))
	printSlice(maxIdxs)
	minIdxs := getMinFieldIndexes(in)
	println("Min", len(minIdxs))
	printSlice(minIdxs)

	return strconv.Itoa(len(maxIdxs))
}

func (d Day12) Execute2(input string) string {
	return "Not Implemented: Part 2"
}
