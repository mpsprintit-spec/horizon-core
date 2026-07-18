package thinking

import (
	"github.com/project-horizon/horizon-core/services/ai/activation"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

func (t *ThinkingEngine) Think(prompt string, contextTokens []string) (Thought, bool) {
	stimulus := splitPrompt(prompt)
	wm := newWorkingMemory(contextTokens)
	defer wm.Clear()
	trace := PathTrace{}
	focus := t.Context.Focus(contextTokens)
	result := t.Activation.ActivateWith(activation.Request{StimulusTokens: stimulus, ContextBoosts: focus.Boosts, Cycles: 8})
	for _, node := range result.RankedNodes {
		trace.Add("Activation", node, result.Confidence[node.ID], "ranked activation")
	}
	if !result.Converged {
		t.LastState = CognitiveState{CurrentContext: contextTokens, UnknownNodes: stimulus, ReasoningDepth: 1}
		return Thought{NeedsWebSearch: true}, false
	}
	rep := t.Understanding.Understand(result, stimulus, wm, &trace)
	hypotheses := BuildHypotheses(rep, result)
	wm.Hypotheses = hypotheses
	state := CognitiveState{CurrentContext: contextTokens, FocusedNodes: rep.DominantNodes, SupportingNodes: rep.SupportingNodes, CompetingNodes: rep.CompetingNodes, UnknownNodes: rep.UnknownTokens, ConflictNodes: rep.ConflictNodes, ReasoningDepth: 3, ActivationHistory: []activation.Result{result}, Confidence: rep.Confidence}
	assessment := t.MetaCognition.Evaluate(state, rep)
	thought := Thought{Resonance: rep.Resonance, Confidence: rep.Confidence, Hypotheses: hypotheses, NeedsWebSearch: assessment.NeedWebSearch}
	best := chooseBestHypothesis(hypotheses)
	for _, h := range hypotheses {
		if h.Confidence < 0.2 {
			continue
		}
		for _, id := range h.Nodes {
			if n := t.Activation.Memory.Registry.GetByID(id); n != nil && !contains(thought.Concepts, n.Token) {
				thought.Concepts = append(thought.Concepts, n.Token)
				trace.Add("Hypothesis", n, h.Confidence, "candidate hypothesis")
			}
		}
	}
	_ = best
	for _, id := range rep.ConflictNodes {
		if n := t.Activation.Memory.Registry.GetByID(id); n != nil {
			thought.Conflicts = append(thought.Conflicts, n.Token)
		}
	}
	if thought.Confidence < t.MetaCognition.MinConfidence {
		thought.NeedsWebSearch = true
	}
	trace.Add("Decision", nil, thought.Confidence, "runtime decision completed")
	t.LastState, t.LastTrace = state, trace
	return thought, len(thought.Concepts) > 0 && thought.Confidence >= 0.2
}

func (t *ThinkingEngine) Reason(prompt string) ([]string, bool) {
	thought, ok := t.Think(prompt, nil)
	return thought.Concepts, ok
}
func chooseBestHypothesis(hs []Hypothesis) Hypothesis {
	var best Hypothesis
	for _, h := range hs {
		if h.Confidence > best.Confidence {
			best = h
		}
	}
	return best
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
