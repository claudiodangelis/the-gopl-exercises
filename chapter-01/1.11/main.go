// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func getTopSite(attrs []html.Attribute) (bool, string) {
	for _, attr := range attrs {
		if strings.HasPrefix(attr.Val, "/siteinfo/") && !strings.HasSuffix(attr.Val, "/") {
			r := strings.Split(attr.Val, "/")
			return true, "http://" + r[len(r)-1]
		}
	}
	return false, ""
}

func main() {
	start := time.Now()
	resp, err := http.Get("http://www.alexa.com/topsites")
	if err != nil {
		fmt.Printf("Error while fetching Top Sites: %v\n", err)
		return
	}

	document, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error while parsing body: %v\n", err)
		return
	}
	ch := make(chan string)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Do something with n...
			isTop, site := getTopSite(n.Attr)
			if isTop {
				fmt.Printf("Found one top site: %s\n", site)
				go fetch(site, ch)
				fmt.Println(<-ch)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(document)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	filename := strconv.FormatInt(time.Now().Unix(), 10) + ".html"
	outputFile, _ := os.Create(filename)
	nbytes, err := io.Copy(outputFile, resp.Body)
	outputFile.Close()
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
