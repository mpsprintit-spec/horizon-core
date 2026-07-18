package learning

import "github.com/project-horizon/horizon-core/services/ai/knowledge"

type LearningUnit struct{ Kb *knowledge.KnowledgeBase }

func NewLearningUnit(kb *knowledge.KnowledgeBase) *LearningUnit { return &LearningUnit{Kb: kb} }
