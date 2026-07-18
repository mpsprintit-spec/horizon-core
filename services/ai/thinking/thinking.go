package thinking

import (
	"strings"

	"github.com/project-horizon/horizon-core/services/ai/activation"
	contextengine "github.com/project-horizon/horizon-core/services/ai/context"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

type ThinkingEngine struct {
	Activation    *activation.Engine
	Context       *contextengine.Engine
	Understanding *UnderstandingEngine
	MetaCognition *MetaCognitiveLayer
	LastState     CognitiveState
	LastTrace     PathTrace
}

type Thought struct {
	Concepts       []string
	Confidence     float64
	Resonance      float64
	Conflicts      []string
	Hypotheses     []Hypothesis
	NeedsWebSearch bool
}

func NewThinkingEngine(kb *knowledge.KnowledgeBase) *ThinkingEngine {
	return &ThinkingEngine{Activation: activation.NewEngine(kb), Context: contextengine.NewEngine(kb), Understanding: NewUnderstandingEngine(kb), MetaCognition: NewMetaCognitiveLayer()}
}

func splitPrompt(prompt string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(prompt)))
}
