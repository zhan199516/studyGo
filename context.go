package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "working 1")
	go watch(ctx, "working 2")
	go watch(ctx, "working3")

	time.Sleep(10 * time.Second)
	fmt.Println("OK!, working stop")
	cancel()
	time.Sleep(5 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "name out,stop....")
			return
		default:
			fmt.Println(name, "goroutine working.....")
			time.Sleep(2 * time.Second)
		}
	}
}
