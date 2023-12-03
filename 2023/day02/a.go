package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var RED_MAX int = 12
var GREEN_MAX int = 13
var BLUE_MAX int = 14

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			log.Fatalf("Could not convert %s to int", strings.Split(parts[0], " ")[1])
		}
		legal := true
		for _, game := range strings.Split(parts[1], ";") {
			for _, pair := range strings.Split(game, ",") {
				num, err := strconv.Atoi(strings.Split(pair, " ")[1])
				if err != nil {
					log.Fatalf("Could not convert cube count %s to int", strings.Split(pair, " ")[1])
				}
				color := strings.Split(pair, " ")[2]
				if (color == "red" && num > RED_MAX) ||
					(color == "green" && num > GREEN_MAX) ||
					(color == "blue" && num > BLUE_MAX) {
					legal = false
					break
				}
			}
			if !legal {
				break
			}
		}
		if legal {
			sum += id
		}
	}
	fmt.Printf("Sum of IDs of valid games: %d\n", sum)
}
