package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "google.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"bing.com",
		"fun.com",
	}

	want := map[string]bool{
		"google.com": false,
		"bing.com":   true,
		"fun.com":    true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
