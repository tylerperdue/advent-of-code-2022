package input

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("[os.Open]: %w", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines, nil
}
