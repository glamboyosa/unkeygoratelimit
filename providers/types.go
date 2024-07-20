package providers

type UnkeyRateLimiterPayload struct {
	Namespace  string          `json:"namespace,omitempty"`
	Identifier string          `json:"identifier,omitempty"`
	Limit      int             `json:"limit,omitempty"`
	Duration   int             `json:"duration,omitempty"`
	Cost       int             `json:"cost,omitempty"`
	Async      bool            `json:"async,omitempty"`
	Meta       UnkeyMeta       `json:"meta,omitempty"`
	Resources  []UnkeyResource `json:"resources,omitempty"`
}
type UnkeyRateLimiterOptions struct {
	Cost      int             `json:"cost,omitempty"`
	Async     bool            `json:"async,omitempty"`
	Meta      UnkeyMeta       `json:"meta,omitempty"`
	Resources []UnkeyResource `json:"resources,omitempty"`
}

type UnkeyResource struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UnkeyMeta map[string]interface{}

type RateLimitResult struct {
	Success   bool  `json:"success"`
	Limit     int   `json:"limit"`
	Reset     int64 `json:"reset"`
	Remaining int   `json:"remaining"`
}

type APIResponse struct {
	Result RateLimitResult `json:"result"`
}
