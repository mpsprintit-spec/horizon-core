package thinking

import (
	"context"
	"fmt"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/thinking/llm"
	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// engine implements ThinkingEngine.
type engine struct {
	workflow WorkflowExecutor
	llmMgr   LLMAdapter
	config   *Config
}

// Config holds configuration for the Thinking Engine.
type Config struct {
	LLMConfig *llm.LLMConfig
	// Add more configs as needed
}

// NewEngine creates a fully configured Thinking Engine.
func NewEngine(
	cfg *Config,
	ctxBuilder ContextBuilder,
	capChecker CapabilityChecker,
	promptBuilder PromptBuilder,
	llmAdapter LLMAdapter,
	analyzer ResponseAnalyzer,
	knowledge KnowledgeAccessor,
) (ThinkingEngine, error) {
	if cfg == nil {
		cfg = &Config{
			LLMConfig: llm.DefaultConfig(),
		}
	}

	wf := NewWorkflowExecutor(
		ctxBuilder,
		capChecker,
		promptBuilder,
		llmAdapter,
		analyzer,
		knowledge,
	)

	return &engine{
		workflow: wf,
		llmMgr:   llmAdapter,
		config:   cfg,
	}, nil
}

// Think is the main entry point.
func (e *engine) Think(ctx context.Context, req *model.ThinkingRequest) (*model.ThinkingResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("nil thinking request")
	}

	if req.ID == "" {
		req.ID = generateID()
	}
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	resp, err := e.workflow.Execute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("thinking workflow failed: %w", err)
	}

	return resp, nil
}

// generateID creates a simple request ID.
func generateID() string {
	return fmt.Sprintf("req_%d", time.Now().UnixNano())
}
