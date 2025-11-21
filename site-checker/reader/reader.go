package reader

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrFileNotFound = errors.New("file not found")

func ReadFile(pathname string) ([]string, error) {
	file, err := os.OpenFile(pathname, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			if !strings.HasPrefix(line, "http") {
				line = "https://" + line
			}
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
