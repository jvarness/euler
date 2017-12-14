package main

import "fmt"

/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/
func main() {
	// Size of grid, in this case 20 x 20
	gridSize := 20

	fmt.Printf("Number of lattice paths: %v", latticePaths(gridSize))
}

func latticePaths(gridSize int) int64 {
	return traverse(0, 0, gridSize)
}

func traverse(x int, y int, gridSize int) int64 {
	if x < gridSize && y < gridSize {
		return traverse(x+1, y, gridSize) + traverse(x, y+1, gridSize)
	} else if x < gridSize {
		return traverse(x+1, y, gridSize)
	} else if y < gridSize {
		return traverse(x, y+1, gridSize)
	}
	return int64(1)
}
