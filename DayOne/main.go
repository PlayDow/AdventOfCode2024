package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("number.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// part one
	var totalNumber int = 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var tableau1 []int
	var tableau2 []int
	var choice bool = true

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error : %v", err)
		}

		if choice == true {
			tableau1 = append(tableau1, num)
			choice = false
		} else {
			tableau2 = append(tableau2, num)
			choice = true
		}
	}

	sort.Ints(tableau1)
	sort.Ints(tableau2)

	for i := range tableau1 {
		totalNumber += int(math.Abs(float64(tableau1[i] - tableau2[i])))
	}

	//fmt.Print(totalNumber)

	// part two
	var totalNumberTwo int = 0
	similarity := make(map[int]int)

	for _, number := range tableau2 {
		similarity[number]++
	}

	for _, number := range tableau1 {
		totalNumberTwo += number * similarity[number]
	}

	// error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(totalNumberTwo)
}
