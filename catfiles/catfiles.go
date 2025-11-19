package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func readFile(filename string, wg *sync.WaitGroup, w io.Writer) {
	defer wg.Done()
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(w, "\nError opening file %s", filename)
	} else {
		fmt.Fprintf(w, "\nContents of file %s\n%v", filename, string(data))
	}
}

func ReadFiles(files []string, w io.Writer) {
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go readFile(file, &wg, w)
	}
	wg.Wait()
}
