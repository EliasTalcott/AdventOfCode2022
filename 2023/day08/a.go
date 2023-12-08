package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	directions := lines[0]
	nodes := make(map[string]*Node)
	for _, line := range lines[2:] {
		parts := strings.Fields(line)
		identifier := parts[0]
		left := parts[2][1:4]
		right := parts[3][0:3]
		nodes[identifier] = &Node{left: left, right: right}
	}

	current := "AAA"
	next := "AAA"
	steps := 0
	for {
		for _, direction := range directions {
			steps++
			if direction == 'L' {
				next = nodes[current].left
			} else {
				next = nodes[current].right
			}
			if next == "ZZZ" {
				break
			}
			current = next
		}
		if next == "ZZZ" {
			break
		}
	}
	fmt.Printf("Total steps to get from AAA to ZZZ: %d\n", steps)
}
