package thinking

import (
	"context"
	"fmt"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// workflowExecutor implements WorkflowExecutor.
type workflowExecutor struct {
	contextBuilder    ContextBuilder
	capabilityChecker CapabilityChecker
	promptBuilder     PromptBuilder
	llmAdapter        LLMAdapter
	responseAnalyzer  ResponseAnalyzer
	knowledgeAccessor KnowledgeAccessor
}

// NewWorkflowExecutor creates the workflow executor with dependencies.
func NewWorkflowExecutor(
	ctxBuilder ContextBuilder,
	capChecker CapabilityChecker,
	promptB PromptBuilder,
	llm LLMAdapter,
	analyzer ResponseAnalyzer,
	knowledge KnowledgeAccessor,
) WorkflowExecutor {
	return &workflowExecutor{
		contextBuilder:    ctxBuilder,
		capabilityChecker: capChecker,
		promptBuilder:     promptB,
		llmAdapter:        llm,
		responseAnalyzer:  analyzer,
		knowledgeAccessor: knowledge,
	}
}

// Execute runs the full thinking workflow.
func (w *workflowExecutor) Execute(ctx context.Context, req *model.ThinkingRequest) (*model.ThinkingResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}

	// 1. Build context
	ctxData, err := w.contextBuilder.Build(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("context build failed: %w", err)
	}

	// 2. Capability check
	_, err = w.capabilityChecker.Check(ctx)
	if err != nil {
		// Continue gracefully
	}

	// 3. Build prompt
	prompt, err := w.promptBuilder.Build(ctx, ctxData)
	if err != nil {
		return nil, fmt.Errorf("prompt build failed: %w", err)
	}

	// 4. LLM call
	llmResp, err := w.llmAdapter.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("llm call failed: %w", err)
	}

	// 5. Analyze response
	analysis, err := w.responseAnalyzer.Analyze(ctx, llmResp.Content)
	if err != nil {
		// Partial success
		analysis = &model.Analysis{ResponseText: llmResp.Content}
	}

	// 6. Handle learning if needed
	if len(analysis.LearningItems) > 0 && w.knowledgeAccessor != nil {
		for _, item := range analysis.LearningItems {
			_ = w.knowledgeAccessor.Store(ctx, &item) // fire and forget
		}
	}

	// 7. Construct final response
	resp := &model.ThinkingResponse{
		ID:                  req.ID,
		Output:              analysis.ResponseText,
		Confidence:          analysis.Confidence,
		Actions:             analysis.Actions,
		LearningSuggestions: analysis.LearningItems,
		Timestamp:           getCurrentTime(),
	}

	return resp, nil
}

// Placeholder helper.
func getCurrentTime() time.Time {
	return time.Now()
}
