package main

import (
	"context"
	"fmt"

	unkey "github.com/glamboyosa/unkeygoratelimit"
)

func main() {
	rateLimiter := unkey.New("your-root-key", unkey.UnkeyRateLimiterNew{
		Namespace: "example",
		Limit:     100,
		Duration:  60,
	})

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
