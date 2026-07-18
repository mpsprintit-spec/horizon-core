package thinking

import (
	"github.com/project-horizon/horizon-core/services/ai/activation"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

func (t *ThinkingEngine) Think(prompt string, contextTokens []string) (Thought, bool) {
	stimulus := splitPrompt(prompt)
	focus := t.Context.Focus(contextTokens)
	result := t.Activation.ActivateWith(activation.Request{StimulusTokens: stimulus, ContextBoosts: focus.Boosts, Cycles: 6})
	if !result.Converged {
		return Thought{}, false
	}

	thought := Thought{Resonance: result.Resonance}
	var confidenceTotal float64
	for _, node := range result.RankedNodes {
		if contains(stimulus, node.Token) {
			continue
		}
		if len(thought.Concepts) >= 5 {
			break
		}
		thought.Concepts = append(thought.Concepts, node.Token)
		confidenceTotal += result.Confidence[node.ID] * node.Activation
		if hasStrongInhibition(node) {
			thought.Conflicts = append(thought.Conflicts, node.Token)
		}
	}
	if len(thought.Concepts) == 0 {
		for _, node := range result.RankedNodes {
			thought.Concepts = append(thought.Concepts, node.Token)
			break
		}
	}
	thought.Confidence = clamp01((confidenceTotal + result.Resonance) / float64(len(result.RankedNodes)+1))
	return thought, true
}

func (t *ThinkingEngine) Reason(prompt string) ([]string, bool) {
	thought, ok := t.Think(prompt, nil)
	return thought.Concepts, ok
}

func hasStrongInhibition(node *knowledge.ConceptNode) bool {
	for _, synapse := range node.Synapses {
		if synapse.Inhibitory && synapse.Weight >= 0.6 {
			return true
		}
	}
	return false
}

func contains(tokens []string, token string) bool {
	for _, t := range tokens {
		if t == token {
			return true
		}
	}
	return false
}
func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
