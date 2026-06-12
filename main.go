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

	filePath := "urls.txt"

	fileReadStartTime := time.Now()

	memUsageBeforeLoading := getMemoryUsage()

	fileData, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error in file opening")
		return
	}

	memUsageAfterLoading := getMemoryUsage()

	urls := []string{}
	url := ""

	for _, char := range fileData {
		if char == '\n' {
			if url != "" {
				urls = append(urls, url)
			}
			url = ""
		} else {
			url += string(char)
		}
	}

	totalFileReadTime := time.Since(fileReadStartTime)

	fmt.Println("Total No.of lines: ", len(urls))
	fmt.Println("Total time to read the file: ", totalFileReadTime)
	fmt.Println("Memory Usage: ", (memUsageAfterLoading - memUsageBeforeLoading))
}
