package llm

import "time"

// LLMResponse represents response from LLM provider.
type LLMResponse struct {
	Content      string                 `json:"content"`
	Model        string                 `json:"model"`
	TokensUsed   int                    `json:"tokens_used"`
	FinishReason string                 `json:"finish_reason"`
	Timestamp    time.Time              `json:"timestamp"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}
