package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println(val)

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	responseChan := make(chan Response)

	go func() {
		val, err := fetchSlowData()
		responseChan <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data took time to long")
		case response := <-responseChan:
			return response.value, response.err
		}
	}
}

func fetchSlowData() (int, error) {
	time.Sleep(time.Millisecond * 120)

	return 666, nil
}
