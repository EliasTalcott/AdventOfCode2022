package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	numbers := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	var keys []string
	for _, key := range reflect.ValueOf(numbers).MapKeys() {
		keys = append(keys, key.Interface().(string))
	}

	sum := 0
	for scanner.Scan() {
		var digit1, digit2 byte
		line := []byte(scanner.Text())
		for i, char := range line {
			if '0' <= char && char <= '9' {
				if digit1 == 0 {
					digit1 = char
				}
				digit2 = char
			} else {
				for j := range line[i:] {
					str := string(line[i : j+i+1])
					for _, key := range keys {
						if key == str {
							if digit1 == 0 {
								digit1 = numbers[str]
							}
							digit2 = numbers[str]
						}
					}
				}
			}
		}
		sum += ((int(digit1) - 48) * 10) + int(digit2) - 48
	}
	fmt.Printf("Sum of calibration values: %d\n", sum)
}
