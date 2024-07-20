# UnkeyGoRatelimit

> "( Globally consistent, fast ) - choose two"

Global ratelimiting built for the modern web, now for Go.

The `unkeygoratelimit` package provides functionality to interact with the Unkey Rate Limiter API. It allows you to check rate limits and handle rate limiting with customizable options and timeouts.

## Getting Started

### Setup

1. **Obtain Your Root Key**:

   - Visit app.unkey.com to get your root key. Follow the setup instructions provided on the platform.

   - Set up namespaces as needed for your application.

2. **Set Environment Variables**:

   - Ensure you set the `ROOT_KEY` environment variable with your obtained root key.

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/glamboyosa/unkeygoratelimit"
)

func main() {
	rateLimiter := unkeygoratelimit.New("your-root-key", unkeygoratelimit.UnkeyRateLimiterNew{
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

```

### With Options

You can use additional options by passing an `UnkeyRateLimiterOptions` struct:

```go
package main

import (
	"context"
	"fmt"
	"github.com/yourusername/unkeygoratelimit"
	"github.com/yourusername/unkeygoratelimit/providers"
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
		Meta:      providers.UnkeyMeta{}, // Set your meta data here
		Resources: []providers.UnkeyResource{}, // Set your resources here
	}

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
```

### With Timeout

[Unkey](https://unkey.com) provides safe guards in case of severe network degredations or other unforseen events. AKA `timeouts`. When configuring a timeout, you will provide a `fallback` result that will be used if an error occurs:

```go
package main

import (
	"context"
	"fmt"
	"github.com/yourusername/unkeygoratelimit"
	"github.com/yourusername/unkeygoratelimit/providers"
)

func main() {
	rateLimiter := unkeygoratelimit.New("your-root-key", unkeygoratelimit.UnkeyRateLimiterNew{
		Namespace: "example",
		Limit:     100,
		Duration:  60,
		Timeout: &unkeygoratelimit.UnkeyRateLimiterTimeout{
			Ms: 5000,
			Fallback: providers.RateLimitResult{
				Success:   true,
				Limit:     0,
				Reset:     0,
				Remaining: 0,
			},
		},
	})

	result, err := rateLimiter.Ratelimit(context.Background(), "user_123", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Rate Limit Result: %+v\n", result)
}
```

> [!NOTE]  
> When using the rate limiter, ensure you handle errors properly. If an error occurs and a timeout fallback is set, the function will return the **fallback result**. For proper error handling, check both the result and the error returned.

## API and Types

`UnkeyRateLimiterNew`
Defines the configuration needed to initialize a new rate limiter instance.

```go
type UnkeyRateLimiterNew struct {
    Namespace string                    // Required. The namespace for your application.
    Limit     int                       // Required. The rate limit.
    Duration  int                       // Required. The duration of the rate limit in seconds.
    Timeout   *UnkeyRateLimiterTimeout // Optional. Timeout settings and fallback result.
}
```

`UnkeyRateLimiterTimeout`
Specifies the timeout configuration and the fallback result in case of errors.

```go
type UnkeyRateLimiterTimeout struct {
Ms int // Required. Timeout duration in milliseconds.
Fallback providers.RateLimitResult // Optional. Fallback result to use in case of an error.
}
```

`New(rootKey string, i UnkeyRateLimiterNew) unkeyRateLimiterNewInit`

Initializes a new rate limiter instance.

- Parameters:

      - rootKey: The root key for authorization.
      - i: Configuration object of type UnkeyRateLimiterNew.

  Returns:
  An instance of `unkeyRateLimiterNewInit`.

`Ratelimit(ctx context.Context, identifier string, opts *providers.UnkeyRateLimiterOptions) (providers.RateLimitResult, error)`
Makes a rate limit request to the Unkey API.

Parameters:

- `ctx`: Context for request management.
- `identifier`: The identifier for the rate limit request.
- `opts`: Optional. Additional options for the rate limit request.

Returns:

- `providers.RateLimitResult`: The result of the rate limit request.
- `error`: Any error encountered during the request.

`providers.RateLimitResult`
Contains the result of the rate limit request.

```go
type RateLimitResult struct {
    Success    bool // Indicates if the request was successful.
    Limit      int  // The limit applied.
    Reset      int64 // Time when the limit will reset (Unix timestamp).
    Remaining  int  // Number of remaining requests.
}

```

`providers.APIResponse`
Represents the full API response structure.

```go
type APIResponse struct {
    Result RateLimitResult // The result of the rate limit request.
}
```

## Contributing

If you have suggestions or improvements, please open an issue or a pull request on GitHub. Contributions are welcome!
