package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	duration := 1000 * time.Millisecond
	ctx := context.Background()
	d := time.Now().Add(duration)
	ctx = context.WithValue(ctx, "test", "hello")
	ctx, cancel := context.WithDeadline(ctx, d)
	

	defer cancel()

	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response complited, status code=%d\n", res.StatusCode)
		fmt.Println(ctx.Value("test"))
	case <-ctx.Done():
		fmt.Println("request too long")
	}
}
