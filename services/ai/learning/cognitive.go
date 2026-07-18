package learning

import (
	"time"

	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

type ConceptEvidence struct {
	Token  string
	Source SourceEvidence
}

type CognitiveLearner struct{ Kb *knowledge.KnowledgeBase }

func NewCognitiveLearner(kb *knowledge.KnowledgeBase) *CognitiveLearner {
	return &CognitiveLearner{Kb: kb}
}
func (l *CognitiveLearner) LearnPath(tokens []string, success bool) {
	if len(tokens) == 0 {
		return
	}
	var prev *knowledge.ConceptNode
	for _, token := range tokens {
		node := l.Kb.Store(token)
		if success {
			node.Importance = minFloat(1, node.Importance+0.03)
			node.Plasticity = minFloat(1, node.Plasticity+0.02)
		} else {
			node.Plasticity = maxFloat(0.05, node.Plasticity-0.02)
		}
		if prev != nil {
			w, c := 0.58, 0.6
			if !success {
				w, c = 0.2, 0.35
			}
			l.Kb.Connect(prev, node, w, c, false)
		}
		prev = node
	}
}
func (l *CognitiveLearner) LearnEvidence(items []ConceptEvidence) {
	for _, item := range items {
		node := l.Kb.Store(item.Token)
		node.LastActivation = time.Now().UTC()
		node.Importance = minFloat(1, node.Importance+(item.Source.Confidence*item.Source.Reputation*0.02))
	}
}
func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
