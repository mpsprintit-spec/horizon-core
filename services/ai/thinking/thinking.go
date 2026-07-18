package thinking

import "horizon-core/services/ai/knowledge"

type ThinkingEngine struct {
	Kb *knowledge.KnowledgeBase
}

func NewThinkingEngine(kb *knowledge.KnowledgeBase) *ThinkingEngine {
	return &ThinkingEngine{Kb: kb}
}
