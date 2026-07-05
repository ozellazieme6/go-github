package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

const maxRetries = 3

// handleSecondaryRateLimit implements the retry logic for 429/403 responses.
func handleSecondaryRateLimit(resp *http.Response, attempt int) (time.Duration, bool) {
	if resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode != http.StatusForbidden {
		return 0, false
	}

	if attempt >= maxRetries {
		return 0, false
	}

	retryAfter := resp.Header.Get("Retry-After")
	if retryAfter == "" {
		// Exponential backoff: 1s, 2s, 4s
		backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
		return backoff, true
	}

	// Logic for parsing Retry-After header would go here
	return 0, true
}

func main() {
	fmt.Println("Secondary rate limit retry logic initialized.")
}