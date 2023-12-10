package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	row  int
	col  int
	pipe byte
}

const UP = 'U'
const RIGHT = 'R'
const DOWN = 'D'
const LEFT = 'L'
const INVALID = '0'

var allowsUp = []byte{'S', '|', 'L', 'J'}
var allowsRight = []byte{'S', '-', 'L', 'F'}
var allowsDown = []byte{'S', '|', '7', 'F'}
var allowsLeft = []byte{'S', '-', 'J', '7'}

var verticalPipes = []byte{'|', '7', 'F'}

func checkNeighbors(grid []string, current *coordinate, previous byte) (*coordinate, byte, error) {
	if previous != DOWN &&
		current.row > 0 &&
		slices.Contains(allowsUp, current.pipe) &&
		slices.Contains(allowsDown, grid[current.row-1][current.col]) {
		return &coordinate{row: current.row - 1, col: current.col, pipe: grid[current.row-1][current.col]}, UP, nil
	}
	if previous != LEFT &&
		current.col < len(grid[current.row]) &&
		slices.Contains(allowsRight, current.pipe) &&
		slices.Contains(allowsLeft, grid[current.row][current.col+1]) {
		return &coordinate{row: current.row, col: current.col + 1, pipe: grid[current.row][current.col+1]}, RIGHT, nil
	}
	if previous != UP &&
		current.row < len(grid) &&
		slices.Contains(allowsDown, current.pipe) &&
		slices.Contains(allowsUp, grid[current.row+1][current.col]) {
		return &coordinate{row: current.row + 1, col: current.col, pipe: grid[current.row+1][current.col]}, DOWN, nil
	}
	if previous != RIGHT &&
		current.col > 0 &&
		slices.Contains(allowsLeft, current.pipe) &&
		slices.Contains(allowsRight, grid[current.row][current.col-1]) {
		return &coordinate{row: current.row, col: current.col - 1, pipe: grid[current.row][current.col-1]}, LEFT, nil
	}
	return current, INVALID, errors.New(fmt.Sprintf("Could not move from %s at row %d column %d", string(current.pipe), current.row, current.col))
}

func followPath(grid []string, start *coordinate) (map[string]byte, error) {
	coordinates := map[string]byte{fmt.Sprintf("%d,%d", start.row, start.col): start.pipe}
	current := start
	var previous byte
	var err error
	for {
		current, previous, err = checkNeighbors(grid, current, previous)
		if err != nil {
			return nil, err
		}
		coordinates[fmt.Sprintf("%d,%d", current.row, current.col)] = current.pipe
		if current.pipe == 'S' {
			break
		}
	}
	return coordinates, nil
}

func countIncludedTiles(grid []string, loop map[string]byte) int {
	count := 0
	for i, row := range grid[1 : len(grid)-1] {
		verticalPipeCount := 0
		for j, char := range []byte(row[:len(row)-1]) {
			_, inLoop := loop[fmt.Sprintf("%d,%d", i+1, j)]
			if inLoop {
				if slices.Contains(verticalPipes, char) {
					verticalPipeCount++
				}
			} else if verticalPipeCount%2 == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := strings.Split(string(data), "\n")

	var start coordinate
out:
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == 'S' {
				var pipe byte
				up := i > 0 && slices.Contains(allowsDown, grid[i-1][j])
				right := j < len(grid[0]) && slices.Contains(allowsLeft, grid[i][j+1])
				down := i < len(grid) && slices.Contains(allowsUp, grid[i+1][j])
				left := j > 0 && slices.Contains(allowsRight, grid[i][j-1])
				switch {
				case up && right:
					pipe = 'L'
				case up && down:
					pipe = '|'
				case up && left:
					pipe = 'J'
				case right && down:
					pipe = 'F'
				case right && left:
					pipe = '-'
				case down && left:
					pipe = '7'
				default:
					pipe = '0'
				}
				start = coordinate{row: i, col: j, pipe: pipe}
				break out
			}
		}
	}
	loop, err := followPath(grid, &start)
	if err != nil {
		log.Fatal(err)
	}
	includedTiles := countIncludedTiles(grid, loop)
	fmt.Printf("Tiles enclosed by loop: %d\n", includedTiles)
}
