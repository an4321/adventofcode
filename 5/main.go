package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_content, err := os.ReadFile("./input.txt")
	rules_content, err := os.ReadFile("./rules.txt")
	if err != nil {
		log.Fatal(err)
	}

	rules := parse_input(strings.ReplaceAll(string(rules_content), "|", " "))
	data := parse_input(strings.ReplaceAll(string(input_content), ",", " "))

	// rules = [][]int{
	// 	{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13}, {75, 53},
	// 	{29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53}, {61, 29}, {47, 13},
	// 	{75, 47}, {97, 75}, {47, 61}, {75, 61}, {47, 29}, {75, 13}, {53, 13},
	// }

	// data = [][]int{
	// 	{75, 47, 61, 53, 29},
	// 	{97, 61, 53, 29, 13},
	// 	{75, 29, 13},
	// 	{75, 97, 47, 61, 53},
	// 	{61, 13, 29},
	// 	{97, 13, 75, 29, 47},
	// }

    // Sol: 4637
	fmt.Println("Part1:", part1(rules, data))
}

func part1(rules [][]int, data [][]int) int {
	total := 0
	for i := range data {
		safe := true

		for j := range data[i] {
			for k := range rules {
				if rules[k][1] == data[i][j] {
					for l := range data[i][j+1:] {
						if data[i][j+1+l] == rules[k][0] {
							safe = false
						}
					}
				}
			}
		}
		if safe {
			total += data[i][len(data[i])/2]
		}
	}

	return total
}

func parse_input(input string) [][]int {
	var result [][]int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		stringNumbers := strings.Fields(line)
		var row []int

		for _, str := range stringNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}