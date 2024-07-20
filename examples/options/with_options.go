package main

import (
	"context"
	"fmt"
	"os"

	"github.com/glamboyosa/unkeygoratelimit"
	"github.com/glamboyosa/unkeygoratelimit/providers"
)

func main() {
	rateLimiter := unkeygoratelimit.New(os.Getenv("ROOT_KEY"), unkeygoratelimit.UnkeyRateLimiterNew{
		Namespace: "example",
		Limit:     100,
		Duration:  60,
	})

	opts := &providers.UnkeyRateLimiterOptions{
		Cost:      5,
		Async:     true,
	}

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
