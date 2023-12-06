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

type Card struct {
	matches int
	count   int
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var id int
	cards := make(map[int]*Card)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		idPart := strings.Split(parts[0], " ")
		id, err = strconv.Atoi(idPart[len(idPart)-1])
		if err != nil {
			log.Fatalf("Could not convert game ID %s to int", strings.Split(parts[0], " ")[1])
		}
		cards[id] = &Card{matches: 0, count: 1}

		game := strings.Split(parts[1], "|")
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

		for _, num := range strings.Split(game[1], " ") {
			if num == "" {
				continue
			}
			playerNumber, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Could not convert player number %s to int", num)
			}
			if slices.Contains(winningNumbers, playerNumber) {
				cards[id].matches++
			}
		}
	}

	total := 0
	for i := 1; i <= id; i++ {
		card, ok := cards[i]
		if ok {
			for j := i + 1; j <= i+card.matches; j++ {
				addTo := j
				if addTo > id {
					addTo = id
				}
				_, ok2 := cards[addTo]
				if ok2 {
					cards[addTo].count += card.count
				} else {
					cards[addTo] = &Card{matches: 0, count: card.count}
				}
			}
			total += card.count
		}
	}
	fmt.Printf("Total cards: %d\n", total)
}
