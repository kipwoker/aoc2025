package solutions

import (
	"strconv"
	"strings"
)

type Day02 struct{}

func (d Day02) Day() string {
	return "02"
}

type Range struct {
	Start int
	End   int
}

func parse02(input string) []Range {
	rawRanges := strings.Split(input, ",")
	ranges := make([]Range, 0, len(rawRanges))
	for _, rawRange := range rawRanges {
		parts := strings.Split(rawRange, "-")
		if len(parts) != 2 {
			continue
		}
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

func commonSolution02(input string, part int) string {
	ranges := parse02(input)

	sum := 0
	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if part == 1 && check01(i) {
				sum += i
			} else if part == 2 && check02(i) {
				sum += i
			}
		}
	}

	return strconv.Itoa(sum)
}

func check01(i int) bool {
	s := strconv.Itoa(i)
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func check02(i int) bool {
	s := strconv.Itoa(i)
	for j := 2; j <= len(s); j++ {
		if len(s)%j != 0 {
			continue
		}

		partSize := len(s) / j
		pattern := s[0:partSize]
		hasMatch := true
		for k := partSize; k+partSize <= len(s); k += partSize {
			part := s[k : k+partSize]
			if part != pattern {
				hasMatch = false
				break
			}
		}
		if hasMatch {
			return true
		}
	}

	return false
}

func (d Day02) Execute1(input string) string {
	return commonSolution02(input, 1)
}

func (d Day02) Execute2(input string) string {
	return commonSolution02(input, 2)
}
