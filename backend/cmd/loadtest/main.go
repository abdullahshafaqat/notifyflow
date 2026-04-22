package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type NotificationPayload struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	baseURL := "http://localhost:8080"
	numRequests := 1000000
	concurrency := 50 // 50 concurrent requests

	fmt.Printf("🚀 Starting load test: %d requests with %d concurrent workers\n\n", numRequests, concurrency)

	var (
		successCount int64
		failureCount int64
		totalTime    time.Duration
	)

	startTime := time.Now()

	// Semaphore to limit concurrent requests
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // Acquire
			defer func() { <-sem }() // Release

			payload := NotificationPayload{
				To:      fmt.Sprintf("user%d@example.com", id),
				Message: fmt.Sprintf("Test message %d at %s", id, time.Now().Format("15:04:05")),
			}

			body, _ := json.Marshal(payload)

			resp, err := http.Post(
				baseURL+"/send",
				"application/json",
				bytes.NewBuffer(body),
			)

			if err != nil {
				atomic.AddInt64(&failureCount, 1)
				fmt.Printf("❌ Request %d failed: %v\n", id, err)
				return
			}
			defer resp.Body.Close()

			respBody, _ := io.ReadAll(resp.Body)
			var respData Response
			json.Unmarshal(respBody, &respData)

			if resp.StatusCode == http.StatusOK && respData.Status == "success" {
				atomic.AddInt64(&successCount, 1)
				if id%10000 == 0 {
					fmt.Printf("✅ Request %d success\n", id)
				}
			} else {
				atomic.AddInt64(&failureCount, 1)
				fmt.Printf("❌ Request %d failed with status %d: %s\n", id, resp.StatusCode, respData.Message)
			}
		}(i)
	}

	wg.Wait()
	totalTime = time.Since(startTime)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("📊 LOAD TEST RESULTS\n")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Total Requests:  %d\n", numRequests)
	fmt.Printf("Success Count:   %d ✅\n", successCount)
	fmt.Printf("Failure Count:   %d ❌\n", failureCount)
	fmt.Printf("Success Rate:    %.2f%%\n", float64(successCount)/float64(numRequests)*100)
	fmt.Printf("Total Time:      %.2f seconds\n", totalTime.Seconds())
	fmt.Printf("Avg Time/Req:    %.2f ms\n", totalTime.Seconds()*1000/float64(numRequests))
	fmt.Printf("Throughput:      %.2f req/sec\n", float64(numRequests)/totalTime.Seconds())
	fmt.Println(strings.Repeat("=", 60))

	// Give workers time to finish processing before test ends
	fmt.Printf("\n⏳ Waiting 15 seconds for workers to finish processing...\n")
	time.Sleep(15 * time.Second)

	fmt.Printf("✅ Load test complete! Check MongoDB for all notifications.\n")
}
