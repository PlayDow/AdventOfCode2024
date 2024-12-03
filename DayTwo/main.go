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
	file, err := os.Open("number.txt")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		var numbers []int

		for _, elem := range elements {
			num, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatalf("Error : %v", err)
			}
			numbers = append(numbers, num)
		}

		result = append(result, numbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error : %v", err)
	}

	//var totalNumber int = 0
	var totalSafe int = 0

	maxDiff := 3

	for _, line := range result {
		if isArithmeticSequence(line, maxDiff) {
			totalSafe++
		} else if Safe(line, maxDiff) || Safe(Reverse(line), maxDiff) {
			totalSafe++
		}
	}

	fmt.Print(totalSafe)
}

func isArithmeticSequence(numbers []int, maxDiff int) bool {
	if len(numbers) < 2 {
		return false
	}

	increas := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		currentDiff := numbers[i] - numbers[i-1]

		if currentDiff == 0 || Change(currentDiff) > maxDiff {
			return false
		}

		if (increas && currentDiff <= 0) || (!increas && currentDiff >= 0) {
			return false
		}
	}

	return true
}

func Safe(numbers []int, maxDiff int) bool {
	if len(numbers) < 2 {
		return false
	}

	var delete bool
	var stop bool
	var currentDiff int

	increas := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {

		if delete {
			currentDiff = numbers[i] - numbers[i-2]
			delete = false
			stop = true
		} else {
			currentDiff = numbers[i] - numbers[i-1]
		}

		if currentDiff == 0 || Change(currentDiff) > maxDiff {
			if stop {
				return false
			} else {
				delete = true
				continue
			}
		}

		if (increas && currentDiff <= 0) || (!increas && currentDiff >= 0) {
			if stop {
				return false
			} else {
				delete = true
				continue
			}
		}
	}

	return true
}

func Change(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Reverse(arr []int) []int {
	reversed := make([]int, len(arr))
	for i, j := 0, len(arr)-1; j >= 0; i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}
	return reversed
}
