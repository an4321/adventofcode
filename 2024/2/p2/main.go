package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Sample Data
	// data := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// }

	// Data
	content, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := parse_input(string(content))

	total := 0
	for i := range data {
		if is_safe(data[i]) {
			total += 1
		}
	}

	// SOLLUTION: 400
	fmt.Println("TOTAL:", total)
}

func is_safe(slice []int) bool {
	if is_sorted(slice) {
		return gol_d_lock(slice)
	} else {
		for i := range slice {
			temp := make([]int, 0, len(slice)-1)
			temp = append(temp, slice[:i]...)
			temp = append(temp, slice[i+1:]...)

			if is_sorted(temp) {
				return gol_d_lock(temp)
			}
		}
	}
	return false
}

func is_sorted(slice []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] >= slice[i+1] {
			ascending = false
		}
		if slice[i] <= slice[i+1] {
			descending = false
		}
	}

	return ascending || descending
}

func gol_d_lock(slice []int) bool {
	prev := slice[0]

	for i := range slice {
		diff := math.Abs(float64(slice[i] - prev))
		if diff <= 3 {
			prev = slice[i]
		} else {
			return false
		}
	}
	return true
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
