// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopl.io/ch5_test/links"
)

var (
	dirPrefix string = "goproxy.cn"
	dir       string
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	var file *os.File
	defer file.Close()
	if file, err = os.OpenFile(dir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0); err != nil {
		fmt.Printf("open file [%v] error: %v", dir, err)
		os.Exit(1)
	}
	for _, link := range list {
		if strings.HasPrefix(link, "https://"+dirPrefix) {
			if _, err = file.Write([]byte(link + "\n")); err != nil {
				fmt.Printf("write file [%v] error: %v", dir, err)
				os.Exit(1)
			}
		}
	}
	return list
}

//!-crawl

//!+main
func main() {
	dir, _ = os.Getwd()
	dir += "\\" + dirPrefix
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
