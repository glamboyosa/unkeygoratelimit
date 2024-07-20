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
		Timeout: &unkeygoratelimit.UnkeyRateLimiterTimeout{
			Ms: 5000,
			Fallback: providers.RateLimitResult{
				Success:   true,
				Limit:     100,
				Reset:     1630000000000,
				Remaining: 50,
			},
		},
	})

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)	
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
