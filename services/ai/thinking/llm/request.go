package llm

import "github.com/project-horizon/horizon-core/services/ai/thinking/model"

// LLMRequest wraps the input for LLM calls.
type LLMRequest struct {
	Prompt       string                   `json:"prompt"`
	Config       *LLMConfig               `json:"config"`
	SystemPrompt string                   `json:"system_prompt,omitempty"`
	History      []model.ConversationTurn `json:"history,omitempty"`
}
