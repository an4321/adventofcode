package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix := [][]rune{}

	// Get Input
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Part 1: 2583
	fmt.Println("PART 1:", part1(matrix))
	fmt.Println("PART 2:", part2(matrix))
}

func part1(matrix [][]rune) int {
	total := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'X' {
				// East
				if len(matrix[i][j:]) > 3 {
					if matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
						total++
					}
				}
				// West
				if len(matrix[i][:j])+1 > 3 {
					if matrix[i][j-1] == 'M' && matrix[i][j-2] == 'A' && matrix[i][j-3] == 'S' {
						total++
					}
				}
				// North
				if i+1 > 3 {
					if matrix[i-1][j] == 'M' && matrix[i-2][j] == 'A' && matrix[i-3][j] == 'S' {
						total++
					}
				}
				// South
				if len(matrix)-i > 3 {
					if matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
						total++
					}
				}
				// North East
				if i+1 > 3 && len(matrix[i][j:]) > 3 {
					if matrix[i-1][j+1] == 'M' && matrix[i-2][j+2] == 'A' && matrix[i-3][j+3] == 'S' {
						total++
					}
				}
				// North West
				if i+1 > 3 && len(matrix[i][:j])+1 > 3 {
					if matrix[i-1][j-1] == 'M' && matrix[i-2][j-2] == 'A' && matrix[i-3][j-3] == 'S' {
						total++
					}
				}
				// South East
				if len(matrix)-i > 3 && len(matrix[i][j:]) > 3 {
					if matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
						total++
					}
				}
				// South West
				if len(matrix)-i > 3 && len(matrix[i][:j])+1 > 3 {
					if matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
						total++
					}
				}
			}
		}
	}
	return total
}

func part2(matrix [][]rune) int {
	total := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'A' {
				if i > 0 && j > 0 && i+1 < len(matrix) && j+1 < len(matrix[i]) {
					// Not proud of this.
					if (matrix[i-1][j-1] == 'M' && matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'S' && matrix[i+1][j+1] == 'S') ||
						(matrix[i-1][j-1] == 'M' && matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'M' && matrix[i+1][j+1] == 'S') ||
						(matrix[i-1][j-1] == 'S' && matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'S' && matrix[i+1][j+1] == 'M') ||
						(matrix[i-1][j-1] == 'S' && matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'M' && matrix[i+1][j+1] == 'M') {
						total++
					}
				}
			}
		}
	}
	return total
}
