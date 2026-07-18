package learning

import "horizon-core/services/ai/knowledge"

type LearningUnit struct {
	Kb *knowledge.KnowledgeBase
}

func NewLearningUnit(kb *knowledge.KnowledgeBase) *LearningUnit {
	return &LearningUnit{Kb: kb}
}
