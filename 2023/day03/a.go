package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partNumber(data *[]string, line, start, end int) int {
	searchStart := max(start-1, 0)
	searchEnd := min(end+1, len((*data)[0])-1)

	for i := max(line-1, 0); i <= min(line+1, len(*data)-1); i++ {
		for j := searchStart; j <= searchEnd; j++ {
			char := (*data)[i][j]
			if !('0' <= char && char <= '9') && char != '.' {
				number, err := strconv.Atoi(string((*data)[line][start : end+1]))
				if err != nil {
					log.Fatalf("Could not convert part number %s to int", string((*data)[line][start:end+1]))
				}
				return number
			}
		}
	}
	return 0
}

func main() {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(rawData), "\n")

	sum := 0
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
					sum += partNumber(&data, i, start, end)
				}
			} else if inNumber {
				// Get part number that terminates before end of line
				inNumber = false
				sum += partNumber(&data, i, start, end)
			}
		}
	}
	fmt.Printf("Sum of part numbers: %d\n", sum)
}
