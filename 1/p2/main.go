package main

import (
	"bufio"
	"fmt"
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

	// println(similarity_score([]int{3,4,2,1,3,3}, []int{4,3,5,3,9,3}))
	// SOLLUTION: 23177084
	println(similarity_score(left, right))
}

func similarity_score(left, right []int) int {
	var sim int = 0

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				sim = sim + left[i]
			}
		}
	}

	return int(sim)
}
