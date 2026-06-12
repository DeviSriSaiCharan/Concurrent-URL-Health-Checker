package main

import (
	"bufio"
	"fmt"
	"io"
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

	file, err := os.Open(filePath)
	fileData := bufio.NewReader(file)

	if err != nil {
		fmt.Println("Error in file opening")
		return
	}
	memUsageAfterLoading := getMemoryUsage()

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

	totalMemUsage := getMemoryUsage()
	totalFileReadTime := time.Since(fileReadStartTime)

	fmt.Println("Total No.of lines: ", len(urls))
	fmt.Println("Total time to read the file: ", totalFileReadTime)
	fmt.Println("Memory Usage: ", (memUsageAfterLoading - memUsageBeforeLoading))
	fmt.Println("Total memory usage: ", (totalMemUsage - memUsageBeforeLoading))
}
