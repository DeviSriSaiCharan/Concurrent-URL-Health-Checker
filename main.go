package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/DeviSriSaiCharan/GoLang-Learnings/checker"
)

var client http.Client = http.Client{
	Timeout: 3 * time.Second,
}

func getMemoryUsage() float64 {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	memUsage := float64(m.Alloc) / 1024 / 1024 // Convert bytes to mega bytes

	return memUsage
}

func main() {

	filePath := "./urls.txt"
	workers := 150

	urls, err := getUrlsFromTextFile(filePath)

	if err != nil {
		return
	}

	var wg sync.WaitGroup

	results := make(chan checker.HealthResult, len(urls))
	jobs := make(chan string, len(urls))

	healthCheckerStartTime := time.Now()

	for i := range workers {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	healthCheckerEndTime := time.Since(healthCheckerStartTime)

	close(results)

	fmt.Printf("Total time to check health for %d urls: %v\n", len(results), healthCheckerEndTime)
}

func getUrlsFromTextFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	fileData := bufio.NewReader(file)

	urls := []string{}

	if err != nil {
		fmt.Println("Error in file opening")
		return urls, err
	}

	for {
		line, _, err := fileData.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file: ", err)
		} else {
			if len(line) != 0 {
				urls = append(urls, string(line))
			}
		}
	}

	return urls, nil
}

func worker(id int, jobs <-chan string, results chan<- checker.HealthResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		result := checker.CheckUrlHealth(url, client)
		results <- result
	}

}
