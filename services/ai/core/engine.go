package core

import (
	"horizon-core/services/ai/execution"
	"horizon-core/services/ai/knowledge"
	"horizon-core/services/ai/learning"
	"horizon-core/services/ai/thinking"
)

type HorizonEngine struct {
	Knowledge *knowledge.KnowledgeBase
	Learning  *learning.LearningUnit
	Thinking  *thinking.ThinkingEngine
	Execution *execution.ExecutionCore
}

func NewHorizonEngine() *HorizonEngine {
	kb := knowledge.NewKnowledgeBase()
	return &HorizonEngine{
		Knowledge: kb,
		Learning:  learning.NewLearningUnit(kb),
		Thinking:  thinking.NewThinkingEngine(kb),
		Execution: execution.NewExecutionCore(),
	}
}
