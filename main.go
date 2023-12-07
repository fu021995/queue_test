package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

const MaxWorker = 2

type Request struct {
	Samples int    `json:"samples"`
	Payload string `json:"payload"`
}

func processRequest(request Request) {
	// 在这里编写处理逻辑，根据需要处理请求的数据
	for i := 0; i < request.Samples; i++ {
		go func() {
			// 在这里编写处理逻辑，根据需要处理请求的数据
			fmt.Printf("Processing request with %d samples: %s\n", request.Samples, request.Payload)
			time.Sleep(3 * time.Second)
		}()
	}
	time.Sleep(3 * time.Second)
}

func worker(wg *sync.WaitGroup, requests <-chan Request) {
	defer wg.Done()
	for request := range requests {
		processRequest(request)
	}
}

func main() {
	requests := make(chan Request)
	var wg sync.WaitGroup

	// 启动多个 Goroutine 来处理请求
	for i := 0; i < MaxWorker; i++ {
		wg.Add(1)
		go worker(&wg, requests)
	}

	// 模拟从用户接收到的请求
	jsonData := `[{"samples": 1, "payload": "Request 1"}, {"samples": 2, "payload": "Request 2"}, {"samples": 3, "payload": "Request 3"},{"samples": 4, "payload": "Request 3"},{"samples": 5, "payload": "Request 2"}]`
	var incomingRequests []Request
	err := json.Unmarshal([]byte(jsonData), &incomingRequests)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	// 将请求发送到通道中进行排队处理
	for _, request := range incomingRequests {
		requests <- request
	}
	close(requests)

	wg.Wait()

	fmt.Println("All requests processed.")

}

