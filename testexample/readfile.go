package readfile

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile opens a file and returns its contents as a []string
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't load file %s: %s\n", filePath, err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
