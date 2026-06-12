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

![File Reading Benchmarks](images/Screenshot%202026-06-12%20at%2010.33.18 AM.png)
