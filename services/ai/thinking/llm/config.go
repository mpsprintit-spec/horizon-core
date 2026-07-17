// Package llm provides adapter for LLM interactions.
package llm

// LLMConfig holds configuration for LLM provider.
type LLMConfig struct {
	// Provider name e.g. "openai", "anthropic", etc.
	Provider string
	// Model name to use.
	Model string
	// APIKey for authentication.
	APIKey string
	// BaseURL for custom endpoints.
	BaseURL string
	// MaxTokens limit for generation.
	MaxTokens int
	// Temperature controls randomness.
	Temperature float64
	// TopP for nucleus sampling.
	TopP float64
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *LLMConfig {
	return &LLMConfig{
		Provider:    "openai",
		Model:       "gpt-4o",
		MaxTokens:   2048,
		Temperature: 0.7,
		TopP:        0.9,
	}
}
