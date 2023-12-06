package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(strings.Split(line, ":")[1], "|")

		winningNumbers := []int{}
		for _, num := range strings.Split(game[0], " ") {
			if num == "" {
				continue
			}
			winningNumber, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Could not convert winning number %s to int", num)
			}
			winningNumbers = append(winningNumbers, winningNumber)
		}

		score := 0
		for _, num := range strings.Split(game[1], " ") {
			if num == "" {
				continue
			}
			playerNumber, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Could not convert player number %s to int", num)
			}
			if slices.Contains(winningNumbers, playerNumber) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score
	}
	fmt.Printf("Total points: %d\n", total)
}
