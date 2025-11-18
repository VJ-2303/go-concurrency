package main

import (
	"fmt"
	"os"
	"sync"
)

func readFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("\nError opening file %s", filename)
	} else {
		fmt.Printf("\nContents of file %s\n%v", filename, string(data))
	}
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Usage go run main.go <file1> <file2> ....")
		os.Exit(1)
	}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go readFile(file, &wg)
	}
	wg.Wait()
}
