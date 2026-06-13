package checker

import (
	"io"
	"net/http"
	"time"
)

type HealthResult struct {
	Url          string
	StatusCode   int
	ResponseTime time.Duration
}

func CheckUrlHealth(url string, client http.Client) HealthResult {

	startTime := time.Now()

	resp, err := client.Get(url)

	if err != nil {
		return HealthResult{
			Url:          url,
			StatusCode:   0,
			ResponseTime: time.Since(startTime),
		}
	}

	defer resp.Body.Close()

	io.Copy(io.Discard, resp.Body)

	responseTime := time.Since(startTime)

	return HealthResult{
		Url:          url,
		StatusCode:   resp.StatusCode,
		ResponseTime: responseTime,
	}
}
