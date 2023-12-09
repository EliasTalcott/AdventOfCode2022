package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Could not convert string %s to int\n", str)
	}
	return num
}

func reduceSequence(nums [][]int) [][]int {
	done := false
	for !done {
		currentRow := nums[len(nums)-1]
		newRow := []int{}
		done = true
		for i := 1; i < len(currentRow); i++ {
			difference := currentRow[i] - currentRow[i-1]
			if difference != 0 {
				done = false
			}
			newRow = append(newRow, difference)
		}
		nums = append(nums, newRow)
	}
	return nums
}

func nextValue(nums [][]int) int {
	nums[len(nums)-1] = append(nums[len(nums)-1], 0)
	for i := len(nums) - 1; i > 0; i-- {
		next := nums[i-1][0] - nums[i][0]
		nums[i-1] = append([]int{next}, nums[i-1]...)
	}
	return nums[0][0]
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		originalNums := []int{}
		for _, field := range strings.Fields(scanner.Text()) {
			originalNums = append(originalNums, stringToInt(field))
		}
		nums := reduceSequence([][]int{originalNums})
		total += nextValue(nums)
	}
	fmt.Printf("Total of next values: %d\n", total)
}
