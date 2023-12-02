package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		var digit1, digit2 byte
		line := scanner.Text()
		for _, char := range []byte(line) {
			if '0' <= char && char <= '9' {
				if digit1 == 0 {
					digit1 = char
				}
				digit2 = char
			}
		}
		sum += ((int(digit1) - 48) * 10) + int(digit2) - 48
	}
	fmt.Printf("Sum of calibration values: %d\n", sum)
}
