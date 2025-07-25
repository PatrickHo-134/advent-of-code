package utils

import (
	"bufio"
	"log"
	"os"
)

func ProcessFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	content := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content
}
