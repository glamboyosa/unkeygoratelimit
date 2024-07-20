package providers

type UnkeyRateLimiterPayload struct {
	Namespace  string          `json:"namespace"`
	Identifier string          `json:"identifier"`
	Limit      int             `json:"limit"`
	Duration   int             `json:"duration"`
	Cost       int             `json:"cost"`
	Async      bool            `json:"async"`
	Meta       UnkeyMeta       `json:"meta"`
	Resources  []UnkeyResource `json:"resources"`
}
type UnkeyRateLimiterOptions struct {
	Cost      int             `json:"cost"`
	Async     bool            `json:"async"`
	Meta      UnkeyMeta       `json:"meta"`
	Resources []UnkeyResource `json:"resources"`
}

type UnkeyResource struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UnkeyMeta map[string]interface{}

type RateLimitResult struct {
    Success   bool `json:"success"`
    Limit     int  `json:"limit"`
    Reset     int64 `json:"reset"`
    Remaining int  `json:"remaining"`
}

type APIResponse struct {
    Result RateLimitResult `json:"result"`
}
