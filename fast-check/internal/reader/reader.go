package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFiles(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var URLs []string

	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			if strings.HasPrefix(url, "https://") {
				URLs = append(URLs, url)
			} else {
				url = "https://" + url
				URLs = append(URLs, url)
			}
		}
	}
	return URLs, scanner.Err()
}
