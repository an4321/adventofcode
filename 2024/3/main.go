package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Data
	raw_data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := string(raw_data)

	// Sample Data
	// data = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	// data = "mul(2,4) don't() mul(2,2) do() don't() mul(2,4)"

	// SOLLUTION: 175015740
	fmt.Println("PART 1:", part1(data))
	// SOLLUTION: 112272912
	fmt.Println("PART 2:", part2(data))
}

func sigmax(slice [][]int) int {
	total_mul := 0
	for i := range slice {
		total_mul += slice[i][0] * slice[i][1]
	}
	return total_mul
}

func part1(data string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(data, -1)
	var result [][]int

	// Convert matches to int
	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result = append(result, []int{num1, num2})
		}
	}

	return sigmax(result)
}

func part2(data string) int {
	safe_data := regexp.MustCompile(`don't\(\).*?do\(\)`).ReplaceAllString(data, "")
	safe_data = regexp.MustCompile(`(.*don't\(\)).*`).ReplaceAllString(safe_data, `$1`)
	return part1(safe_data)
}
