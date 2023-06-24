package Models

import (
	"bufio"
	"os"
)

func CountLines(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var count int
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}

	return count
}
