package llm

import (
	"context"
)

// LLMProvider defines the interface for different LLM backends.
type LLMProvider interface {
	// GenerateText sends prompt to provider and returns response.
	GenerateText(ctx context.Context, prompt string, config *LLMConfig) (string, error)
}

// LLMManager manages LLM interactions and provider selection.
type LLMManager interface {
	// Generate handles the full LLM request lifecycle.
	Generate(ctx context.Context, prompt string, opts ...func(*LLMConfig)) (*LLMResponse, error)
}
