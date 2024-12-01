package main

import (
	"fmt"
	"math"
	"sort"

	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) > 1 {
			leftValue, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting first part to int:", err)
				continue
			}
			left = append(left, leftValue)

			rightValue, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting second part to int:", err)
				continue
			}
			right = append(right, rightValue)
		}
	}

	// SOLLUTION: 2057374
	println(total_dist(left, right))
}

func total_dist(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var dist float64 = 0
	for i := 0; i < len(left); i++ {
		d := math.Abs(float64(left[i]) - float64(right[i]))
		dist = dist + d
	}
	return int(dist)
}
