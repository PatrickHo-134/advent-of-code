package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func processFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		numbers := strings.Split(line, " ")
		row := []int{}

		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)

			if err != nil {
				log.Fatalf("Failed to parse number: %v", err)
			}

			row = append(row, num)
		}

		reports = append(reports, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return reports
}

func levelsAreAllIncreasing(report []int) bool {
	isIncreasing := true
	currentValue := 0

	if len(report) == 0 {
		return false
	}

	for _, num := range report {
		if num <= currentValue {
			isIncreasing = false
		} else {
			currentValue = num
		}
	}

	return isIncreasing
}

func levelsAreAllDecreasing(report []int) bool {
	if len(report) == 0 {
		return false
	}

	isDecreasing := true
	currentValue := report[0]

	for _, num := range report[1:] {
		if num >= currentValue {
			isDecreasing = false
		} else {
			currentValue = num
		}
	}

	return isDecreasing
}

func isUnsafeLevel(currentLevel int, previousLevel int) bool {
	difference := int(math.Abs(float64(currentLevel - previousLevel)))

	return difference < 1 || difference > 3
}

func adjacentLevelsAreOk(report []int) bool {
	if len(report) == 0 {
		return false
	}

	isOk := true
	previousValue := report[0]

	for _, num := range report[1:] {
		if isUnsafeLevel(num, previousValue) {
			isOk = false
		}

		previousValue = num
	}

	return isOk
}

func reportIsSafe(report []int) bool {
	return (levelsAreAllDecreasing(report) || levelsAreAllIncreasing(report)) && adjacentLevelsAreOk(report)
}

func canPassProblemDampener(report []int) bool {
	isSafeReport := false

	if reportIsSafe(report) {
		isSafeReport = true
	} else {
		for idx := range report {
			newReport := append([]int{}, report[:idx]...)
			newReport = append(newReport, report[idx+1:]...)
			if reportIsSafe(newReport) {
				isSafeReport = true
				break
			}
		}
	}

	return isSafeReport
}

func printPart1(reports [][]int) {
	countSafeReports := 0

	for _, report := range reports {
		if reportIsSafe(report) {
			countSafeReports++
		}
	}

	fmt.Println("Part 1 - Number of safe reports: ", countSafeReports)
}

func printPart2(reports [][]int) {
	countSafeReports := 0

	for _, report := range reports {
		if canPassProblemDampener(report) {
			countSafeReports++
		}
	}

	fmt.Println("Part 2 - Number of safe reports: ", countSafeReports)
}
func main() {
	reports := processFile("input.txt")

	printPart1(reports)
	printPart2(reports)

}
