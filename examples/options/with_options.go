package main

import (
	"context"
	"fmt"
	"os"

	"github.com/glamboyosa/unkeygoratelimit"
	"github.com/glamboyosa/unkeygoratelimit/providers"
)

func main() {
	userId := "user_4f2G8T1zS4XkNoVfX8bW5E57P3U9"
	rateLimiter := unkeygoratelimit.New(os.Getenv("ROOT_KEY"), unkeygoratelimit.UnkeyRateLimiterNew{
		Namespace: "osa.test",
		Limit:     100,
		Duration:  120000,
	})

	opts := &providers.UnkeyRateLimiterOptions{
		Cost:  5,
		Async: false,
	}

	result, err := rateLimiter.Ratelimit(context.Background(), userId, opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
