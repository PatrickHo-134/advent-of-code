package main

import (
	"testing"
)

var exampleReports = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestLevelsAreAllIncreasing(t *testing.T) {
	expectedResults := []bool{
		false,
		true,
		false,
		false,
		false,
		true,
	}

	for i, report := range exampleReports {
		result := levelsAreAllIncreasing(report)

		if result != expectedResults[i] {
			t.Errorf("levelsAreAllIncreasing(%v) = %v; want %v", report, result, expectedResults[i])
		}
	}
}

func TestLevelsAreAllDecreasing(t *testing.T) {
	expectedResults := []bool{
		true,
		false,
		true,
		false,
		false,
		false,
	}

	for i, report := range exampleReports {
		result := levelsAreAllDecreasing(report)

		if result != expectedResults[i] {
			t.Errorf("levelsAreAllDecreasing(%v) = %v; want %v", report, result, expectedResults[i])
		}
	}
}

func TestAdjacentLevelsAreOk(t *testing.T) {
	expectedResults := []bool{
		true,
		false,
		false,
		true,
		false,
		true,
	}

	for i, report := range exampleReports {
		result := adjacentLevelsAreOk(report)

		if result != expectedResults[i] {
			t.Errorf("adjacentLevelsAreOk(%v) = %v; want %v", report, result, expectedResults[i])
		}
	}
}

func TestReportIsSafe(t *testing.T) {
	expectedResults := []bool{
		true,
		false,
		false,
		false,
		false,
		true,
	}

	for i, report := range exampleReports {
		result := reportIsSafe(report)

		if result != expectedResults[i] {
			t.Errorf("reportIsSafe(%v) = %v; want %v", report, result, expectedResults[i])
		}
	}
}

func TestCanPassProblemDampener(t *testing.T) {
	var exampleReports = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{82, 85, 87, 88, 91, 93, 97, 94},
		{24, 25, 28, 31, 28},
	}

	expectedResults := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
		true,
		true,
	}

	for i, report := range exampleReports {
		result := canPassProblemDampener(report)

		if result != expectedResults[i] {
			t.Errorf("reportIsSafe(%v) = %v; want %v", report, result, expectedResults[i])
		}
	}
}
