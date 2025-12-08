package main

import (
	"aoc2025/solutions"
	"fmt"
	"os"
	"time"
)

func printLogo() {
	data, err := os.ReadFile("logo.txt")
	if err != nil {
		panic(err)
	}

	txt := string(data)
	fmt.Println(txt)

	fmt.Println()
}

var solutionImpls = []solutions.Solution{
	solutions.Day00{},
	solutions.Day01{},
	solutions.Day02{},
	solutions.Day03{},
	solutions.Day04{},
	solutions.Day05{},
	solutions.Day06{},
	solutions.Day07{},
	solutions.Day08{},
}

func main() {
	// Configure it
	day := ""       // last by default
	label := "real" // test | real

	if day == "" {
		day = solutionImpls[len(solutionImpls)-1].Day()
	}

	if day == "00" {
		printLogo()
	}

	filePath := fmt.Sprintf("inputs/%s.%s.txt", day, label)

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(data)

	solutionMap := make(map[string]solutions.Solution)
	for _, s := range solutionImpls {
		solutionMap[s.Day()] = s
	}

	dayTitle := fmt.Sprintf("¸.•*´❄`*•.¸ Day %s %s ¸.•*´❄`*•.¸", day, label)
	fmt.Println(dayTitle)
	fmt.Println()

	fmt.Println("Part 1")
	start1 := time.Now()
	output1 := solutionMap[day].Execute1(input)
	fmt.Printf("Result %s\n", output1)
	duration1 := time.Since(start1)
	fmt.Printf("Elapsed %v\n", duration1)
	fmt.Println()

	fmt.Println("Part 2")
	start2 := time.Now()
	output2 := solutionMap[day].Execute2(input)
	fmt.Printf("Result %s\n", output2)
	duration2 := time.Since(start2)
	fmt.Printf("Elapsed %v\n", duration2)
	fmt.Println()

	fmt.Println("❄`*•.¸.•*´❄`*•.¸._.¸.•*´❄`*•.¸.•*´❄")
}
