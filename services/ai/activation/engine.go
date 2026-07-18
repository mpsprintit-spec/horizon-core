package activation

import (
	"sort"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

type Engine struct {
	Memory     *knowledge.KnowledgeBase
	Decay      float64
	SpreadRate float64
	Threshold  float64
	Inhibition float64
}

type Request struct {
	StimulusTokens []string
	ContextBoosts  map[knowledge.NodeID]float64
	Cycles         int
	Now            time.Time
}

type Result struct {
	Converged   bool
	Resonance   float64
	Activations map[knowledge.NodeID]float64
	Confidence  map[knowledge.NodeID]float64
	RankedNodes []*knowledge.ConceptNode
}

func NewEngine(memory *knowledge.KnowledgeBase) *Engine {
	return &Engine{Memory: memory, Decay: 0.12, SpreadRate: 0.65, Threshold: 0.25, Inhibition: 0.55}
}

func (e *Engine) Activate(tokens []string, cycles int) Result {
	return e.ActivateWith(Request{StimulusTokens: tokens, Cycles: cycles, Now: time.Now().UTC()})
}

func (e *Engine) ActivateWith(req Request) Result {
	if req.Cycles < 1 {
		req.Cycles = 1
	}
	if req.Now.IsZero() {
		req.Now = time.Now().UTC()
	}
	state := map[knowledge.NodeID]float64{}
	confidence := map[knowledge.NodeID]float64{}
	for _, token := range req.StimulusTokens {
		if n := e.Memory.Fetch(token); n != nil {
			state[n.ID] = 1
			confidence[n.ID] = 1
		}
	}
	for id, boost := range req.ContextBoosts {
		state[id] += boost
		confidence[id] = max(confidence[id], boost)
	}
	state = normalize(state)

	for i := 0; i < req.Cycles; i++ {
		next := map[knowledge.NodeID]float64{}
		nextConfidence := map[knowledge.NodeID]float64{}
		for _, n := range e.Memory.Registry.Nodes() {
			adaptive := n.Threshold - (n.Importance * 0.05) - (float64(n.Frequency) * 0.001)
			n.Threshold = clamp(adaptive, 0.12, 0.8)
			next[n.ID] = n.RestingActivation
		}
		for id, level := range state {
			n := e.Memory.Registry.GetByID(id)
			if n == nil {
				continue
			}
			next[id] += level * (1 - e.Decay)
			nextConfidence[id] = max(nextConfidence[id], confidence[id])
			for _, s := range n.Synapses {
				agePenalty := temporalPenalty(req.Now, s.LastActivation)
				pulse := level * s.Weight * s.Confidence * e.SpreadRate * agePenalty
				if s.Inhibitory {
					next[s.TargetID] -= pulse * e.Inhibition
				} else {
					next[s.TargetID] += pulse
				}
				nextConfidence[s.TargetID] = max(nextConfidence[s.TargetID], confidence[id]*s.Confidence*agePenalty)
				s.Activation = pulse
			}
		}
		state = normalize(next)
		confidence = normalize(nextConfidence)
	}
	return e.converge(state, confidence, req.Now)
}

func (e *Engine) converge(state, confidence map[knowledge.NodeID]float64, now time.Time) Result {
	var ranked []*knowledge.ConceptNode
	var resonance float64
	for id, level := range state {
		n := e.Memory.Registry.GetByID(id)
		if n == nil {
			continue
		}
		n.Activation = level
		if level >= max(e.Threshold, n.Threshold) {
			n.Frequency++
			n.LastActivation = now
			ranked = append(ranked, n)
			resonance += level * max(confidence[id], 0.1)
		}
	}
	sort.Slice(ranked, func(i, j int) bool { return ranked[i].Activation > ranked[j].Activation })
	return Result{Converged: len(ranked) > 0, Resonance: resonance, Activations: state, Confidence: confidence, RankedNodes: ranked}
}

func temporalPenalty(now, last time.Time) float64 {
	if last.IsZero() {
		return 0.7
	}
	days := now.Sub(last).Hours() / 24
	if days <= 1 {
		return 1
	}
	return clamp(1-(days*0.01), 0.35, 1)
}

func normalize(in map[knowledge.NodeID]float64) map[knowledge.NodeID]float64 {
	for k, v := range in {
		in[k] = clamp(v, 0, 1)
	}
	return in
}

func clamp(v, low, high float64) float64 {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
