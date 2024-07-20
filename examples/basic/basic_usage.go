package main

import (
	"context"
	"fmt"
	"os"

	unkey "github.com/glamboyosa/unkeygoratelimit"
)

func main() {
	userId := "user_4f2G8T1zS4XkNoVfX8bW5E57P3U9"
	rateLimiter := unkey.New(os.Getenv("ROOT_KEY"), unkey.UnkeyRateLimiterNew{
		Namespace: "osa.test",
		Limit:     100,
		Duration:  120000,
	})

	result, err := rateLimiter.Ratelimit(context.Background(), userId, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
