package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://example.com",
		"http://httpbin.org/delay/5",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		go func(url string, ctx context.Context, cancel context.CancelFunc) {
			defer wg.Done()
			defer cancel()
			fetchURL(ctx, url)
		}(url, ctx, cancel)
	}
	wg.Wait()
}

func fetchURL(ctx context.Context, url string) {
	c := make(chan string)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			close(c) 
			return
		}
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			close(c) 
			return
		}
		fmt.Println("Fetched HTML:")
		c <- "ok"
		close(c)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
		return
	case v := <-c:
		if v != "" {
			fmt.Println(v)
		}
	}
}
