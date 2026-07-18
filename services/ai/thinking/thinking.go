package thinking

import (
	"strings"

	"github.com/project-horizon/horizon-core/services/ai/activation"
	contextengine "github.com/project-horizon/horizon-core/services/ai/context"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

type ThinkingEngine struct {
	Activation *activation.Engine
	Context    *contextengine.Engine
}

type Thought struct {
	Concepts   []string
	Confidence float64
	Resonance  float64
	Conflicts  []string
}

func NewThinkingEngine(kb *knowledge.KnowledgeBase) *ThinkingEngine {
	return &ThinkingEngine{Activation: activation.NewEngine(kb), Context: contextengine.NewEngine(kb)}
}

func splitPrompt(prompt string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(prompt)))
}
