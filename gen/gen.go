package gen

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

	mainText, err := os.ReadFile("../main.go")
	if err != nil {
		panic(err)
	}
	mainLines := strings.Split(string(mainText), "\n")
	var newLines []string
	for _, line := range mainLines {
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
