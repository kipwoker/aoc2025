package solutions

import (
	"strconv"
	"strings"
)

type Day06 struct{}

func (d Day06) Day() string {
	return "06"
}

type Problem struct {
	Nums []int
	Op   string
}

func parse06p1(input string) []Problem {
	lines := strings.Split(input, "\n")

	m := len(lines)
	var n int

	matrix := [][]string{}
	for _, line := range lines {
		fields := strings.Fields(line)
		matrix = append(matrix, fields)
		n = len(fields)
	}

	result := []Problem{}
	for i := 0; i < n; i++ {
		nums := []int{}
		for j := 0; j < m-1; j++ {
			num, _ := strconv.Atoi(matrix[j][i])
			nums = append(nums, num)
		}
		op := matrix[m-1][i]
		result = append(result, Problem{
			Nums: nums,
			Op:   op,
		})
	}

	return result
}

func digitsToNumber(digits []rune) int {
	numStr := string(digits)
	num, _ := strconv.Atoi(numStr)
	return num
}

func parse06p2(input string) []Problem {
	lines := strings.Split(input, "\n")
	m := len(lines)
	var n int

	matrix := make([][]rune, len(lines))
	for i, s := range lines {
		chars := []rune(s)
		matrix[i] = chars
		n = len(chars)
	}

	result := []Problem{}
	nums := []int{}
	var op string
	for i := 0; i < n; i++ {
		allSpaces := true
		digits := []rune{}

		for j := range m {
			ch := matrix[j][i]
			if ch >= '0' && ch <= '9' {
				allSpaces = false
				digits = append(digits, ch)
			} else if ch == '*' || ch == '+' {
				allSpaces = false
				op = string(ch)
			}
		}
		if allSpaces {
			result = append(result, Problem{Nums: nums, Op: op})
			nums = []int{}
			op = ""
		} else {
			num := digitsToNumber(digits)
			nums = append(nums, num)
		}
	}
	result = append(result, Problem{Nums: nums, Op: op})

	return result
}

func solve06(ps []Problem) string {
	sum := 0
	for _, pr := range ps {
		switch pr.Op {
		case "+":
			s := 0
			for _, num := range pr.Nums {
				s += num
			}
			sum += s
		case "*":
			p := 1
			for _, num := range pr.Nums {
				p *= num
			}
			sum += p
		}
	}

	return strconv.Itoa(sum)
}

func (d Day06) Execute1(input string) string {
	ps := parse06p1(input)
	return solve06(ps)
}

func (d Day06) Execute2(input string) string {
	ps := parse06p2(input)
	return solve06(ps)
}
