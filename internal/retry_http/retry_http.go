package utils

import (
	"context"
	"eppv2/internal/constants"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func RetryHTTP(client *http.Client, request *http.Request, retries int, retriesInterval time.Duration, logger *zap.SugaredLogger) (*http.Response, error) {
	var resp *http.Response
	var err error

	for retries > 0 {
		startTime := time.Now()
		ctx, _ := context.WithTimeout(context.Background(), constants.HTTP_RETRY_TIMEOUT)
		req := request.WithContext(ctx)
		resp, err = client.Do(req)
		if err != nil {
			// if err is not nil, return the resp and error as is.
			return resp, err
		}

		shouldRetry := checkRetry(resp.StatusCode)
		if !shouldRetry {
			return resp, err
		} else {
			retries -= 1

			resp.Body.Close()

			sleepTime := time.Until(startTime.Add(retriesInterval))
			logger.Infow("Request Failed",
				"request url", request.URL,
				"retries left", retries,
				"retrying in", sleepTime)

			time.Sleep(sleepTime)
		}
	}

	return resp, err
}

func checkRetry(statusCode int) bool {
	// 408 Request Timeout
	// 425 Too Early
	// 429 Too Many Requests
	// 500 Internal Server Error
	// 502 Bad Gateway
	// 503 Service Unavailable
	// 504 Gateway Timeout

	validStatusCodeForRetry := []int{http.StatusRequestTimeout, http.StatusTooEarly, http.StatusTooManyRequests, http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout}
	for _, value := range validStatusCodeForRetry {
		if value == statusCode {
			return true
		}
	}

	return false
}
