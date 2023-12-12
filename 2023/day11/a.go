package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type galaxy struct {
	row int
	col int
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func sumTo(values []int, to int) int {
	sum := 0
	for i := 0; i < to; i++ {
		sum += values[i]
	}
	return sum
}

func findGalaxies(grid [][]byte) []galaxy {
	// Check for occupied rows and columns
	rowOccupied := make([]int, len(grid))
	colOccupied := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				rowOccupied[i] = 1
				colOccupied[j] = 1
			}
		}
	}

	// Find galaxies accounting for space expansion
	galaxies := []galaxy{}
	for i, row := range grid {
		for j, char := range row {
			if char == '#' {
				rowExpansion := (i - sumTo(rowOccupied, i)) * (2 - 1)
				colExpansion := (j - sumTo(colOccupied, j)) * (2 - 1)
				galaxies = append(galaxies, galaxy{row: i + rowExpansion, col: j + colExpansion})
			}
		}
	}
	return galaxies
}

func galaxyDistance(a, b galaxy) int {
	return abs(b.col-a.col) + abs(b.row-a.row)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	galaxies := findGalaxies(bytes.Split(data, []byte{'\n'}))
	total := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += galaxyDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Printf("Total distance of shortest paths between galaxies: %d\n", total)
}
