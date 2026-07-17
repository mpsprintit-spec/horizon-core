// Package thinking provides the core intelligence engine for Horizon Core.
// It acts as the single brain, orchestrating all decision-making processes.
package thinking

import (
	"context"

	"github.com/project-horizon/horizon-core/services/ai/thinking/llm"
	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// ThinkingEngine defines the main interface for the thinking engine.
type ThinkingEngine interface {
	// Think processes a user request and returns a structured response.
	// It orchestrates the entire cognitive workflow.
	Think(ctx context.Context, req *model.ThinkingRequest) (*model.ThinkingResponse, error)
}

// ContextBuilder builds unified context from various sources.
type ContextBuilder interface {
	Build(ctx context.Context, req *model.ThinkingRequest) (*model.Context, error)
}

// PromptBuilder generates prompts for the LLM based on context.
type PromptBuilder interface {
	Build(ctx context.Context, ctxData *model.Context) (string, error)
}

// ResponseAnalyzer analyzes LLM responses and extracts structured data.
type ResponseAnalyzer interface {
	Analyze(ctx context.Context, llmResp string) (*model.Analysis, error)
}

// CapabilityChecker checks availability of various Horizon modules.
type CapabilityChecker interface {
	Check(ctx context.Context) (*model.Capability, error)
}

// WorkflowExecutor executes the thinking workflow steps.
type WorkflowExecutor interface {
	Execute(ctx context.Context, req *model.ThinkingRequest) (*model.ThinkingResponse, error)
}

// LLMAdapter provides abstraction for LLM interactions.
type LLMAdapter interface {
	Generate(ctx context.Context, prompt string, opts ...LLMOption) (*llm.LLMResponse, error)
}

// LLMOption represents optional configurations for LLM calls.
type LLMOption func(*llm.LLMConfig)

// KnowledgeAccessor provides access to long-term knowledge.
type KnowledgeAccessor interface {
	Query(ctx context.Context, query string) ([]string, error)
	Store(ctx context.Context, knowledge *model.LearningRequest) error
}

// WebSearcher handles information retrieval from external sources.
type WebSearcher interface {
	Search(ctx context.Context, query string) ([]string, error)
}

// DeviceExecutor executes actions on connected devices.
type DeviceExecutor interface {
	Execute(ctx context.Context, action *model.Action) error
}
