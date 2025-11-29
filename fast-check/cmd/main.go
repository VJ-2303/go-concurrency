package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/VJ-2303/fast-check/internal/reader"
	"github.com/VJ-2303/fast-check/internal/types"
	"github.com/VJ-2303/fast-check/internal/worker"
)

const usage string = `Usage
fast-check <filename>
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(1)
	}
	path := os.Args[1]

	URLs, err := reader.ReadFiles(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	result := make(chan types.Result, len(URLs))

	for _, url := range URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			worker.CheckSite(url, result)
		}(url)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for r := range result {
		if r.Error != nil {
			fmt.Printf("site : %s Error : %s\n", r.URL, r.Error)
		} else if r.StatusCode != http.StatusOK {
			fmt.Printf("site : %s Bad Status %d\n", r.URL, r.StatusCode)
		} else {
			fmt.Printf("site : %s Okay\n", r.URL)
		}
	}
}
