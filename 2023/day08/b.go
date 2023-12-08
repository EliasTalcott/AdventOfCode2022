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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums []int) int {
	currentGcd := (nums[0] * nums[1]) / gcd(nums[0], nums[1])
	for i := 1; i < len(nums); i++ {
		currentGcd = (currentGcd * nums[i]) / gcd(currentGcd, nums[i])
	}
	return currentGcd
}

func pathLength(start string, directions string, nodes map[string]*Node) int {
	current := start
	next := start
	steps := 0
	for {
		for _, direction := range directions {
			steps++
			if direction == 'L' {
				next = nodes[current].left
			} else {
				next = nodes[current].right
			}
			if next[2] == 'Z' {
				break
			}
			current = next
		}
		if next[2] == 'Z' {
			break
		}
	}
	return steps
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	directions := lines[0]
	nodes := make(map[string]*Node)
	startingNodes := []string{}
	for _, line := range lines[2:] {
		parts := strings.Fields(line)
		identifier := parts[0]
		if identifier[2] == 'Z' {
			startingNodes = append(startingNodes, identifier)
		}
		left := parts[2][1:4]
		right := parts[3][0:3]
		nodes[identifier] = &Node{left: left, right: right}
	}

	pathLengths := make([]int, len(startingNodes))
	for i, startingNode := range startingNodes {
		pathLengths[i] = pathLength(startingNode, directions, nodes)
	}
	fmt.Printf("Steps for all starting nodes to arrive at nodes ending in Z: %d\n", lcm(pathLengths))
}
