package reader_test

import (
	"testing"

	"github.com/VJ-2303/sitecheck/reader"
	"github.com/google/go-cmp/cmp"
)

func TestReadLines(t *testing.T) {
	path := "./../testdata/site_test.txt"
	want := []string{
		"https://gemini.google.com",
		"https://www.google.com",
		"https://youtube.com",
	}
	got, err := reader.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
