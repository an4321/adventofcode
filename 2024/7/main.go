package main

import (
	"fmt"
)

func main() {
	var result [][]rune
	permutations([]rune{'+', '*'}, 4, []rune{}, &result)

	for _, slice := range result {
		fmt.Printf("%s\n", string(slice))
	}
}

func permutations(runes []rune, target_len int, current []rune, result *[][]rune) {
	if len(current) == target_len {
		temp := make([]rune, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for _, r := range runes {
		current = append(current, r)
		permutations(runes, target_len, current, result)
		current = current[:len(current)-1]
	}
}
