package main

import (
	"reflect"
	"testing"
)

func TestCountFrequencies(t *testing.T) {
	input := []int{4, 3, 5, 3, 9, 3}

	expected := map[int]int{
		3: 3,
		4: 1,
		5: 1,
		9: 1,
	}

	result := countFrequencies(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("countFrequencies(%v) = %v; want %v", input, result, expected)
	}
}

func TestSimilarityScore(t *testing.T) {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}

	expected := 31

	result := similarityScore(leftList, rightList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test similarityScore; expected %v; actual %v", expected, result)
	}
}
