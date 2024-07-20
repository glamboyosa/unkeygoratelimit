package main

import (
	"context"
	"fmt"

	"github.com/glamboyosa/unkeygoratelimit"
	"github.com/glamboyosa/unkeygoratelimit/providers"
)

func main() {
	rateLimiter := unkeygoratelimit.New("your-root-key", unkeygoratelimit.UnkeyRateLimiterNew{
		Namespace: "example",
		Limit:     100,
		Duration:  60,
	})

	opts := &providers.UnkeyRateLimiterOptions{
		Cost:      5,
		Async:     true,
		Meta:      providers.UnkeyMeta{}, // can be left off
		Resources: []providers.UnkeyResource{}, // can be left off
	}

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
