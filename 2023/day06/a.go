package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isWholeNumber(num float64) int {
	if num == math.Trunc(num) {
		return 1
	}
	return 0
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Could not convert string %s to int", str)
	}
	return num
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(data), "\n")

	times := strings.Fields(strings.Split(parts[0], ":")[1])
	records := strings.Fields(strings.Split(parts[1], ":")[1])
	races := make(map[int]int)
	for i := range times {
		races[stringToInt(times[i])] = stringToInt(records[i])
	}

	total := 1
	for time, record := range races {
		recordDiff := math.Sqrt(math.Pow(float64(time), 2.0) - 4.0*float64(record))
		minHold := (float64(time) - recordDiff) / 2.0
		maxHold := (float64(time) + recordDiff) / 2.0
		numCombinations := int(maxHold) - int(math.Ceil(minHold)) + 1 - isWholeNumber(minHold) - isWholeNumber(maxHold)
		total *= numCombinations
	}
	fmt.Printf("Winning combinations: %d\n", total)
}
