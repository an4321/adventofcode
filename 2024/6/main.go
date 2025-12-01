package main

import "fmt"

// Could't do it.
func main() {
	Map := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	}

	x := 0
	y := 0
	for i := range Map {
		for j := range Map[i] {
			fmt.Printf(" %c", Map[i][j])
			if Map[i][j] == '^' {
				Map[i][j] = '.'
				x = i
				y = j
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n^ (%d,%d)\n", x, y)

	inside := true
	path := [][2]int{}
	for inside {
		inside = false
		slice := make([]rune, len(Map))

		// UP
		// Reversed List
		for i, row := range Map {
            if len(Map) > i+2 {
                slice[len(Map)-1-i] = row[y]
            } else {
                inside = false
            }
		}
		// fmt.Printf("%c\n", slice)
		// Check for #
		for _, r := range slice {
			if r == '#' {
				inside = true
			}
		}
		if !inside {
			break
		}
		// Add path
		fmt.Printf("\nActive: %c\n", slice[len(Map)-x:])
		for i, r := range slice[len(Map)-x:] {
			if r == '.' {
				path = append(path, [2]int{x - i - 1, y})
			} else {
				x = x - i
				break
			}
		}
		fmt.Printf("UP (%d,%d)\n", x, y)
		// End UP

		// Right
		inside = false
		slice = Map[x][y+1:]
		for _, r := range slice {
			if r == '#' {
				inside = true
			}
		}
		if !inside {
			break
		}
		// Add path
		fmt.Printf("\nActive: %c\n", slice)
		for i, r := range slice {
			if r == '.' {
				path = append(path, [2]int{x, y + i + 1})
			} else {
				y = y + i
				break
			}
		}
		fmt.Printf("RIGHT (%d,%d)\n", x, y)
		// End Right

		// Down
		inside = false
		slice = []rune{}
		// 4,6
		for _, row := range Map {
			slice = append(slice, row[y])
		}
		// fmt.Printf("\n🀄: %c\n", slice)
		for i := range Map {
			fmt.Printf("%c\n", Map[i])
		}
		for _, r := range slice {
			if r == '#' {
				inside = true
			}
		}
		if !inside {
			break
		}
		// Add path
		fmt.Printf("\nActive: %c\n", slice[x+1:])
		for i, r := range slice[x+1:] {
			if r == '.' {
				path = append(path, [2]int{x + i + 1, y})
			} else {
				x = x + i
				break
			}
		}
		fmt.Printf("Down (%d,%d)\n", x, y)
		// End Down

		// Left
		for i := range Map {
			fmt.Printf("%c\n", Map[i])
		}
		inside = false
		// slice = Map[x][:y]
		copy(slice, Map[x][:y])
		// Flip it
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		for _, r := range slice {
			if r == '#' {
				inside = true
			}
		}
		if !inside {
			break
		}
		// Add path
		fmt.Printf("\nActive: %c\n", slice)
		for i, r := range slice {
			if r == '.' {
				path = append(path, [2]int{x, y - i - 1})
			} else {
				y = y - i
				break
			}
		}
		fmt.Printf("Left (%d,%d)\n", x, y)
		// End Left

		// inside = false
		if !inside {
			break
		}
	}

	fmt.Println("\n", path)

	uniqueElements := make(map[string]bool)

	// Iterate over the slice and add elements to the map
	for _, pair := range path {
		key := fmt.Sprintf("%d-%d", pair[0], pair[1])
		uniqueElements[key] = true
	}

	// Get the total number of unique elements
	totalUniqueElements := len(uniqueElements)

	fmt.Println("Total unique elements:", totalUniqueElements)
}
