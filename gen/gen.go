package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func genNext(input string) string {
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	n += 1
	return fmt.Sprintf("%02d", n)
}

func replaceMain(last, next string) {
	bytes, err := os.ReadFile("../main.go")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, line)
		if strings.Contains(line, "solutions.Day"+last) {
			newLine := strings.Replace(line, "solutions.Day"+last, "solutions.Day"+next, 1)
			newLines = append(newLines, newLine)
		}
	}
	output := strings.Join(newLines, "\n")
	err = os.WriteFile("../main.go", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

func genDayFile(next string) {
	bytes, err := os.ReadFile("../solutions/day00.go")
	if err != nil {
		log.Fatal(err)
	}

	text := string(bytes)
	output := strings.ReplaceAll(text, "Day00", "Day"+next)

	err = os.WriteFile("../solutions/day"+next+".go", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

func genInputs(next string) {
	paths := []string{
		fmt.Sprintf("../inputs/%s.test.txt", next),
		fmt.Sprintf("../inputs/%s.real.txt", next),
	}
	for _, path := range paths {
		err := os.WriteFile(path, []byte(""), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	entries, err := os.ReadDir("../solutions")
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	last := entries[len(entries)-1].Name()
	last = strings.Replace(last, "day", "", 1)
	last = strings.Replace(last, ".go", "", 1)
	next := genNext(last)

	genDayFile(next)
	replaceMain(last, next)
	genInputs(next)
}
