package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Error      error
}

func checkURL(url string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	start := time.Now()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)

	duration := time.Since(start)

	if err != nil {
		ch <- Result{
			URL:      url,
			Duration: duration,
			Error:    err,
		}
		return
	}

	defer resp.Body.Close()

	ch <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   duration,
	}
}

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var urls []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())

		if url == "" || strings.HasPrefix(url, "#") {
			continue
		}

		urls = append(urls, url)
	}

	return urls, scanner.Err()
}

func main() {
	var urls []string

	if len(os.Args) > 1 {
		urls = os.Args[1:]
	} else {
		var err error

		urls, err = readURLs("urls.txt")

		if err != nil {
			fmt.Println("Usage:")
			fmt.Println("go run main.go google.com github.com")
			fmt.Println("")
			fmt.Println("Or add websites to urls.txt")
			return
		}
	}

	if len(urls) == 0 {
		fmt.Println("No URLs provided.")
		return
	}

	fmt.Println("GoPulse")
	fmt.Println("------------------------------")

	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	// Limit simultaneous requests to avoid exhausting resources on large URL lists.
	const maxConcurrentChecks = 10
	slots := make(chan struct{}, maxConcurrentChecks)

	for _, url := range urls {
		slots <- struct{}{}
		wg.Add(1)

		go func(url string) {
			defer func() { <-slots }()
			checkURL(url, results, &wg)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Printf(
				"[DOWN] %-30s %v\n",
				result.URL,
				result.Error,
			)

			continue
		}

		status := "ONLINE"

		if result.StatusCode >= 400 {
			status = "ERROR"
		}

		fmt.Printf(
			"[%d] %-30s %-7s %v\n",
			result.StatusCode,
			result.URL,
			status,
			result.Duration.Round(time.Millisecond),
		)
	}
}
