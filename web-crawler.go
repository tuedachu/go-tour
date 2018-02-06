package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type cache struct {
	m   map[string]int
	mux sync.Mutex
}

//global variable
var c cache

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	if depth <= 0 {
		ch <- 1
		return
	}
	c.mux.Lock()
	if _, ok := c.m[url]; ok {
		fmt.Println("URL ", url, " already visited.")
		c.mux.Unlock()
		ch <- 1
		return
	} else {
		c.m[url] = 1
		c.mux.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			ch <- 1
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		subch := make(chan int, len(urls)-1)
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher, subch)
		}
		for i := 0; i < len(urls); i++ {
			<-subch
		}
	}
	ch <- 1
	return
}

func main() {
	c.m = make(map[string]int)
	ch := make(chan int, 0)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
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
