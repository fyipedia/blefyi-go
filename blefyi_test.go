package blefyi_test

import (
	"testing"

	blefyi "github.com/fyipedia/blefyi-go"
)

func TestNewClient(t *testing.T) {
	c := blefyi.NewClient()
	if c.BaseURL != blefyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", blefyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := blefyi.NewClient()
	result, err := c.Search("heart rate")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
