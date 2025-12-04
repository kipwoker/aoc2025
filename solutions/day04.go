package solutions

import (
	"strconv"
	"strings"
)

type Day04 struct{}

func (d Day04) Day() string {
	return "04"
}

func parse04(input string) PointSet {
	set := NewPointSet()
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, ch := range line {
			if ch == '@' {
				set.Add(Point{X: j, Y: i})
			}
		}
	}

	return set
}

func getToRemove(set PointSet) PointSet {
	toRemove := NewPointSet()
	for p := range set {
		ns := p.GetNeighbors8()
		nsCount := 0
		for _, n := range ns {
			if _, exists := set[n]; exists {
				nsCount++
			}
		}
		if nsCount < 4 {
			toRemove.Add(p)
		}
	}

	return toRemove
}

func (d Day04) Execute1(input string) string {
	set := parse04(input)
	toRemove := getToRemove(set)
	counter := toRemove.Size()

	return strconv.Itoa(counter)
}

func (d Day04) Execute2(input string) string {
	set := parse04(input)
	counter := 0

	for {
		toRemove := getToRemove(set)
		if toRemove.Size() == 0 {
			break
		}
		counter += toRemove.Size()
		for p := range toRemove {
			set.Remove(p)
		}
	}

	return strconv.Itoa(counter)
}
