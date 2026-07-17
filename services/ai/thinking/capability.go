package thinking

import (
	"context"

	"github.com/project-horizon/horizon-core/services/ai/thinking/model"
)

// capabilityChecker implements CapabilityChecker.
type capabilityChecker struct {
	knowledgeAccessor KnowledgeAccessor
	webSearcher       WebSearcher
	deviceExecutor    DeviceExecutor
	// Add more as needed
}

// NewCapabilityChecker creates a capability checker with injected dependencies.
func NewCapabilityChecker(
	knowledge KnowledgeAccessor,
	search WebSearcher,
	device DeviceExecutor,
) CapabilityChecker {
	return &capabilityChecker{
		knowledgeAccessor: knowledge,
		webSearcher:       search,
		deviceExecutor:    device,
	}
}

// Check assesses available capabilities.
func (c *capabilityChecker) Check(ctx context.Context) (*model.Capability, error) {
	cap := &model.Capability{
		HasKnowledge:       c.knowledgeAccessor != nil,
		HasWebSearch:       c.webSearcher != nil,
		HasDevices:         c.deviceExecutor != nil,
		HasLearning:        true, // always available internally
		SupportedProviders: []string{"openai", "anthropic"},
	}
	return cap, nil
}
