package main

import (
	"context"
	"fmt"
)

type tracer struct {
	trace string
}

func childFunction(ctx context.Context) {
	fmt.Println(ctx.Value(tracer{"main"}))
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, tracer{"main"}, "context value result")

	childFunction(ctx)
}
