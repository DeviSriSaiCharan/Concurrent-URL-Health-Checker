package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"

	"github.com/DeviSriSaiCharan/GoLang-Learnings/checker"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
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

	printResults(results)
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

func printResults(results <-chan checker.HealthResult) {
	i := 1
	data := [][]string{}

	for result := range results {
		status := strconv.Itoa(result.StatusCode)
		var coloredStatus string

		switch {
		case result.StatusCode >= 200 && result.StatusCode < 300:
			coloredStatus = color.New(color.FgGreen, color.Bold).Sprint(status)

		case result.StatusCode >= 300 && result.StatusCode < 400:
			coloredStatus = color.New(color.FgYellow, color.Bold).Sprint(status)

		default:
			coloredStatus = color.New(color.FgRed, color.Bold).Sprint(status)
		}

		outputResult := []string{
			strconv.Itoa(i),
			result.Url,
			coloredStatus,
			strconv.FormatInt(result.ResponseTime.Milliseconds(), 10) + " ms",
		}

		data = append(data, outputResult)
		i++
	}

	_, height, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println("Error getting terminal size")
		return
	}

	pageSize := height - 7

	if pageSize <= 0 {
		pageSize = 5
	}

	totalData := len(data)
	totalPages := (totalData + pageSize - 1) / pageSize // Ceil division to get total pages
	pageToRender := 1

	terminalInputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\r\033[H\033[2J")

		startIndex := (pageToRender - 1) * pageSize
		endIndex := startIndex + pageSize

		if endIndex > totalData {
			endIndex = totalData
		}

		pageData := data[startIndex:endIndex]

		table := tablewriter.NewTable(os.Stdout)
		table.Header(
			"#",
			"URL",
			"Status Code",
			"Response Time (ms)",
		)
		table.Bulk(pageData)
		table.Render()

		fmt.Printf("\nShowing %d-%d of %d.\n", startIndex+1, endIndex, totalData)

		if totalPages > 1 {
			if pageToRender == 1 {
				fmt.Print("[n] Next | [q] Quit: ")
			} else if pageToRender == totalPages {
				fmt.Print("[p] Previous | [q] Quit: ")
			} else {
				fmt.Print("[p] Previous | [n] Next | [q] Quit: ")
			}
		} else {
			fmt.Print("[q] Quit: ")
		}

		input, _ := terminalInputReader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "n":
			if pageToRender < totalPages {
				pageToRender++
			}
		case "p":
			if pageToRender > 1 {
				pageToRender--
			}
		case "q":
			return
		default:
			fmt.Println("Invalid input. Please enter 'n', 'p', or 'q'.")
		}
	}
}
