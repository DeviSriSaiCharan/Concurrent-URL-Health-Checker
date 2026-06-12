package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func getMemoryUsage() float64 {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	memUsage := float64(m.Alloc) / 1024 / 1024 // Convert bytes to mega bytes

	return memUsage
}

func main() {

	filePath := "./urls_big.txt"

	fileReadStartTime := time.Now()

	memUsageBeforeLoading := getMemoryUsage()

	fileData, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error in file opening")
		return
	}

	memUsageAfterLoading := getMemoryUsage()

	urls := []string{}
	url := []byte{}

	for _, char := range fileData {
		if char == '\n' {
			if len(url) != 0 {
				urls = append(urls, string(url))
				url = []byte{}
			}
		} else {
			url = append(url, char)
		}
	}

	totalFileReadTime := time.Since(fileReadStartTime)

	fmt.Println("Total No.of lines: ", len(urls))
	fmt.Println("Total time to read the file: ", totalFileReadTime)
	fmt.Println("Memory Usage: ", (memUsageAfterLoading - memUsageBeforeLoading))
}
