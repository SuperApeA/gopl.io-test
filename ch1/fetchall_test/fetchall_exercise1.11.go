package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const timeout = 3 * time.Second

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	client := &http.Client{Timeout: timeout}

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		fmt.Printf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s\n", secs, nBytes, url)
}
