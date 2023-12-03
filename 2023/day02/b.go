package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var red_max, blue_max, green_max int
		for _, game := range strings.Split(strings.Split(line, ":")[1], ";") {
			for _, pair := range strings.Split(game, ",") {
				num, err := strconv.Atoi(strings.Split(pair, " ")[1])
				if err != nil {
					log.Fatalf("Could not convert cube count %s to int", strings.Split(pair, " ")[1])
				}
				color := strings.Split(pair, " ")[2]
				if color == "red" && num > red_max {
					red_max = num
					continue
				}
				if color == "green" && num > green_max {
					green_max = num
					continue
				}
				if color == "blue" && num > blue_max {
					blue_max = num
				}
			}
		}
		sum += red_max * green_max * blue_max
	}
	fmt.Printf("Sum of powers: %d\n", sum)
}
