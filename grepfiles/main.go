package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func grepFile(key, filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("\nError Opening file:%s", filename)
		return
	}
	reader := bufio.NewScanner(file)
	lines := 0
	for reader.Scan() {
		lines++
		line := reader.Text()
		if strings.Contains(line, key) {
			fmt.Printf("\nFound at file:%s at Lineno:%d\n", filename, lines)
			fmt.Printf(line + "\n")
		}
	}
}

func main() {
	start := time.Now()
	key := os.Args[1]
	files := os.Args[2:]

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go grepFile(key, file, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Took:", elapsed)
}
