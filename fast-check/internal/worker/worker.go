package worker

import (
	"net/http"
	"time"

	"github.com/VJ-2303/fast-check/internal/types"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func CheckSite(URL string, result chan<- types.Result) {
	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		r := types.Result{
			URL:   URL,
			Error: err,
		}
		result <- r
		return
	}
	defer resp.Body.Close()

	r := types.Result{
		URL:        URL,
		StatusCode: resp.StatusCode,
	}

	result <- r
}
