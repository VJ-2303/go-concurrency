package main

import (
	"fmt"
	"os"

	"github.com/VJ-2303/sitecheck/checker"
	"github.com/VJ-2303/sitecheck/reader"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Usage sitecheck <filename>")
		os.Exit(1)
	}
	sites, err := reader.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := checker.CheckSites(sites)
	fmt.Printf("UpSites: %d \t DownSites: \t %d ErrorSites: \t %d", result.UpSites, result.DownSites, result.ErrorSites)
}
