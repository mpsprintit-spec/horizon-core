package core

import (
	"github.com/project-horizon/horizon-core/services/ai/execution"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
	"github.com/project-horizon/horizon-core/services/ai/learning"
	"github.com/project-horizon/horizon-core/services/ai/thinking"
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
