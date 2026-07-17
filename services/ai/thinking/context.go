package thinking

import (
	"context"
	"fmt"

	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// contextBuilder implements ContextBuilder.
type contextBuilder struct {
	knowledgeAccessor KnowledgeAccessor
	webSearcher       WebSearcher
	deviceExecutor    DeviceExecutor
}

// NewContextBuilder creates a new context builder.
func NewContextBuilder(
	knowledge KnowledgeAccessor,
	search WebSearcher,
	device DeviceExecutor,
) ContextBuilder {
	return &contextBuilder{
		knowledgeAccessor: knowledge,
		webSearcher:       search,
		deviceExecutor:    device,
	}
}

// Build assembles context from all sources.
func (b *contextBuilder) Build(ctx context.Context, req *model.ThinkingRequest) (*model.Context, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}

	ctxData := &model.Context{
		UserInput:           req.Input,
		ConversationHistory: req.ConversationHistory,
		SessionContext:      make(map[string]interface{}),
		Metadata:            make(map[string]string),
	}

	// Fetch knowledge if available
	if b.knowledgeAccessor != nil {
		knowledge, err := b.knowledgeAccessor.Query(ctx, req.Input)
		if err == nil {
			ctxData.KnowledgeResults = knowledge
		}
	}

	// Web search if needed (placeholder logic)
	if b.webSearcher != nil {
		// In real impl, decide based on query
		searchRes, err := b.webSearcher.Search(ctx, req.Input)
		if err == nil {
			ctxData.SearchResults = searchRes
		}
	}

	// Device state
	if b.deviceExecutor != nil {
		// Placeholder
		ctxData.DeviceState = map[string]interface{}{"status": "online"}
	}

	return ctxData, nil
}
