package llm

import (
	"context"
	"fmt"
	"time"
)

// manager implements LLMManager.
type manager struct {
	provider LLMProvider
	config   *LLMConfig
}

// NewManager creates a new LLM manager.
func NewManager(provider LLMProvider, cfg *LLMConfig) LLMManager {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	return &manager{
		provider: provider,
		config:   cfg,
	}
}

// Generate orchestrates LLM generation.
func (m *manager) Generate(ctx context.Context, prompt string, opts ...func(*LLMConfig)) (*LLMResponse, error) {
	cfg := *m.config // copy
	for _, opt := range opts {
		opt(&cfg)
	}

	respText, err := m.provider.GenerateText(ctx, prompt, &cfg)
	if err != nil {
		return nil, fmt.Errorf("llm generation failed: %w", err)
	}

	return &LLMResponse{
		Content:      respText,
		Model:        cfg.Model,
		TokensUsed:   len(respText) / 4, // rough estimate
		FinishReason: "stop",
		Timestamp:    getCurrentTime(), // helper would be in internal
	}, nil
}

// getCurrentTime is placeholder - use time.Now in real.
func getCurrentTime() time.Time {
	return time.Now()
}
