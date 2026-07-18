package thinking

import "github.com/project-horizon/horizon-core/services/ai/knowledge"


type ThinkingEngine struct {
	Kb *knowledge.KnowledgeBase
}

func NewThinkingEngine(kb *knowledge.KnowledgeBase) *ThinkingEngine {
	return &ThinkingEngine{Kb: kb}
}
