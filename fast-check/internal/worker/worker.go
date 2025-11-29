package worker

import (
	"net/http"
	"time"

	"github.com/VJ-2303/fast-check/internal/types"
)

func CheckSite(URL string, result chan<- types.Result) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(URL)
	if err != nil {
		r := types.Result{
			URL:        URL,
			StatusCode: 404,
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
