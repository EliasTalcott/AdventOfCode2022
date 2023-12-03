package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	number int
	gear   string
}

func partNumber(data *[]string, line, start, end int) []*Part {
	searchStart := max(start-1, 0)
	searchEnd := min(end+1, len((*data)[0])-1)

	parts := []*Part{}
	for i := max(line-1, 0); i <= min(line+1, len(*data)-1); i++ {
		for j := searchStart; j <= searchEnd; j++ {
			if (*data)[i][j] == '*' {
				number, err := strconv.Atoi(string((*data)[line][start : end+1]))
				if err != nil {
					log.Fatalf("Could not convert part number %s to int", string((*data)[line][start:end+1]))
				}
				parts = append(parts, &Part{number, fmt.Sprintf("%d,%d", i, j)})
			}
		}
	}
	return parts
}

func main() {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(rawData), "\n")

	parts := []*Part{}
	for i := 0; i < len(data); i++ {
		line := data[i]
		inNumber := false
		var start, end int
		for j, char := range line {
			if '0' <= char && char <= '9' {
				// Track position of digits
				if !inNumber {
					inNumber = true
					start = j
				}
				end = j

				// Get part number that terminates at end of line
				if j == len(line)-1 {
					inNumber = false
					parts = append(parts, partNumber(&data, i, start, end)...)
				}
			} else if inNumber {
				// Get part number that terminates before end of line
				inNumber = false
				parts = append(parts, partNumber(&data, i, start, end)...)
			}
		}
	}

	// Add parts to gears that they are touching
	gears := make(map[string][]int)
	for _, part := range parts {
		list, ok := gears[part.gear]
		if ok {
			gears[part.gear] = append(list, part.number)
		} else {
			gears[part.gear] = []int{part.number}
		}
	}

	// Sum ratios of gears with exactly two parts
	sum := 0
	for _, list := range gears {
		if len(list) == 2 {
			sum += list[0] * list[1]
		}
	}
	fmt.Printf("Sum of gear ratios: %d\n", sum)
}
