package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
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

	filePath := "./urls_big.txt"

	file, err := os.Open(filePath)
	fileData := bufio.NewReader(file)

	if err != nil {
		fmt.Println("Error in file opening")
		return
	}

	urls := []string{}

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

	results := []checker.HealthResult{}
	healthCheckerStartTime := time.Now()

	for _, url := range urls {
		result := checker.CheckUrlHealth(url, client)
		results = append(results, result)
	}

	healthCheckerEndTime := time.Since(healthCheckerStartTime)

	fmt.Printf("Total time to check health for %d urls: %v\n", len(urls), healthCheckerEndTime)
}
