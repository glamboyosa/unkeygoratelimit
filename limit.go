package unkeygoratelimit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/glamboyosa/unkeygoratelimit/providers"
)

type UnkeyRateLimiterNew struct {
	Namespace string
	Limit     int
	Duration  int
}
type unkeyRateLimiterNewInit struct {
	Namespace string
	Limit     int
	Duration  int
	rootKey   string
}

func New(rootKey string, i UnkeyRateLimiterNew) unkeyRateLimiterNewInit {
	return unkeyRateLimiterNewInit{
		Namespace: i.Namespace,
		Limit:     i.Limit,
		Duration:  i.Duration,
		rootKey:   rootKey,
	}
}

func (r *unkeyRateLimiterNewInit) Ratelimit(ctx context.Context, identifier string, opts *providers.UnkeyRateLimiterOptions) interface{} {
	url := "https://api.unkey.dev/v1/ratelimits.limit"

	payload := mergePayload(r, opts, identifier)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+r.rootKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v, body: %v", res.StatusCode, string(body))
	}

	fmt.Println("Response:", string(body))
	return nil
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
