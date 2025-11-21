package checker

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	UpSites    int
	DownSites  int
	ErrorSites int
	mu         sync.Mutex
}

func worker(site string, result *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Head(site)
	if err != nil {
		result.mu.Lock()
		result.ErrorSites++
		result.mu.Unlock()
	}
	if resp.StatusCode == 200 {
		fmt.Println(site, "\t UP")
		result.mu.Lock()
		result.UpSites++
		result.mu.Unlock()
	} else {
		fmt.Println(site, "\t DOWN")
		result.mu.Lock()
		result.DownSites++
		result.mu.Unlock()
	}
}

func CheckSites(sites []string) *Result {
	result := &Result{}
	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)
		go worker(site, result, &wg)
	}
	wg.Wait()
	return result
}
