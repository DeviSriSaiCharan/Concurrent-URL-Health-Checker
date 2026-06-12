# Concurrent URL Health Checker

A lightweight and efficient URL Health Checker built with Go that performs concurrent health checks on multiple websites and APIs. The application leverages Go's goroutines and channels to send parallel HTTP requests, significantly reducing the time required to monitor large sets of URLs.

The tool validates URL availability by checking HTTP response status codes, response times, and connectivity, providing a quick overview of service health. Its concurrent architecture ensures scalability and optimal resource utilization, making it suitable for monitoring web services, APIs, and infrastructure endpoints.

## Phase 1: Reading Data From File
I am using `os.ReadFile()` to read the entire file and parsed the content byte by byte.
- It takes the entire file and stores it in the memory

### Benchmarks
| Metric | Value |
|---------|---------|
| Number of URLs | 115 |
| Average time | 814 microseconds |
| Memory Usage | 0.0028762 MB ~ 2Kb |

| Metric | Value |
|---------|---------|
| Number of URLs | 100000 |
| Average time | 100 milliseconds |
| Memory Usage | 2.4768 MB ~ 2Mb |

![File Reading Benchmarks](images/Screenshot%202026-06-12%20at%2010.33.18 AM.png)


## Phase 2: Changed the URL from string to byte type
In Go `string` are immutable, it means its bytes cannot be modified
 - Before i am using `string` type for the `url` variables, so when ever i append a new character to it, go create a new string and copy all the values into the new string and discards the old string
- `slices` are mutable so when we append data to `byte[]` array it dont create a new byte[] everytime until capacity runs out.

### Benchmarks
| Metric | Value |
|---------|---------|
| Number of URLs | 100000 |
| Average time | 30 milliseconds |
| Memory Usage | 2.4768 MB ~ 2Mb |

## PHASE 3: Optimized File Reading
Before we are loading the entire file data into memory so that's the reason the memory usage for reading teh file is `2Mb` which the size fo the file.
- I have changed `os.ReadFile()` to `os.Open()` it gives an pointer to that file instead of loading the entire data, using that pointer we read the data line by line.

### Benchmarks
| Metric | Value |
|---------|---------|
| Number of URLs | 100000 |
| Average time | 12.55 milliseconds |
| Memory Usage | 0.00402 ~ 4Kb |