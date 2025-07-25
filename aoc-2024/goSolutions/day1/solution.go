package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processFile(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	firstList := []int{}
	secondList := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")
		firstNum, err1 := strconv.Atoi(numbers[0])
		if err1 != nil {
			log.Fatalf("Failed to parse number: %v", err1)
		} else {
			firstList = append(firstList, firstNum)
		}

		secondNum, err2 := strconv.Atoi(numbers[3])
		if err2 != nil {
			log.Fatalf("Failed to parse number: %v", err2)
		} else {
			secondList = append(secondList, secondNum)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return firstList, secondList
}

// Assumptions:
// - The length of input list is always an even number so we can split it into 2 equal lists
// - Elements of the input list are integers

func calculateDistance(thisList []int, thatList []int) int {
	sort.Ints(thisList)
	sort.Ints(thatList)

	totalDistance := 0

	for i, num := range thisList {
		otherNum := thatList[i]

		if num <= otherNum {
			totalDistance += (otherNum - num)
		} else {
			totalDistance += (num - otherNum)
		}
	}

	return totalDistance
}

func countFrequencies(intList []int) map[int]int {
	result := make(map[int]int)

	for _, val := range intList {
		_, exists := result[val]

		if exists {
			result[val]++
		} else {
			result[val] = 1
		}
	}

	return result
}

func similarityScore(thisList []int, thatList []int) int {
	frequencies := countFrequencies(thatList)
	totalScore := 0

	for _, num := range thisList {
		frequency, exists := frequencies[num]

		if exists {
			totalScore = totalScore + num*frequency
		}
	}

	return totalScore
}

func main() {
	firstList, secondList := processFile("input.txt")

	fmt.Println("Part one answer: ", calculateDistance(firstList, secondList))

	fmt.Println("Part two answer: ", similarityScore(firstList, secondList))
}
