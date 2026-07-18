package thinking

import (
	"sort"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/activation"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

type CognitiveState struct {
	CurrentContext    []string
	FocusedNodes      []knowledge.NodeID
	SupportingNodes   []knowledge.NodeID
	CompetingNodes    []knowledge.NodeID
	UnknownNodes      []string
	ConflictNodes     []knowledge.NodeID
	ReasoningDepth    int
	ActivationHistory []activation.Result
	Confidence        float64
	CurrentGoal       string
}

type WorkingMemory struct {
	Activation          map[knowledge.NodeID]float64
	Hypotheses          []Hypothesis
	Candidates          []knowledge.NodeID
	Context             []string
	TemporaryRelations  map[knowledge.NodeID][]knowledge.NodeID
	TemporaryConfidence map[knowledge.NodeID]float64
}

type CognitiveRepresentation struct {
	DominantNodes       []knowledge.NodeID
	SupportingNodes     []knowledge.NodeID
	CompetingNodes      []knowledge.NodeID
	UnknownTokens       []string
	ConflictNodes       []knowledge.NodeID
	HiddenRelations     [][2]knowledge.NodeID
	Confidence          float64
	Resonance           float64
	ActivationSignature map[knowledge.NodeID]float64
}

type Hypothesis struct {
	Nodes      []knowledge.NodeID
	Confidence float64
	Evidence   int
	Conflicts  int
}

type PathTrace struct{ Steps []TraceStep }
type TraceStep struct {
	Stage      string
	NodeID     knowledge.NodeID
	Token      string
	Confidence float64
	Note       string
	Time       time.Time
}

func (p *PathTrace) Add(stage string, node *knowledge.ConceptNode, confidence float64, note string) {
	step := TraceStep{Stage: stage, Confidence: confidence, Note: note, Time: time.Now().UTC()}
	if node != nil {
		step.NodeID = node.ID
		step.Token = node.Token
	}
	p.Steps = append(p.Steps, step)
}

func newWorkingMemory(context []string) *WorkingMemory {
	return &WorkingMemory{Activation: map[knowledge.NodeID]float64{}, Context: context, TemporaryRelations: map[knowledge.NodeID][]knowledge.NodeID{}, TemporaryConfidence: map[knowledge.NodeID]float64{}}
}
func (w *WorkingMemory) Clear() { *w = WorkingMemory{} }

type UnderstandingEngine struct{ Memory *knowledge.KnowledgeBase }

func NewUnderstandingEngine(memory *knowledge.KnowledgeBase) *UnderstandingEngine {
	return &UnderstandingEngine{Memory: memory}
}

func (u *UnderstandingEngine) Understand(result activation.Result, stimulus []string, wm *WorkingMemory, trace *PathTrace) CognitiveRepresentation {
	rep := CognitiveRepresentation{ActivationSignature: result.Activations, Resonance: result.Resonance}
	stimulusSet := map[string]bool{}
	for _, token := range stimulus {
		stimulusSet[token] = true
		if u.Memory.Fetch(token) == nil {
			rep.UnknownTokens = append(rep.UnknownTokens, token)
		}
	}
	for _, node := range result.RankedNodes {
		level := result.Activations[node.ID]
		conf := result.Confidence[node.ID]
		if stimulusSet[node.Token] && len(result.RankedNodes) > 1 {
			continue
		}
		if level >= 0.62 && len(rep.DominantNodes) < 3 {
			rep.DominantNodes = append(rep.DominantNodes, node.ID)
			trace.Add("Understanding", node, conf, "dominant convergence")
		} else if level >= 0.25 && len(rep.SupportingNodes) < 8 {
			rep.SupportingNodes = append(rep.SupportingNodes, node.ID)
			trace.Add("Understanding", node, conf, "supporting activation")
		}
		if hasStrongInhibition(node) {
			rep.ConflictNodes = append(rep.ConflictNodes, node.ID)
		}
		for _, syn := range node.Synapses {
			if syn.Inhibitory && result.Activations[syn.TargetID] > 0.25 {
				rep.CompetingNodes = appendUniqueID(rep.CompetingNodes, syn.TargetID)
			}
		}
	}
	rep.HiddenRelations = u.hiddenRelations(rep.DominantNodes, rep.SupportingNodes)
	rep.Confidence = clamp01((result.Resonance + averageConfidence(result, rep.DominantNodes, rep.SupportingNodes)) / 2)
	wm.Candidates = append(append([]knowledge.NodeID{}, rep.DominantNodes...), rep.SupportingNodes...)
	for id, c := range result.Confidence {
		wm.TemporaryConfidence[id] = c
	}
	return rep
}

func (u *UnderstandingEngine) hiddenRelations(dominant, supporting []knowledge.NodeID) [][2]knowledge.NodeID {
	var out [][2]knowledge.NodeID
	for _, a := range dominant {
		na := u.Memory.Registry.GetByID(a)
		if na == nil {
			continue
		}
		for _, b := range supporting {
			if a == b {
				continue
			}
			if _, ok := na.Synapses[b]; ok {
				out = append(out, [2]knowledge.NodeID{a, b})
			}
		}
	}
	return out
}

func BuildHypotheses(rep CognitiveRepresentation, result activation.Result) []Hypothesis {
	ids := append(append([]knowledge.NodeID{}, rep.DominantNodes...), rep.SupportingNodes...)
	sort.Slice(ids, func(i, j int) bool {
		return result.Confidence[ids[i]]*result.Activations[ids[i]] > result.Confidence[ids[j]]*result.Activations[ids[j]]
	})
	var hs []Hypothesis
	for i, id := range ids {
		if i >= 4 {
			break
		}
		conflicts := 0
		if containsID(rep.ConflictNodes, id) || containsID(rep.CompetingNodes, id) {
			conflicts = 1
		}
		hs = append(hs, Hypothesis{Nodes: []knowledge.NodeID{id}, Confidence: clamp01(result.Confidence[id]*result.Activations[id] - float64(conflicts)*0.2), Evidence: len(rep.HiddenRelations) + 1, Conflicts: conflicts})
	}
	return hs
}

func appendUniqueID(ids []knowledge.NodeID, id knowledge.NodeID) []knowledge.NodeID {
	if containsID(ids, id) {
		return ids
	}
	return append(ids, id)
}
func containsID(ids []knowledge.NodeID, id knowledge.NodeID) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
func averageConfidence(result activation.Result, groups ...[]knowledge.NodeID) float64 {
	var sum float64
	var count int
	for _, g := range groups {
		for _, id := range g {
			sum += result.Confidence[id] * result.Activations[id]
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
