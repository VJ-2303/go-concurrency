package worker_test

import (
	"net/http"
	"testing"

	"github.com/VJ-2303/fast-check/internal/types"
	"github.com/VJ-2303/fast-check/internal/worker"
)

func TestCheckSite_ReturnsStatusOk(t *testing.T) {
	result := make(chan types.Result)
	url := "https://google.com"
	go worker.CheckSite(url, result)

	r := <-result
	if r.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d. got %d instead", http.StatusOK, r.StatusCode)
	}
}
