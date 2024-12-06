// package main
//
// import "fmt"
//
// func main() {
// 	Map := [][]rune{
// 		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
// 		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
// 	}
//
// 	for _, row := range Map {
// 		fmt.Println(string(row))
// 	}
// 	fmt.Println()
//
// 	x := 0
// 	y := 0
// 	for i := range Map {
// 		for j := range Map[i] {
// 			if Map[i][j] == '^' {
// 				Map[i][j] = '.'
// 				x = i
// 				y = j
// 			}
// 		}
// 	}
//
// 	total := 0
// 	inside := true
// 	for inside {
// 		// Up
// 		// Slice Vertical
// 		slice := []rune{}
// 		for _, row := range Map {
// 			slice = append(slice, row[y])
// 		}
// 		fmt.Printf("%c\n", slice[:x])
// 		// Check Safe & +distance
// 		for i, r := range slice[:x] {
// 			if r == '#' {
// 				inside = true
// 				x = i + 1
// 			} else if r == '.' {
// 				total++
// 			}
// 		}
// 		fmt.Println("Up:", total)
//
// 		// Right
//         total = 0
// 		slice = Map[x][y+1:]
//         fmt.Printf("%c\n", slice)
// 		for i, r := range slice {
// 			if r == '.' {
// 				total++
// 				x = i + 1
// 			} else if  {
// 				inside = true
// 			}
// 		}
// 		fmt.Println("Down:", total)
//
// 		inside = false
// 		if !inside {
// 			break
// 		}
// 	}
//
// 	// fmt.Println("Dist:", dist(Map[6], position[1]))
// }
//
// func dist(slice []rune, position int) int {
// 	total := 0
// 	for i := range slice[position+1:] {
// 		if slice[i+position+1] == '#' {
// 			break
// 		} else {
// 			total += 1
// 		}
// 	}
// 	return total
// }
