package main

import (
	"context"
	"fmt"
	"os"

	"github.com/glamboyosa/unkeygoratelimit"
	"github.com/glamboyosa/unkeygoratelimit/providers"
)

func main() {
	rateLimiter := unkeygoratelimit.New(os.Getenv("TEST_ROOT_KEY"), unkeygoratelimit.UnkeyRateLimiterNew{
		Namespace: "osa.test",
		Limit:     100,
		Duration:  120000,
		Timeout: &unkeygoratelimit.UnkeyRateLimiterTimeout{
			Ms: 10,
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
