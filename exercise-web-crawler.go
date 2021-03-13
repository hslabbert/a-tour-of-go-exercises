package main

import (
	"fmt"
	"sync"
)

// SafeUrlCache is safe to use concurrently.
type SafeUrlCache struct {
	mu sync.Mutex
	v  map[string]bool
	wg sync.WaitGroup
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (uc *SafeUrlCache) Crawl(url string, depth int, fetcher Fetcher) {
	defer uc.wg.Done()
	if depth <= 0 {
		return
	}
	uc.mu.Lock()
	body, urls, err := fetcher.Fetch(url)
	uc.v[url] = true
	uc.mu.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if _, ok := uc.v[u]; !ok {
			uc.wg.Add(1)
			go uc.Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	uc := SafeUrlCache{v: make(map[string]bool)}
	uc.wg.Add(1)
	uc.Crawl("https://golang.org/", 4, fetcher)
	uc.wg.Wait()
	fmt.Println("Final url cache:", uc.v)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
