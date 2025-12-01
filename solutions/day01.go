package solutions

import (
	"strconv"
	"strings"
)

type Day01 struct{}

func (d Day01) Day() string {
	return "01"
}

type Instruction struct {
	Command string
	Value   int
}

func parse(input string) []Instruction {
	lines := strings.Split(input, "\n")
	ins := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		ch := line[:1]

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		ins = append(ins, Instruction{Command: ch, Value: num})
	}

	return ins
}

func isChanded(prev, current int) bool {
	return prev*current < 0
}

func (d Day01) Execute1(input string) string {
	ins := parse(input)
	pos := 50
	mod := 100

	counter := 0
	for _, instruction := range ins {
		switch instruction.Command {
		case "L":
			pos -= instruction.Value
		case "R":
			pos += instruction.Value
		}

		pos %= mod
		if pos == 0 {
			counter++
		}
	}

	return strconv.Itoa(counter)
}

func (d Day01) Execute2(input string) string {
	ins := parse(input)
	pos := 50
	mod := 100

	counter := 0
	for _, instruction := range ins {
		currentPos := pos
		switch instruction.Command {
		case "L":
			pos -= instruction.Value

		case "R":
			pos += instruction.Value
		}

		if pos > 0 {
			counter += pos / mod
		} else {
			counter += -pos / mod
		}

		if pos == 0 || isChanded(currentPos, pos) {
			counter++
		}

		pos %= mod
	}

	return strconv.Itoa(counter)
}
