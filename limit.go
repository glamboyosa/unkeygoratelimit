package unkeygoratelimit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/glamboyosa/unkeygoratelimit/providers"
)

type UnkeyRateLimiterNew struct {
	Namespace string
	Limit     int
	Duration  int
	Timeout   *UnkeyRateLimiterTimeout
}

type unkeyRateLimiterNewInit struct {
	Namespace string
	Limit     int
	Duration  int
	Timeout   *UnkeyRateLimiterTimeout
	rootKey   string
}

type UnkeyRateLimiterTimeout struct {
	Ms       int
	Fallback providers.RateLimitResult
}

func New(rootKey string, i UnkeyRateLimiterNew) unkeyRateLimiterNewInit {
	var timeout *UnkeyRateLimiterTimeout
	if i.Timeout != nil {
		timeout = i.Timeout
	}

	return unkeyRateLimiterNewInit{
		Namespace: i.Namespace,
		Limit:     i.Limit,
		Duration:  i.Duration,
		rootKey:   rootKey,
		Timeout:   timeout,
	}
}
func (r *unkeyRateLimiterNewInit) Ratelimit(ctx context.Context, identifier string, opts *providers.UnkeyRateLimiterOptions) (providers.RateLimitResult, error) {
	url := "https://api.unkey.dev/v1/ratelimits.limit"
	payload := mergePayload(r, opts, identifier)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return providers.RateLimitResult{}, fmt.Errorf("error marshalling payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		// Return the fallback value if Timeout is not nil
		if r.Timeout != nil {
			return r.Timeout.Fallback, fmt.Errorf("error creating request: %v. Using Fallback: %v", err, r.Timeout.Fallback)
		}
		return providers.RateLimitResult{}, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+r.rootKey)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	if r.Timeout != nil {
		client.Timeout = time.Duration(r.Timeout.Ms) * time.Millisecond
	}
	res, err := client.Do(req)
	if err != nil {
		// Return the fallback value if Timeout is not nil
		if r.Timeout != nil {
			return r.Timeout.Fallback, fmt.Errorf("error making request: %v. Using Fallback: %v", err, r.Timeout.Fallback)
		}
		return providers.RateLimitResult{}, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		// Return the fallback value if Timeout is not nil
		if r.Timeout != nil {
			return r.Timeout.Fallback, fmt.Errorf("error reading response body: %v. Using Fallback: %v", err, r.Timeout.Fallback)
		}
		return providers.RateLimitResult{}, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		if r.Timeout != nil {
			return r.Timeout.Fallback, fmt.Errorf("unexpected status code: %v, body: %v using Fallback %v", res.StatusCode, string(body), r.Timeout.Fallback)
		}

		return providers.RateLimitResult{}, fmt.Errorf("unexpected status code: %v, body: %v using Fallback %v", res.StatusCode, string(body), nil)
	}

	var apiResponse providers.RateLimitResult
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return providers.RateLimitResult{}, err
	}

	return apiResponse, nil
}

func mergePayload(r *unkeyRateLimiterNewInit, opts *providers.UnkeyRateLimiterOptions, id string) providers.UnkeyRateLimiterPayload {

	if opts == nil {

		payload := providers.UnkeyRateLimiterPayload{
			Namespace:  r.Namespace,
			Identifier: id,
			Limit:      r.Limit,
			Duration:   r.Duration,
		}
		return payload
	} else {

		payload := providers.UnkeyRateLimiterPayload{
			Namespace:  r.Namespace,
			Identifier: id,
			Limit:      r.Limit,
			Duration:   r.Duration,
			Cost:       opts.Cost,
			Async:      opts.Async,
			Meta:       opts.Meta,
			Resources:  opts.Resources,
		}
		return payload
	}
}
