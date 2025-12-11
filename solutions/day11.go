package solutions

import (
	"strconv"
	"strings"
)

type Day11 struct{}

func (d Day11) Day() string {
	return "11"
}

func parse11(input string) map[string][]string {
	lines := strings.Split(input, "\n")

	result := map[string][]string{}
	for _, l := range lines {
		parts := strings.Split(l, ":")
		key := parts[0]
		values := strings.Fields(parts[1])
		result[key] = values
	}

	return result
}

func findPathsCount(g map[string][]string, start string, end string) int {
	q := NewQueue[string]()
	q.Enqueue(start)

	count := 0

	for {
		state, ok := q.Dequeue()
		if !ok {
			break
		}

		if state == end {
			count++
			continue
		}

		vals, ok := g[state]
		if !ok {
			println("Not found", state)
			continue
		}

		for _, v := range vals {
			q.Enqueue(v)
		}
	}

	return count
}

func (d Day11) Execute1(input string) string {
	g := parse11(input)
	c := findPathsCount(g, "you", "out")

	return strconv.Itoa(c)
}

func (d Day11) Execute2(input string) string {
	g := parse11(input)
	cache := map[string]int{}

	var dfs func(node string, fft bool, dac bool) int
	dfs = func(node string, fft bool, dac bool) int {
		key := node + strconv.FormatBool(fft) + strconv.FormatBool(dac)
		cached, found := cache[key]
		if found {
			//println(key)
			return cached
		}

		if node == "out" {
			val := 0
			if fft && dac {
				val = 1
			}
			cache[key] = val
			return val
		}

		newFft := fft || node == "fft"
		newDac := dac || node == "dac"

		vals, _ := g[node]
		sum := 0
		for _, v := range vals {
			sum += dfs(v, newFft, newDac)
		}
		cache[key] = sum

		return sum
	}

	c := dfs("svr", false, false)

	return strconv.Itoa(c)
}
