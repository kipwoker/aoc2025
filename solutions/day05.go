package solutions

import (
	"strconv"
	"strings"
)

type Day05 struct{}

func (d Day05) Day() string {
	return "05"
}

type Ingredients struct {
	Fresh       []Interval
	Experiments []int
}

func parse05Input(input string) Ingredients {
	lines := strings.Split(input, "\n")

	intervals := make([]Interval, 0, 200)
	experiments := make([]int, 0, 1100)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			intervals = append(intervals, Interval{Start: start, End: end})
			continue
		}
		value, _ := strconv.Atoi(line)
		experiments = append(experiments, value)
	}

	return Ingredients{
		Fresh:       intervals,
		Experiments: experiments,
	}
}

func (d Day05) Execute1(input string) string {
	rawIngs := parse05Input(input)
	ings := Ingredients{
		Fresh:       CollapseIntervals(rawIngs.Fresh),
		Experiments: rawIngs.Experiments,
	}

	counter := 0
	for _, exp := range ings.Experiments {
		for _, iv := range ings.Fresh {
			if iv.Inside(exp) {
				counter++
			}
		}
	}

	return strconv.Itoa(counter)
}

func (d Day05) Execute2(input string) string {
	rawIngs := parse05Input(input)
	ivs := CollapseIntervals(rawIngs.Fresh)

	sum := 0
	for _, iv := range ivs {
		sum += iv.End - iv.Start + 1
	}

	return strconv.Itoa(sum)
}
