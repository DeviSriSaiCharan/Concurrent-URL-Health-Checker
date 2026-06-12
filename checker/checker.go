package checker

import (
	"net/http"
)

type HealthResult struct {
	Url        string
	StatusCode int
	Err        error
}

func CheckUrlHealth(url string, client http.Client) HealthResult {

	resp, err := client.Get(url)

	if err != nil {
		return HealthResult{
			Url:        url,
			StatusCode: 0,
			Err:        err,
		}
	}

	defer resp.Body.Close()

	return HealthResult{
		Url:        url,
		StatusCode: resp.StatusCode,
		Err:        nil,
	}
}
