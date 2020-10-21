package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	readFlags()
	query := readQuery()
	if query == "" {
		return
	}
	results, err := search(query)
	if err != nil {
		handle(err)
		return
	}
	results = limit(results)
	reverse(results)
	display(results)
}

func readFlags() {
	fResultLimit.Read()
	flag.Parse()
}

func readQuery() string {
	args := flag.Args()
	if len(args) < 1 {
		return ""
	}
	return strings.Join(args, querySeparator)
}

func search(q string) ([]Result, error) {
	q = url.QueryEscape(q)
	crawlingURL := fmt.Sprintf(crawURL, q)
	return crawl(crawlingURL)
}

func crawl(url string) ([]Result, error) {
	var results []Result
	c := colly.NewCollector()
	c.OnHTML("div.SearchSnippet", func(e *colly.HTMLElement) {
		path := e.ChildText(".SearchSnippet-header > a")
		packageURL := e.ChildAttr(".SearchSnippet-header > a", "href")
		desc := e.ChildText(".SearchSnippet-synopsis")
		result := Result{
			Path:        path,
			URL:         buildURL(packageURL),
			Description: desc,
		}
		results = append(results, result)
	})
	err := c.Visit(url)
	return results, err
}

func buildURL(packageURL string) string {
	for _, repo := range repos {
		if strings.Contains(packageURL, repo) {
			return scheme + packageURL[1:]
		}
	}
	return ""
}

func display(results []Result) {
	fmt.Printf("\n")

	for i, result := range results {
		if i != 0 {
			fmt.Printf("\n----------\n\n")
		}
		fmt.Println(result.Path)
		if result.URL != "" {
			fmt.Println(result.URL)
		}
		if result.Description != "" {
			fmt.Println(result.Description)
		}
	}

	fmt.Printf("\n")
}

func handle(err error) {
	log.Printf("Failed to search due to: %v\n", err)
}

func reverse(results []Result) {
	l := len(results)
	m := l / 2
	for i := 0; i < m; i++ {
		results[i], results[l-i-1] = results[l-i-1], results[i]
	}
}

func limit(results []Result) []Result {
	return results[:fResultLimit.Value()]
}
