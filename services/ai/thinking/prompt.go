package thinking

import (
	"context"
	"fmt"
	"strings"

	"github.com/project-horizon/horizon-core/services/ai/thinking/internal"
	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// promptBuilder implements PromptBuilder.
type promptBuilder struct{}

// NewPromptBuilder creates a new prompt builder.
func NewPromptBuilder() PromptBuilder {
	return &promptBuilder{}
}

// Build constructs the full prompt for LLM.
func (p *promptBuilder) Build(ctx context.Context, ctxData *model.Context) (string, error) {
	if ctxData == nil {
		return "", fmt.Errorf("nil context data")
	}

	var sb strings.Builder

	sb.WriteString("You are Horizon Thinking Engine, the central intelligence.\n\n")
	sb.WriteString("User Query: " + ctxData.UserInput + "\n\n")

	if len(ctxData.ConversationHistory) > 0 {
		sb.WriteString("Conversation History:\n")
		for _, turn := range ctxData.ConversationHistory {
			sb.WriteString(turn.Role + ": " + turn.Content + "\n")
		}
		sb.WriteString("\n")
	}

	if len(ctxData.KnowledgeResults) > 0 {
		sb.WriteString("Relevant Knowledge:\n")
		for _, k := range ctxData.KnowledgeResults {
			sb.WriteString("- " + k + "\n")
		}
		sb.WriteString("\n")
	}

	if len(ctxData.SearchResults) > 0 {
		sb.WriteString("Search Results:\n")
		for _, s := range ctxData.SearchResults {
			sb.WriteString("- " + s + "\n")
		}
		sb.WriteString("\n")
	}

	// Instructions for structured output
	sb.WriteString("Analyze the query and respond in the following JSON format:\n")
	sb.WriteString(`{
  "intent": "string",
  "confidence": 0.0,
  "response_text": "string",
  "actions": [],
  "learning_items": []
}`)

	prompt := internal.CleanPrompt(sb.String())
	return prompt, nil
}
