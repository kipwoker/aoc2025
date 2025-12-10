package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day10 struct{}

func (d Day10) Day() string {
	return "10"
}

type Scheme struct {
	Pattern string
	Buttons [][]int
	Joltage []int
}

func parse10(input string) []Scheme {
	//[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}

	result := []Scheme{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")

		inds := strings.Trim(parts[0], "[]")

		buttons := [][]int{}
		for i := 1; i < len(parts)-1; i++ {
			btn := strings.Split(strings.Trim(parts[i], "()"), ",")

			button := []int{}
			for _, id := range btn {
				button = append(button, parseInt(id))
			}
			buttons = append(buttons, button)
		}

		last := strings.Split(strings.Trim(parts[len(parts)-1], "{}"), ",")
		joltage := []int{}
		for _, id := range last {
			joltage = append(joltage, parseInt(id))
		}

		result = append(result, Scheme{Pattern: inds, Buttons: buttons, Joltage: joltage})
	}

	return result
}

type PatternState struct {
	Snapshot string
	Depth    int
}

func press(state string, buttons []int) string {
	result := []rune(state)
	for _, idx := range buttons {
		if result[idx] == '.' {
			result[idx] = '#'
		} else {
			result[idx] = '.'
		}
	}

	return string(result)
}

func findFastWay(sc Scheme) int {
	initPattern := []rune{}
	for i := 0; i < len(sc.Pattern); i++ {
		initPattern = append(initPattern, '.')
	}
	initSnapshot := string(initPattern)
	initState := PatternState{
		Snapshot: initSnapshot,
		Depth:    0,
	}

	q := NewQueue[PatternState]()
	q.Enqueue(initState)

	min := -1

	for {
		state, hasState := q.Dequeue()
		if !hasState {
			break
		}

		if min != -1 && min <= state.Depth {
			continue
		}

		if state.Snapshot == sc.Pattern {
			if min == -1 || state.Depth < min {
				min = state.Depth
			}
			continue
		}

		for _, btn := range sc.Buttons {
			q.Enqueue(PatternState{
				Snapshot: press(state.Snapshot, btn),
				Depth:    state.Depth + 1,
			})
		}
	}

	return min
}

func press2(state []int, button []int) []int {
	result := slices.Clone(state)
	for _, idx := range button {
		result[idx]--
	}

	return result
}

func pressMany(state []int, button []int, times int) []int {
	result := slices.Clone(state)
	for _, idx := range button {
		result[idx] -= times
	}

	return result
}

func isGoodButton(button []int, includeIdx, excludeIdx int) bool {
	foundInclude := false
	foundExclude := false
	for _, idx := range button {
		if idx == includeIdx {
			foundInclude = true
		}
		if idx == excludeIdx {
			foundExclude = true
		}
	}

	return foundInclude && !foundExclude
}

func getGoodButtons(buttons [][]int, includeIdx, excludeIdx int) [][]int {
	result := [][]int{}
	for _, button := range buttons {
		if isGoodButton(button, includeIdx, excludeIdx) {
			result = append(result, button)
		}
	}

	return result
}

func isFailedState(state []int) bool {
	for _, s := range state {
		if s < 0 {
			return true
		}
	}

	return false
}

type Press struct {
	ButtonIdx int
	Count     int
}

func findFastWay2(sc Scheme) int {
	res := math.MaxInt
	n := len(sc.Joltage)

	index := map[int][]int{}
	for i, b := range sc.Buttons {
		for _, idx := range b {
			list, ok := index[idx]
			if !ok {
				list = []int{}
			}
			list = append(list, i)
			index[idx] = list
		}
	}

	solos := []Press{}
	for k, v := range index {
		if len(v) == 1 {
			solos = append(solos, Press{ButtonIdx: v[0], Count: sc.Joltage[k]})
		}
	}

	soloCursor := 0

	var backtrack func(state []int, buttons [][]int, depth int)
	backtrack = func(state []int, buttons [][]int, depth int) {
		if isFailedState(state) {
			return
		}

		limit := slices.Max(state)
		if limit == 0 {
			res = min(res, depth)
			return
		}

		if limit+depth >= res {
			return
		}

		if soloCursor < len(solos) {
			pressCtx := solos[soloCursor]
			soloCursor++
			backtrack(
				pressMany(state, sc.Buttons[pressCtx.ButtonIdx], pressCtx.Count),
				buttons,
				depth+pressCtx.Count,
			)
			return
		}

		for i := range n {
			for j := range n {
				if state[i] > state[j] {
					good := getGoodButtons(buttons, i, j)
					count := len(good)
					if count == 0 {
						return
					}

					if count == 1 {
						backtrack(
							press2(state, good[0]),
							buttons,
							depth+1,
						)
						return
					}
				}
			}
		}

		for i := range len(buttons) {
			button := buttons[i]
			tail := buttons[i:]
			backtrack(
				press2(state, button),
				tail,
				depth+1,
			)
		}
	}

	backtrack(sc.Joltage, sc.Buttons, 0)

	return res
}

func (d Day10) Execute1(input string) string {
	schemes := parse10(input)

	sum := 0
	for _, s := range schemes {
		sum += findFastWay(s)
	}

	return strconv.Itoa(sum)
}

func (d Day10) Execute2(input string) string {
	schemes := parse10(input)

	sum := 0
	for i, s := range schemes {
		val := findFastWay2(s)
		println("Scheme", i) //, "Result", val)
		sum += val
	}

	return strconv.Itoa(sum)
}
