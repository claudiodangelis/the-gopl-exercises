// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, errCopy := io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close()
		if errCopy != nil {
			fmt.Fprint(os.Stderr, "Error while fetching the url")
			os.Exit(1)
		}
	}
}

//!-
