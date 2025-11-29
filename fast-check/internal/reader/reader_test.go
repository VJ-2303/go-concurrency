package reader_test

import (
	"errors"
	"os"
	"testing"

	"github.com/VJ-2303/fast-check/internal/reader"
	"github.com/google/go-cmp/cmp"
)

func TestReadFiles_ReturnsErrorWhenFileNoFound(t *testing.T) {
	path := "testdata/notfound.txt"

	_, err := reader.ReadFiles(path)
	println(err.Error())
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected not file not exists error")
	}
}

func TestReadFiles_ReturnsNoError(t *testing.T) {
	path := "testdata/sites.txt"

	want := []string{
		"https://google.com",
		"https://youtube.com",
	}
	got, err := reader.ReadFiles(path)
	if err != nil {
		t.Fatalf("expected no error, but got %q", err)
	}
	if !cmp.Equal(got, want) {
		t.Fatal(cmp.Diff(got, want))
	}
}
