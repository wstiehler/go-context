package main

import (
	"context"
	"fmt"
	"time"
)

type tracer struct {
	trace string
}

func main() {
	forever := make(chan struct{})

	ctx := context.Background()

	ctx = context.WithValue(ctx, tracer{"main"}, "context value result")

	ctx, _ = context.WithTimeout(ctx, 20000*time.Millisecond)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				switch ctx.Err() {
				case context.DeadlineExceeded:
					fmt.Println("context timeout exceeded")
				case context.Canceled:
					fmt.Println("context cancelled by force. Whole process is complete")
				}
				forever <- struct{}{}
				return
			default:
				fmt.Println("tracer: ", ctx.Value(tracer{"main"}))
			}
			time.Sleep(500 * time.Microsecond)
		}
	}(ctx)
	<-forever
	fmt.Println("finish")
}
