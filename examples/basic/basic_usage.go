package main

import (
	"context"
	"fmt"
	"os"

	unkey "github.com/glamboyosa/unkeygoratelimit"
)

func main() {
	rateLimiter := unkey.New(os.Getenv("ROOT_KEY"), unkey.UnkeyRateLimiterNew{
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
